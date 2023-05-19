// Testing file for main.go
package main_test

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

/**
 * BASIC TESTS - named with the function they test after underscore. NOTE: use net/http/httptest.
 */
// Check for appropriate content headers and contents
func Test_postRoot(t *testing.T) {}

// Check for appropriate content headers and contents
func Test_getPeopleTime(t *testing.T) {}
 
/**
 * EXTRA TESTS - tests I thought would be useful for coverage or to check some
 * synchronization cases. Also added a test if I noticed an error, so that it
 * doesn't go unnoticed later. NOTE: use "go" before a function to make a new
 * goroutine/thread of execution.
 */
func TestWriteReadSynchronization(t *testing.T) {}

func TestWriteReadTwiceSynchronization(t *testing.T) {}
