package groupanagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected [][]string
	}
	tests := []test{
		/*
			{
				name:  "test 1",
				input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
				expected: [][]string{
					{"bat"},
					{"nat", "tan"},
					{"ate", "eat", "tea"},
				},
			},
		*/
		{
			name:     "test 2",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "test 3",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "test 4",
			input:    []string{"", "b", ""},
			expected: [][]string{{"", ""}, {"b"}},
		},
		{
			name:     "test 5",
			input:    []string{"ac", "c"},
			expected: [][]string{{"c"}, {"ac"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := groupAnagrams(tt.input)
			if len(actual) != len(tt.expected) {
				t.Logf("Actual len %d differs from expected len %d", len(actual), len(tt.expected))
				t.FailNow()
			}
			for i, v1 := range actual {
				for j, v2 := range v1 {
					if v2 != tt.expected[i][j] {
						t.Errorf("Actual: %s; Expected: %s; Position: %d %d", v2, tt.expected[i][j], i, j)
					}
				}
			}
		})
	}
}
