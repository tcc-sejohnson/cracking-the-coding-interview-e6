package is_unique_test

import (
	"testing"

	"github.com/tcc-sejohnson/cracking-the-coding-interview-e6/ch1/1.1/is_unique"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		Unique bool
		Input  string
	}{
		{Unique: true, Input: "abcdefghijklmnopqrstuvwxyz"},
		{Unique: false, Input: "asdfqwera"},
		{Unique: false, Input: "the quick brown fox jumps over the lazy dog."},
		{Unique: true, Input: "quixote"},
		{Unique: false, Input: ""},
	}

	for _, test := range tests {
		unique := is_unique.IsUnique(test.Input)
		if test.Unique != unique {
			t.Errorf("expected input %v to return a %v result", test.Input, test.Unique)
		}
	}
}

func BenchmarkTestIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		is_unique.IsUnique("the quick brown fox jumps over the lazy dog.")
	}
}
