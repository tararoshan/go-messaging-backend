package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
)

// Used for storing entries in the arrays of the message map.
type MessageEntry struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	At      int64  `json:"at"`
}

// Designed to deterministically handle concurrent requests. Write preferring.
type MessageMap struct {
	Data                 map[string][]MessageEntry
	Mu                   sync.RWMutex
}

type ByAt []MessageEntry

// Allow the slices in the map (stored as values) to be binary-searchable, sorted by time (At field)
func (slice ByAt) Len() int           { return len(slice) }
func (slice ByAt) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }
func (slice ByAt) Less(i, j int) bool { return slice[i].At < slice[j].At }

// Initializes internal map, number of reader/writers, map lock
func makeMessageMap() *MessageMap {
	messagemap := &MessageMap{}
	m := make(map[string][]MessageEntry)
	messagemap.Data = m
	return messagemap
}

func getPeoplePair(nameA string, nameB string) string {
	// Determine the PeoplePair to use, for consistency
	result := nameA + nameB 
	if nameB > nameA {
		result = nameB + nameA
	}
	return result
}

func (messagemap *MessageMap) enterMessage(message MessageEntry) {
	alphabeticalpair := getPeoplePair(message.To, message.From)

	messagemap.Mu.Lock()
	defer messagemap.Mu.Unlock()

	// Add the key-value to the map or append to slice
	// Note: I assume the messages are given lock priority in the order they
	// were sent, so it's okay to add to the end!
	if messagesplice, ok := messagemap.Data[alphabeticalpair]; ok {
		messagemap.Data[alphabeticalpair] = append(messagesplice, message)
	} else {
		// Make the slice
		size := 1
		capacity := 4
		messagemap.Data[alphabeticalpair] = make([]MessageEntry, size, capacity)
		messagemap.Data[alphabeticalpair][0] = message
	}
}

func (messagemap *MessageMap) getPeopleMessagesAfterTimestamp(peoplepair string, timestamp int64) []MessageEntry {
	messagemap.Mu.RLock()
	defer messagemap.Mu.RUnlock()

	messageslice := messagemap.Data[peoplepair]
	// Find the index
	startindex := sort.Search(len(messageslice), func(i int) bool {
		return messageslice[i].At >= timestamp
	})
	// Return the slice of the slice
	return messageslice[startindex:]
}

// Save time by not copying the array twice when you return to caller
func (messagemap *MessageMap) printPeopleMessagesAfterTimestamp(peoplepair string, timestamp int64, writer http.ResponseWriter) {
	messagemap.Mu.RLock()
	defer messagemap.Mu.RUnlock()

	messageslice := messagemap.Data[peoplepair]
	// Find the index
	startindex := sort.Search(len(messageslice), func(i int) bool {
		return messageslice[i].At >= timestamp
	})

	// Write the slice of the slice as JSON
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(messageslice[startindex:])
	if err != nil {
		fmt.Println("Error when encoding: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
