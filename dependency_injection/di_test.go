package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kevin")

	got := buffer.String()
	want := "Hello, Kevin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
