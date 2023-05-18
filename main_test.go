// Testing file for main.go
package main

import (
	"fmt"
	"testing"
)

/**
 * BASIC TESTS - named with the function they test after underscore. NOTE: use net/http/httptest.
 */
func Test_postRoot(t *testing.T) {}

func Test_getPeopleTime(t *testing.T) {}
 
/**
 * EXTRA TESTS - tests I thought would be useful for coverage or to check some
 * synchronization cases. Also added a test if I noticed an error, so that it
 * doesn't go unnoticed later.
 */
func TestWriteReadSynchronization(t *testing.T) {}

func TestWriteReadTwiceSynchronization(t *testing.T) {}
