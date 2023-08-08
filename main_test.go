package main

import "testing"

func helloWorld() string {
	return "Hello World!!"
}

func TestHelloWorld(t *testing.T) {
	if helloWorld() != "Hello World!!" {
		t.Fatal("Test fail")
	}
}
