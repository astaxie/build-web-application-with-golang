package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
)

//定义一个访问者结构体
type Visitor struct{}

//定义其碰到目录的时候的行为
func (v *Visitor) VisitDir(path string, f *os.FileInfo) bool {
	fmt.Println(path)
	return true
}

//定义其碰到文件的时候的行为
func (v *Visitor) VisitFile(path string, f *os.FileInfo) {
	fmt.Println(path)
	//@todo 如何读取md文件转化为html
	input, err := ioutil.ReadAll()
	output := blackfriday.MarkdownBasic(input)
}

func main() {
	v := &Visitor{}
	errors := make(chan os.Error, 64) //错误消息使用64个缓存，可以随意
	filepath.Walk("./", v, errors)
	select {
	case err := <-errors:
		panic(err)
	default:
	}
}
