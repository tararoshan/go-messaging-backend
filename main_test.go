// Testing file for main.go
package main

import (
	"bytes"
	"time"
	"net/http/httptest"
	"testing"
)

/**
 * BASIC TESTS - named with the function they test after underscore.
 * NOTE: use net/http/httptest.
 */
// Check for appropriate content headers and contents
func TestPostRoot(t *testing.T) {
	mockbody := []byte(`{"from": "zack", "to": "charles", "message": "pizza tonight?"}`)
	request := httptest.NewRequest("GET", "/", bytes.NewBuffer(mockbody))
	recorder := httptest.NewRecorder()
	postRoot(recorder, request)

	messages := messagemap.getPeopleMessagesAfterTimestamp("charleszack", time.Now().Add(-5 * time.Second).Unix())

	expectedLen := 1
	if len(messages) != expectedLen  {
		t.Errorf("Wrong number of messages. Expected %d but found %d", expectedLen, len(messages))
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
func TestGetPeopleTime(t *testing.T) {}
 
/**
 * EXTRA TESTS - tests I thought would be useful for coverage or to check some
 * synchronization cases. Also added a test if I noticed an error, so that it
 * doesn't go unnoticed later. NOTE: use "go" before a function to make a new
 * goroutine/thread of execution.
 */
func TestWriteReadSynchronization(t *testing.T) {}

func TestWriteReadTwiceSynchronization(t *testing.T) {}
