package main

import "testing"

func TestHelloMessage(t *testing.T) {
	if HelloMessage() != "Hello" {
		t.Errorf("hellomessage() = %s; expected Hello", HelloMessage())
	}
}
