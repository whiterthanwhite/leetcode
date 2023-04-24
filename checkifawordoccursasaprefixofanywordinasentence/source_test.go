package checkifawordoccursasaprefixofanywordinasentence

import "testing"

func TestIsPrefixOfWord(t *testing.T) {
	type test struct {
		name       string
		sentence   string
		searchWord string
		expected   int
	}
	tests := []test{
		{
			name:       "test 1",
			sentence:   "i love eating burger",
			searchWord: "burg",
			expected:   4,
		},
		{
			name:       "test 2",
			sentence:   "this problem is an easy problem",
			searchWord: "pro",
			expected:   2,
		},
		{
			name:       "test 3",
			sentence:   "i am tired",
			searchWord: "you",
			expected:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isPrefixOfWord(tt.sentence, tt.searchWord)
			if actual != tt.expected {
				t.Errorf("Actual: %d; expected: %d", actual, tt.expected)
			}
		})
	}
}
