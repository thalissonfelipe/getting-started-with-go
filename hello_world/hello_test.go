package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Felipe")
	expect := "Hello, Felipe"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}
