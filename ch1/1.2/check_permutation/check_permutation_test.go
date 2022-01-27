package check_permutation_test

import (
	"fmt"
	"testing"

	"github.com/dchest/uniuri"
	cp "github.com/tcc-sejohnson/cracking-the-coding-interview-e6/ch1/1.2/check_permutation"
)

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		FirstInput  string
		SecondInput string
		Want        bool
	}{
		{FirstInput: "asdf", SecondInput: "sadf", Want: true},
		{FirstInput: "", SecondInput: "", Want: true},
		{FirstInput: "tacocat", SecondInput: "tacocat", Want: true},
		{FirstInput: "This is not a permutation.", SecondInput: "This won't match, man.", Want: false},
		{FirstInput: "ttt", SecondInput: "tttt", Want: false},
		{FirstInput: "tttt", SecondInput: "ttt", Want: false}, // make sure we handle out-of-bounds
	}

	for _, test := range tests {
		got := cp.IsPermutation(test.FirstInput, test.SecondInput)
		if test.Want != got {
			t.Errorf("%v and %v should have produced result %v, but it did not", test.FirstInput, test.SecondInput, test.Want)
		}
	}
}

func BenchmarkCheckPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cp.IsPermutation("the quick brown fox jumped over the lazy dog.", "e epguiy z uotrkjb  chem laervtw  oxfoddqonh.")
	}
}

func BenchmarkCheckPermutationOnProgressivelyLongerTrueStrings(b *testing.B) {
	for i := 10; i < 1000000; i = i * 10 {
		testString := uniuri.NewLen(i)
		b.Run(fmt.Sprintf("MachingStringsOfLength%v", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cp.IsPermutation(testString, testString)
			}
		})
	}
}

func BenchmarkCheckPermutationOnProgressivelyLongerFalseStrings(b *testing.B) {
	for i := 10; i < 1000000; i = i * 10 {
		testString := uniuri.NewLen(i)
		testNonmatchingString := uniuri.NewLen(i)
		b.Run(fmt.Sprintf("MachingStringsOfLength%v", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cp.IsPermutation(testString, testNonmatchingString)
			}
		})
	}
}
