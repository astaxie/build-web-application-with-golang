// Example code for Chapter 4.5
// Purpose is to create a server to handle uploading files.
package main

import (
	"apps/ch.4.4/nonce"
	"apps/ch.4.4/validator"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const MiB_UNIT = 1 << 20

var t *template.Template
var submissions nonce.Nonces = nonce.New()

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "index", submissions.NewToken())
	checkError(err)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var errs validator.Errors
	r.ParseMultipartForm(32 * MiB_UNIT)
	token := r.Form.Get("token")
	if err := submissions.CheckThenMarkToken(token); err != nil {
		errs = validator.Errors{[]error{err}}
	} else {
		file, handler, err := r.FormFile("uploadfile")
		checkError(err)
		saveUpload(file, handler)
	}
	err := t.ExecuteTemplate(w, "upload", errs)
	checkError(err)
}
func saveUpload(file multipart.File, handler *multipart.FileHeader) {
	defer file.Close()
	fmt.Printf("Uploaded file info: %#v", handler.Header)
	localFilename := fmt.Sprintf("./uploads/%v.%v", handler.Filename, submissions.NewToken())
	f, err := os.OpenFile(localFilename, os.O_WRONLY|os.O_CREATE, 0666)
	checkError(err)
	defer f.Close()
	_, err = io.Copy(f, file)
	checkError(err)
}
func init() {
	var err error
	t, err = template.ParseFiles("index.gtpl", "upload.gtpl")
	checkError(err)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":9090", nil)
	checkError(err)
}
