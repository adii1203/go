package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aditya")
	want := "Hello, Aditya"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
