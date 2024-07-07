package main

import (
	"io"
	"strings"
	"testing"
)

func TestByteCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"Hello, World!", 13},
		{"こんにちは", 15}, // 5 characters, but 15 bytes in UTF-8
		{"1234567890", 10},
		{"Line 1\nLine 2\nLine 3", 20},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		data, err := io.ReadAll(reader)
		if err != nil {
			t.Fatal(err)
		}

		result, err := countBytes(data)
		if err != nil {
			t.Fatalf("Error counting bytes: %v", err)
		}
		if result != tc.expected {
			t.Errorf("countBytes(%q) = %d; Expected %d", tc.input, tc.expected, result)
		}
	}

}

func TestWordCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"OneWord", 1},
		{"Hello, World!", 2},
		{"This is a test.", 4},
		{"   Spaces   between   words   ", 3},
		{"Line 1\nLine 2\nLine 3", 6},
		{"Hyphenated-word", 1},
		{"Multiple   spaces   between", 3},
		{"End with space ", 3},
		{" Start with space", 3},
		{"こんにちは世界", 7}, //function handles UTF-8 (each character as one word)
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		data, err := io.ReadAll(reader)
		if err != nil {
			t.Fatal(err)
		}

		results, err := countWords(data)
		if err != nil {
			t.Fatalf("Error counting words: %v", err)
		}

		if results != tc.expected {
			t.Errorf("countWords(%q) = %d; Expected %d", tc.input, results, tc.expected)
		}

	}
}

func TestLineCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"One line", 1},
		{"Line 1\nLine 2\nLine 3", 3},
		{"\n\n\n", 3},
		{"No newline at end", 1},
		{"Newline at end\n", 1},
		{"Multiple\n\nEmpty\n\nLines", 5},
		{"こんにちは\n世界\n", 2},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		data, err := io.ReadAll(reader)
		if err != nil {
			t.Fatal(err)
		}

		results, err := countLines(data)
		if err != nil {
			t.Fatalf("Error counting lines: %v", err)
		}

		if results != tc.expected {
			t.Errorf("countLines(%q) = %d; Expected %d", tc.input, results, tc.expected)
		}
	}
}
