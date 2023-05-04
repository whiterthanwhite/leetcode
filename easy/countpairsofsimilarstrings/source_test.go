package countpairsofsimilarstrings

import (
	"fmt"
	"testing"
)

func TestSimilarPairs(t *testing.T) {
	type test struct {
		paremeter []string
		output    int
	}
	tests := []test{
		{
			paremeter: []string{"aba", "aabb", "abcd", "bac", "aabc"},
			output:    2,
		},
		{
			paremeter: []string{"aabb", "ab", "ba"},
			output:    3,
		},
		{
			paremeter: []string{"nba", "cba", "dba"},
			output:    0,
		},
	}
	for tn, ct := range tests {
		t.Run(fmt.Sprintf("test %d", tn), func(t *testing.T) {
			actual := similarPairs(ct.paremeter)
			t.Log(actual)
			if actual != ct.output {
				t.Fail()
			}
		})
	}
}
