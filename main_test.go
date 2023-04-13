package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		input        string
		expectResult bool
	}

	testCases := []Case{
		{
			input:        "ana",
			expectResult: true,
		},
		{
			input:        "mamam",
			expectResult: true,
		},
		{
			input:        "leonardo",
			expectResult: false,
		},
		{
			input:        "",
			expectResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.expectResult {
			t.Error("FAIL", testCase)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	type Case struct {
		input        string
		expectResult bool
	}

	testCases := []Case{
		{
			input:        "ana",
			expectResult: true,
		},
		{
			input:        "mamam",
			expectResult: true,
		},
		{
			input:        "leonardo",
			expectResult: false,
		},
		{
			input:        "",
			expectResult: false,
		},
	}

	for _, testCase := range testCases {
		if isPalindrome2(testCase.input) != testCase.expectResult {
			t.Error("FAIL", testCase)
		}
	}
}

func TestInvertText(t *testing.T) {
	type Case struct {
		input        string
		expectResult string
	}

	testCases := []Case{
		{
			input:        "ana",
			expectResult: "ana",
		},
		{
			input:        "mamam",
			expectResult: "mamam",
		},
		{
			input:        "leonardo",
			expectResult: "odranoel",
		},
	}

	for _, testCase := range testCases {
		if invertText(testCase.input) != testCase.expectResult {
			t.Error("FAIL", testCase)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("LEONARDO")
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2("LEONARDO")
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome3("LEONARDO")
	}
}
