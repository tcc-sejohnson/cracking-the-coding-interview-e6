package urlify_clever_test

import (
	"testing"

	urlify "github.com/tcc-sejohnson/cracking-the-coding-interview-e6/ch1/1.3/urlify_clever"
)

func TestUrlify(t *testing.T) {
	tests := []struct {
		Input  string
		Length int
		Want   string
	}{
		{Input: "Mr John Smith    ", Length: 13, Want: "Mr%20John%20Smith"},
		{Input: "             ", Length: 0, Want: ""},
		{Input: "The quick brown fox jumped over the lazy dog.                ", Length: 45, Want: "The%20quick%20brown%20fox%20jumped%20over%20the%20lazy%20dog."},
	}

	for _, test := range tests {
		got := urlify.Urlify(test.Input, test.Length)
		if test.Want != got {
			t.Errorf("failed to urlify. Input: %v, want: %v, got: %v", test.Input, test.Want, got)
		}
	}
}

func BenchmarkUrlify(b *testing.B) {
	input := "Mr John Smith    "
	for i := 0; i < b.N; i++ {
		urlify.Urlify(input, 13)
	}
}
