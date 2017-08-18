package main

import "testing"

func TestAdd(t *testing.T) {
	u := User{}
	result := u.add()
	if !result {
		t.Fatalf("Error")
	}
}
