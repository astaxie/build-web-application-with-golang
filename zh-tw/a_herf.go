package main

import(
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sort"
)
func dir()([]string,error) {
    path, err := os.Getwd()
    if err != nil {
        fmt.Println("err is:", err)
    }
    log.Println(path)
    path =path +"/*.html"

    fmt.Println(path)
    files,err := filepath.Glob(path)
    var s =make([]string,len(files))
    var head uint8 =0
    for _,k :=range files {
        filename := filepath.Base(k)
        head=filename[0]
        if (head < 52) {
            s = append(s, filename)
            fmt.Println(filename)
        }
    }
        sort.Strings(s)

    return s,err
    }


func  htmlfile(filename string,next_path string,last_path string)(error){
       file,err:= os.OpenFile("./"+filename,os.O_RDWR,0666)
       if err !=nil{
           fmt.Println("something is err :",err)
       }
       defer file.Close()
       var add_string1 string = "\n<a href=\"./"+next_path+"\">下一页</a>\n"
       var add_string2 string = "\n<a href=\"./"+last_path+"\">下一页</a>\n"
       file.Seek(1,2)
       _,err=file.WriteString(add_string1)
       if(err!=nil){
           fmt.Println("err:",err)
       }
       _,err=file.WriteString(add_string2)
       file.Seek(0,0)
       if(err!=nil){
           fmt.Println("err:",err)
       }
       var f =make([]byte,50000)
       _,err=file.Read(f)
       if(err!=nil){
           fmt.Println("error:",err)
       }
       //fmt.Println(string(f))
     return err
}
func nextandlast(filenames []string,index int )(filename string,next_path string,last_path string){
        fmt.Println(index,"   ---",index+1)
        filename = filenames[index]
        if(0<index && index<len(filenames)-1) {
            next_path = filenames[index+1]
            last_path = filenames[index-1]
        }else if(index==0){
          next_path =filenames[index+1]
          last_path = filenames[len(filenames)-1]
        }else if(index==len(filenames)){
            next_path =filenames[0]
            last_path =filenames[index-1]
        }else{
           return
         }
        return filename,next_path,last_path
}


func main(){
      files,err:=dir()
      if(err!=nil){
        fmt.Println(err)
      }
      //fmt.Println(files)
     // var tmp string ="\0"
      for i, v :=range files{
          if(v!=""){
              fmt.Println(v)
           filename,next_path,last_path :=nextandlast(files,i)
           err:=htmlfile(filename,next_path,last_path)
           if(err!=nil){
               fmt.Println("err=",err)
           }else if(v==""){
               continue
               }
          }
      }
      log.Println("end")

      return
    }
