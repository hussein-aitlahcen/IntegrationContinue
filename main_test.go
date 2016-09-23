package main

import "testing"

func TestYo(t *testing.T) {
	expected := "Yo"
	obtained := GetString()
	if expected != obtained {
		t.Fatalf("GetString failed exp: %s, obt: %s\n", expected, obtained)
	}
}
