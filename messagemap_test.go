// Testing file for messagemap.go. Unfinished but done.
package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * BASIC TESTS - named with the function they test after underscore.
 */
func Test_makeMessageMap(t *testing.T) {}

func Test_getPeoplePair(t *testing.T) {}

func Test_enterMessage(t *testing.T) {}

func Test_getPeopleMessagesAfterTimestamp(t *testing.T) {}

func Test_printPeopleMessagesAfterTimestamp(t *testing.T) {}

/**
 * EXTRA TESTS - tests I thought would be useful for coverage or to check some
 * synchronization cases. Also added a test if I noticed an error, so that it
 * doesn't go unnoticed later. NOTE: use "go" before a function to make a new
 * goroutine/thread of execution.
 */
func TestMultipleWrites(t *testing.T) {
	messagemap := makeMessageMap()
	expectedMessages := []MessageEntry{
		{From: "zack", To: "charles", Message: "pizza tonight?", At: time.Now().Unix()},
		{From: "charles", To: "zack", Message: "yeah, def.", At: time.Now().Unix()},
		{From: "claire", To: "zara", Message: "wanna get coffee tmrw?", At: time.Now().Unix()},
		{From: "zara", To: "claire", Message: "busy, what about the day after?", At: time.Now().Unix()},
	}

	var wg sync.WaitGroup
	wg.Add(len(expectedMessages))

	for i := range expectedMessages {
		go func(index int) {
			defer wg.Done()
			messagemap.enterMessage(expectedMessages[index])
			fmt.Printf("message: %s\n", expectedMessages[index].Message)
		}(i)
	}
	wg.Wait()

	// Ensure routines didn't write to the same place
	numZaraClair := len(messagemap.Data[getPeoplePair("zara", "claire")])
	expectedNumZaraClaire := 2
	if numZaraClair != expectedNumZaraClaire {
		t.Errorf("Expected %d messages between Zara and Claire, but found %d.", expectedNumZaraClaire, numZaraClair)
	}

	numZackCharles := len(messagemap.Data[getPeoplePair("zack", "charles")])
	expectednumZackCharlese := 2
	if numZackCharles != expectednumZackCharlese {
		t.Errorf("Expected %d messages between Zack and Charles, but found %d.", expectednumZackCharlese, numZackCharles)
	}
}
 
func TestSortedMessageMap(t *testing.T) {}
