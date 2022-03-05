package main

import (
	"app/lib"
	"testing"
)

func TestTruncate(t *testing.T) {
	text := "this text contains exactly 45 chracters."
	maxLen := 30
	trailChar := "..."
	testText := text[:maxLen] + trailChar
	r := lib.Truncate(text, maxLen)
	if r != testText {
		t.Fatal("failed test")
	}
}
