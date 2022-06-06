package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	event := Event{Path: "/health"}
	response, err := Handler(nil, event)
	if err != nil {
		t.Errorf("Unexpected error: %e", err)
	}

	expect := "Received request from: /health"
	if response != expect {
		t.Errorf("Unexpecte response from handler:\n\tgot: \"%s\"\n\texpected: \"%s\"", response, expect)
	}
}
