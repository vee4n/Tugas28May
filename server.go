package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing TLS\n")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Listen and serve with TLS
	err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
	if err != nil {
		panic(err)
	}
}
