package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Fatalf("Add(2, 3) = %d, expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(4, 2)
	expected := 2
	if result != expected {
		t.Fatalf("Subtract(4, 2) = %d, expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 4)
	expected := 8
	if result != expected {
		t.Fatalf("Multiply(2, 4) = %d, expected %d", result, expected)
	}
}
