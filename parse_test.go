package main

import (
	"testing"
)

func TestUnpackH(t *testing.T) {
	text, rest, err := unpackH("e69716f637e2d616e6")
	if err != nil {
		t.Fatalf("%s/%s: %s\n", text, rest, err.Error())
		return
	}
	if text != "nyaos.man" || rest != "" {
		t.Fatalf("expect nyaos.man but '%s' and rest is '%s'\n", text, rest)
	}
	// println(text, rest)
}
