# 8.4 RPC

In previous sections we talked about how to write network applications based on Sockets and HTTP. We learned that both of them use the "information exchange" model, in which clients send requests and servers respond to them. This kind of data exchange is based on a specific format so that both sides are able to communicate with one another. However, many independent applications do not use this model, but instead call services just like they would call normal functions.

RPC was intended to be the function call mode for networked systems. Clients execute RPCs like they call native functions, except they package the function parameters and send them through the network to the server. The server can then unpack these parameters and process the request, executing the results back to the client.

In computer science, a remote procedure call (RPC) is a type of inter-process communication that allows a computer program to cause a subroutine or procedure to execute in another address space (commonly on another computer on a shared network) without the programmer explicitly coding the details for this remote interaction. That is, the programmer writes essentially the same code whether the subroutine is local to the executing program, or remote. When the software in question uses object-oriented principles, RPC is called remote invocation or remote method invocation.

## RPC working principle

![](images/8.4.rpc.png?raw=true)

Figure 8.8 RPC working principle

Normally, an RPC call from client to server has the following ten steps:

- 1. Call the client handle, execute transfer arguments.
- 2. Call local system kernel to send network messages.
- 3. Send messages to remote hosts.
- 4. The server receives handle and arguments.
- 5. Execute remote processes.
- 6. Return execution result to corresponding handle.
- 7. The server handle calls remote system kernel.
- 8. Messages sent back to local system kernel.
- 9. The client handle receives messages from system kernel.
- 10. The client gets results from corresponding handle.

## Go RPC

Go has official support for RPC in its standard library on three levels, which are TCP, HTTP and JSON RPC. Note that Go RPC is not like other traditional RPC systems. It requires you to use Go applications on both client and server sides because it encodes content using Gob.

Functions of Go RPC must abide by the following rules for remote access, otherwise the corresponding calls will be ignored.

- Functions are exported (capitalized).
- Functions must have two arguments with exported types.
- The first argument is for receiving from the client, and the second one has to be a pointer and is for replying to the client.
- Functions must have a return value of error type.

For example:

	func (t *T) MethodName(argType T1, replyType *T2) error

Where T, T1 and T2 must be able to be encoded by the `package/gob` package.

Any kind of RPC has to go through a network to transfer data. Go RPC can either use HTTP or TCP. The benefits of using HTTP is that you can reuse some functions from the `net/http` package.

### HTTP RPC 

HTTP server side code:

	package main

	import (
		"errors"
		"fmt"
		"net/http"
		"net/rpc"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

	func main() {

		arith := new(Arith)
		rpc.Register(arith)
		rpc.HandleHTTP()

		err := http.ListenAndServe(":1234", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

We registered a RPC service of Arith, then registered this service on HTTP through `rpc.HandleHTTP`. After that, we are able to transfer data through HTTP.

Client side code:

	package main

	import (
		"fmt"
		"log"
		"net/rpc"
		"os"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}


	func main() {
		if len(os.Args) != 2 {
			fmt.Println("Usage: ", os.Args[0], "server")
			os.Exit(1)
		}
		serverAddress := os.Args[1]

		client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		args := Args{17, 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

		var quot Quotient
		err = client.Call("Arith.Divide", args, &quot)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	}

We compile the client and the server side code separately then start the server and client. You'll then have something similar as follows after you input some data.

	$ ./http_c localhost
	Arith: 17*8=136
	Arith: 17/8=2 remainder 1

As you can see, we defined a struct for the return type. We use it as type of function argument on the server side, and as the type of the second and third arguments on the client `client.Call`. This call is very important. It has three arguments, where the first one is the name of the function that is going to be called, the second is the argument you want to pass, and the last one is the return value (of pointer type). So far we see that it's easy to implement RPC in Go.

### TCP RPC

Let's try the RPC that is based on TCP, here is the server side code:

	package main

	import (
		"errors"
		"fmt"
		"net"
		"net/rpc"
		"os"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

	func main() {

		arith := new(Arith)
		rpc.Register(arith)

		tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
		checkError(err)

		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			rpc.ServeConn(conn)
		}

	}

	func checkError(err error) {
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		}
	}

The difference between HTTP RPC and TCP RPC is that we have to control connections by ourselves if we use TCP RPC, then pass connections to RPC for processing.

As you may have guessed, this is a blocking pattern. You are free to use goroutines to extend this application as a more advanced experiment.

The client side code:

	package main

	import (
		"fmt"
		"log"
		"net/rpc"
		"os"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	func main() {
		if len(os.Args) != 2 {
			fmt.Println("Usage: ", os.Args[0], "server:port")
			os.Exit(1)
		}
		service := os.Args[1]

		client, err := rpc.Dial("tcp", service)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		args := Args{17, 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

		var quot Quotient
		err = client.Call("Arith.Divide", args, &quot)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	}

The only difference in the client side code is that HTTP clients use DialHTTP whereas TCP clients use Dial(TCP).

### JSON RPC

JSON RPC encodes data to JSON instead of gob. Let's see an example of a Go JSON RPC on the server:

	package main

	import (
		"errors"
		"fmt"
		"net"
		"net/rpc"
		"net/rpc/jsonrpc"
		"os"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

	func main() {

		arith := new(Arith)
		rpc.Register(arith)

		tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
		checkError(err)

		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			jsonrpc.ServeConn(conn)
		}

	}

	func checkError(err error) {
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		}
	}

JSON RPC is based on TCP and doesn't support HTTP yet.

The client side code:

	package main

	import (
		"fmt"
		"log"
		"net/rpc/jsonrpc"
		"os"
	)

	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	func main() {
		if len(os.Args) != 2 {
			fmt.Println("Usage: ", os.Args[0], "server:port")
			log.Fatal(1)
		}
		service := os.Args[1]

		client, err := jsonrpc.Dial("tcp", service)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		args := Args{17, 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

		var quot Quotient
		err = client.Call("Arith.Divide", args, &quot)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	}

## Summary

Go has good support for HTTP, TPC and JSON RPC implementation which allow us to easily develop distributed web applications; however, it is regrettable that Go doesn't have built-in support for SOAP RPC, although some open source third-party packages do offer this.

## Links

- [Directory](preface.md)
- Previous section: [REST](08.3.md)
- Next section: [Summary](08.5.md)
