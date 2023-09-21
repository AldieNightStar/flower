package util

import "testing"

func TestSliceOf(t *testing.T) {
	s := SliceOf("fabcx", 1, 4)
	if s != "abc" {
		t.Fatalf("Should be abc: %s", s)
	}

	s = SliceOf("x", 0, 1000)
	if s != "x" {
		t.Fatalf("Should be x: %s", s)
	}

	s = SliceOf("", 0, 1)
	if s != "" {
		t.Fatalf("Should be empty: %s", s)
	}

	s = SliceOf("123", 0, 50)
	if s != "123" {
		t.Fatalf("Should be 123: %s", s)
	}

	s = SliceOf("99", 5, 10)
	if s != "" {
		t.Fatalf("Should be empty: %s", s)
	}

	s = SliceOf("123", -4, 100)
	if s != "123" {
		t.Fatalf("Should be 123: %s", s)
	}
}
