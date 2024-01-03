package main

import (
	"testing"
	"time"
)


func TestSetTimeValidInput(t *testing.T) {
	testTime := "15:04"
	_, err := time.Parse("15:04", testTime)
	if err != nil {
		t.Errorf("Parsing time failed for valid input %s", testTime)
	}
}

func TestSetTimeInvalidInput(t *testing.T) {
	testTime := "invalid-time"
	_, err := time.Parse("15:04", testTime)
	if err == nil {
		t.Errorf("Expected an error for invalid time input %s", testTime)
	}
}
