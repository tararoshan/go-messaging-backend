package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

    // use default server multiplexer
	err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed. Typically expected.\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(writer http.ResponseWriter, request_ptr *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(writer, "Here's a sample root response!\n")
}

func getHello(writer http.ResponseWriter, request_ptr *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(writer, "Hello response!\n")
}