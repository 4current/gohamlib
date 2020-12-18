package main

import "testing"

func TestModes(t *testing.T) {
	var m Mode = AM
	if m.String() != "AM" {
		t.Error("AM.String() should equal \"AM\"")
	}

}
