// Testing file for main.go
package main

import (
	"bytes"
	// "fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"github.com/gorilla/mux"
)

/**
 * BASIC TESTS - named with the function they test after underscore.
 * NOTE: use net/http/httptest.
 */
func TestPostRoot(t *testing.T) {
	// Setup for test
	messagemap := makeMessageMap()
	body := []byte(`{"from": "zack", "to": "charles", "message": "pizza tonight?"}`)

	request := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	postRoot(recorder, request, messagemap)

	pp := getPeoplePair("zack", "charles")
	messages := messagemap.getPeopleMessagesAfterTimestamp(pp, 0)

	expectedLen := 1
	if len(messages) != expectedLen  {
		// Use fatal since we could get an out-of-bounds issue later
		t.Fatalf("Wrong number of messages. Expected %d but found %d", expectedLen, len(messages))
	}

	expectedMessage := MessageEntry{
		From: 		"zack",
		To: 		"charles",
		Message: 	"pizza tonight?",
		// Not sure what time the message will be recorded at
		At:			messages[0].At,
	}
	if messages[0] != expectedMessage {
		t.Errorf("The message wasn't stored correctly. Expected %+v but found %+v", expectedMessage, messages[0])
	}
}

// Check for appropriate content headers and contents
func TestGetPeopleTime(t *testing.T) {
	// Setup for tests
	router := mux.NewRouter()
	messagemap := makeMessageMap()

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

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	var requestBodies = []struct {
		body []byte
	}{
		{[]byte(`{"from": "zack", "to": "charles", "message": "pizza tonight?"}`)},
		{[]byte(`{"from": "charles", "to": "zack", "message": "yeah, def."}`)},
		{[]byte(`{"from": "claire", "to": "zara", "message": "wanna get coffee tmrw?"}`)},
		{[]byte(`{"from": "zara", "to": "claire", "message": "busy, what about the day after?"}`)},
	}

	var lastMessageTime int64
	for _, tt := range requestBodies {
		time.Sleep(1 * time.Second)
		lastMessageTime = time.Now().Unix()
		_, err := http.Post(testServer.URL, "application/json", bytes.NewBuffer(tt.body))
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
	}

	// Check just zack and charles's messages
	response, err := http.Get(testServer.URL + "/zack/charles")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	// Check response Content-Type header
	expectedContentType := "application/json" 
	if response.Header.Get("Content-Type") != expectedContentType {
		t.Errorf("Incorrect header. Expected %s, found %s", expectedContentType, response.Header.Get("Content-Type"))
	}

	// Check for just two messages with the proper values
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	var messages []MessageEntry
	err = json.Unmarshal(body, &messages)
	if err != nil {
		t.Fatalf("Issue unmarshalling response body: %v", err)
	}
	expectedMessages := []MessageEntry{
		{From: "zack", To: "charles", Message: "pizza tonight?", At: messages[0].At},
		{From: "charles", To: "zack", Message: "yeah, def.", At: messages[1].At},
	}
	
	for i, expected := range expectedMessages {
		if messages[i] != expected {
			t.Errorf("The message at index %d wasn't stored correctly. Expected %v but found %v", i, expected, messages[i])
		}
	}

	// Check just the last message from the girls
	base := 10
	response, err = http.Get(testServer.URL + "/zara/claire/" +  strconv.FormatInt(lastMessageTime, base))
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		t.Fatalf("Issue unmarshalling response body: %v", err)
	}
	expectedMessages = []MessageEntry{
		{From: "zara", To: "claire", Message: "busy, what about the day after?", At: messages[0].At},
	}
	
	for i, expected := range expectedMessages {
		if messages[i] != expected {
			t.Errorf("The message at index %d wasn't stored correctly. Expected %v but found %v", i, expected, messages[i])
		}
	}

	// Check for no messages from the boys after lastMessageTime
	response, err = http.Get(testServer.URL + "/zack/charles/" +  strconv.FormatInt(lastMessageTime, base))
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		t.Fatalf("Issue unmarshalling response body: %v", err)
	}
	
	if len(messages) != 0 {
		t.Errorf("There was a message found between zack and charles after lastMessageTime: %v", messages)
	}
}
 
/**
 * EXTRA TESTS - tests I thought would be useful for coverage or to check some
 * synchronization cases. Also added a test if I noticed an error, so that it
 * doesn't go unnoticed later. NOTE: use "go" before a function to make a new
 * goroutine/thread of execution.
 */
func TestWriteReadSynchronization(t *testing.T) {}

func TestWriteReadTwiceSynchronization(t *testing.T) {}
