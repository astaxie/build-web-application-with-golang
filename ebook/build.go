package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 定义一个访问者结构体
type Visitor struct{}

func (self *Visitor) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		return nil
	} else {
		if strings.HasSuffix(f.Name(), ".md") {
			fmt.Println(f)
			file, err := os.Open(f.Name())
			if err != nil {
				return err
			}
			input, _ := ioutil.ReadAll(file)
			input = regexp.MustCompile("\\[(.*?)\\]\\(<?(.*?)\\.md>?\\)").ReplaceAll(input, []byte("[$1](<$2.html>)"))
			output := blackfriday.MarkdownCommon(input)
			var out *os.File
			if out, err = os.Create(strings.Replace(f.Name(), ".md", ".html", -1)); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating %s: %v", f.Name(), err)
				os.Exit(-1)
			}
			defer out.Close()
			header := "<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"utf-8\"/>\n</head>\n"
			footer := "</html>\n"
			out.Write([]byte(header))
			if _, err = out.Write(output); err != nil {
				fmt.Fprintln(os.Stderr, "Error writing output:", err)
				os.Exit(-1)
			}
			out.Write([]byte(footer))
		}
	}
	return nil
}

func main() {
	v := &Visitor{}
	err := filepath.Walk("./", func(path string, f os.FileInfo, err error) error {
		return v.visit(path, f, err)
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
