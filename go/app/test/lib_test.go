package main

import (
	"app/lib"
	"testing"
)

func TestTruncate(t *testing.T) {
	text := "this text contains exactly 45 chracters."
	maxLen := 30
	r := lib.Truncate(text, maxLen)
	if len(r) != text[:maxLen] {
		t.Fatal("failed test")
	}
}
