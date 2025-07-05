package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	})
	server("localhost", "3000")
}

func server(addr string, port string) {
	log.Fatal(http.ListenAndServe(addr+":"+port, nil))
}
