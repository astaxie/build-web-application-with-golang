package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func postFile(filename string, targetUrl string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	checkError(err)

	fh, err := os.Open(filename)
	checkError(err)

	_, err = io.Copy(fileWriter, fh)
	checkError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	checkError(err)

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
func main() {
	target_url := "http://localhost:9090/upload"
	filename := "../file.txt"
	postFile(filename, target_url)
}
