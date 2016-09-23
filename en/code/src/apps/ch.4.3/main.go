// Example code for Chapter 4.3 from "Build Web Application with Golang"
// Purpose: Shows how to properly escape input
package main

import (
	"html/template"
	"net/http"
	textTemplate "text/template"
)

var t *template.Template = template.Must(template.ParseFiles("index.gtpl"))

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userInput := r.Form.Get("userinput")
	if 0 < len(r.Form.Get("escape")) {
		t.Execute(w, template.HTMLEscapeString(userInput))
	} else {
		// Variables with type `template.HTML` are not escaped when passed to `.Execute()`
		t.Execute(w, template.HTML(userInput))
	}
}
func templateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userInput := r.Form.Get("userinput")
	if 0 < len(r.Form.Get("escape")) {
		// `html/template.Execute()` escapes input
		t.Execute(w, userInput)
	} else {
		tt := textTemplate.Must(textTemplate.ParseFiles("index.gtpl"))
		// `text/template.Execute()` doesn't escape input
		tt.Execute(w, userInput)
	}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/template", templateHandler)
	http.ListenAndServe(":9090", nil)
}
