package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from Server\n"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Healthy from Server\n"))
	})

	http.ListenAndServe(":8080",nil)

}
