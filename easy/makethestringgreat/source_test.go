package makethestringgreat

import "testing"

func TestMakeTheStringGreat(t *testing.T) {
	type test struct {
		src, wanted string
	}
	tests := []test{
		{
			src:    "leEeetcode",
			wanted: "leetcode",
		},
		{
			src:    "abBAcC",
			wanted: "",
		},
		{
			src:    "s",
			wanted: "s",
		},
		{
			src:    "mC",
			wanted: "mC",
		},
		{
			src:    "Pp",
			wanted: "",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := makeGood(tt.src)
			if actual != tt.wanted {
				t.Errorf("Actual: %s; but Wanted: %s", actual, tt.wanted)
			}
		})
	}
}
