package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Seamus")
	want := "Hello, Seamus"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
