package main

import (
	"fmt"
	"net/http"
)

var USERNAME = "hah"
var PASSWORD = "wah"

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe("localhost:8083", nil)
}

func root(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("basic auth failed"))
		fmt.Println("basic auth failed")
		return
	}

	if username != USERNAME || password != PASSWORD {
		w.Write([]byte("username/password is wrong"))
		fmt.Println("username/password is wrong")
		return
	}

	w.Write([]byte("Hello World"))
}
