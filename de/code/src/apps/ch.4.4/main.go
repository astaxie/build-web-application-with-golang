// Example code for Chapter 3.2 from "Build Web Application with Golang"
// Purpose: Shows how to prevent duplicate submissions by using tokens
// Example code for Chapter 4.4 based off the code from Chapter 4.2
// Run `go run main.go` then access http://localhost:9090
package main

import (
	"apps/ch.4.4/nonce"
	"apps/ch.4.4/validator"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT     = "9090"
	HOST_URL = "http://localhost:" + PORT
)

var submissions nonce.Nonces
var t *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, HOST_URL+"/profile", http.StatusTemporaryRedirect)
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "profile", submissions.NewNonce())
}
func checkProfile(w http.ResponseWriter, r *http.Request) {
	var errs validator.Errors
	r.ParseForm()
	token := r.Form.Get("token")
	if err := submissions.CheckThenMarkToken(token); err != nil {
		errs = validator.Errors{[]error{err}}
	} else {
		p := validator.ProfilePage{&r.Form}
		errs = p.GetErrors()
	}
	t.ExecuteTemplate(w, "submission", errs)
}
func init() {
	submissions = nonce.New()
	t = template.Must(template.ParseFiles("profile.gtpl", "submission.gtpl"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/checkprofile", checkProfile)

	err := http.ListenAndServe(":"+PORT, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
