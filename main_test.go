package main

import "testing"

var expectedMessage string = "Hello"

func TestHelloMessage(t *testing.T) {
	if HelloMessage() != expectedMessage {
		t.Error("got ", HelloMessage(), " expected ", expectedMessage)
	}
}
