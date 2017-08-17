package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	t.Parallel()
	expected := []int{1}
	result := sortF([]int{1})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase2(t *testing.T) {
	t.Parallel()
	expected := []int{1, 2}
	result := sortF([]int{1, 2})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

func TestCase3(t *testing.T) {
	t.Parallel()
	expected := []int{1, 2, 3, 4}
	result := sortF([]int{4, 3, 2, 1})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}
