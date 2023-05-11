package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Use the signature of the handler to register on server routes
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

    // "nil" means use default server multiplexer
	err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed. Typical behavior.\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(writer http.ResponseWriter, request_ptr *http.Request) {
	fmt.Printf("got / request, parsing as POST\n")

	// TODO parse POST request
	// need to store sender, reciever, and message
	io.WriteString(writer, "sent!\n")
}

func getHello(writer http.ResponseWriter, request_ptr *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(writer, "Hello response!\n")
}
