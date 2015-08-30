// Example code for Chapter 5.6 from "Build Web Application with Golang"
// Purpose: Shows you have to perform basic CRUD operations for a redis driver.
package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	DB_PORT = "9191"
	DB_URL  = "127.0.0.1"
)

func main() {
	var client goredis.Client

	// Set the default port in Redis
	client.Addr = DB_URL + ":" + DB_PORT

	// string manipulation
	fmt.Println("Inserting")
	err := client.Set("a", []byte("hello"))
	checkError(err)

	// list operation
	vals := []string{"a", "b", "c", "d"}
	for _, v := range vals {
		err = client.Rpush("l", []byte(v))
		checkError(err)
	}
	fmt.Println("Updating")
	err = client.Set("a", []byte("a is for apple"))
	checkError(err)
	err = client.Rpush("l", []byte("e"))
	checkError(err)

	fmt.Println("Querying")
	val, err := client.Get("a")
	checkError(err)
	fmt.Println(string(val))

	dbvals, err := client.Lrange("l", 0, 4)
	checkError(err)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}

	fmt.Println("Deleting")
	_, err = client.Del("l")
	checkError(err)
	_, err = client.Del("a")
	checkError(err)
}
