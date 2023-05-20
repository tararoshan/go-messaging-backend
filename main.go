package main

import (
	"errors"
	"fmt"
	// "io"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Defined in messagemap.go
	messagemap := makeMessageMap()

	// Use the signature of the handler to register on server routes
	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		postRoot(writer, req, messagemap)
	}).Methods("POST")
	router.HandleFunc("/{userNameA}/{userNameB}/{fromTimeStamp}",
		func(writer http.ResponseWriter, req *http.Request) {
			getPeopleTime(writer, req, messagemap)
		}).Methods("GET")
	router.HandleFunc("/{userNameA}/{userNameB}",
		func(writer http.ResponseWriter, req *http.Request) {
			getPeopleTime(writer, req, messagemap)
		}).Methods("GET")

    // "nil" means use default server multiplexer
	srv := &http.Server{
        Addr:         ":3333",
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: router,
    }

	err := srv.ListenAndServe()
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed. Typical behavior.\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func postRoot(writer http.ResponseWriter, request_ptr *http.Request, messagemap *MessageMap) {
	fmt.Printf("heard / request, parsing as POST\n")
	// fmt.Printf("request: %s\n", request_ptr)
	// io.WriteString(writer, "sent!\n")

	// Parse POST request and add timestamp
	var me MessageEntry
	err := json.NewDecoder(request_ptr.Body).Decode(&me)
	if err != nil {
		fmt.Printf("logging parsing error: %s\n", err)
		http.Error(writer, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	me.At = time.Now().Unix()

	// Access the parsed values
	fmt.Printf("Decoded message: %+v\n", me)
	fmt.Printf("From: %s\n", me.From)
	fmt.Printf("To: %s\n", me.To)
	fmt.Printf("Message: %s\n", me.Message)
	fmt.Printf("Timestamp: %d\n", me.At)

	// TODO need to store sender, reciever, and message in HashMap by peoplepair
	messagemap.enterMessage(me)
	return
}

func getPeopleTime(writer http.ResponseWriter, request_ptr *http.Request, messagemap *MessageMap) {
	fmt.Printf("heard /userA/userB/time request, responding as GET\n")
	// io.WriteString(writer, "GETting your data!\n")

	routeVars := mux.Vars(request_ptr)
	if routeVars["fromTimeStamp"] == "" {
		fmt.Printf("got no fromTimeStamp\n")
		routeVars["fromTimeStamp"] = "0"
	}
	// fmt.Printf("request: %s\n", request_ptr)
	fmt.Printf("userNameA: %s, userNameB: %s, fromTimeStamp: %s\n", routeVars["userNameA"], routeVars["userNameB"], routeVars["fromTimeStamp"])

	fmt.Println()

	timestamp, err := strconv.ParseInt(routeVars["fromTimeStamp"], 10, 64)
	if err != nil {
		fmt.Printf("Error parsing timestamp from message. Error: %s", err)
	}
	// Content-Type header taken care of in method
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	messagemap.printPeopleMessagesAfterTimestamp(getPeoplePair(routeVars["userNameA"], routeVars["userNameB"]), timestamp, writer)
	return
}
