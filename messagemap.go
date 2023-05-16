package main

import (
	// "net/http"
	"sync"
)

// Used for storing entries in the arrays of the message map.
type MessageEntry struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	At      int64  `json:"at"`
}

// Used as the keys in the message map. Names in sorted order.
type PeoplePair struct {
	NameA	string
	NameB	string
}

// Designed to deterministically handle concurrent requests. Write preferring.
type MessageMap struct {
	Data                 map[PeoplePair][]MessageEntry
	Mu                   sync.Mutex
	CondVar              sync.Cond
	WriterActive         bool
	NumWritersWaiting    int
	NumReadersActive     int
}

// Initializes internal map, number of reader/writers, map lock
func makeMessageMap() *MessageMap {
	messagemap := &MessageMap{}
	m := make(map[PeoplePair][]MessageEntry)
	messagemap.Data = m
	messagemap.WriterActive = false
	messagemap.NumWritersWaiting = 0
	messagemap.NumReadersActive = 0
	messagemap.CondVar = *sync.NewCond(&messagemap.Mu)
	return messagemap
}

// func (messagemap *MessageMap) enterMessage(message MessageEntry) {

// }

// func (messagemap *MessageMap) getPeopleMessagesAfterTimestamp(pp PeoplePair, timestamp int64) []MessageEntry {

// }

// // Save time by not copying the array twice
// func (messagemap *MessageMap) printPeopleMessagesAfterTimestamp(pp PeoplePair, timestamp int64, writer http.ResponseWriter) {

// }
