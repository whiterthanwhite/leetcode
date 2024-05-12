package maximumnumberofvowelsinasubstringofgivenlength

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	vowels := make(map[byte]struct{}, 5)
	vowels['a'] = struct{}{}
	vowels['e'] = struct{}{}
	vowels['i'] = struct{}{}
	vowels['o'] = struct{}{}
	vowels['u'] = struct{}{}

	N := len(s)
	max := 0
	for i := 0; i < N; i += k {
		m := 0
		t := ""
		if i+k > N {
			t = s[N-k : N]
		} else {
			t = s[i : i+k]
		}
		for j := 0; j < k; j++ {
			_, ok := vowels[t[j]]
			if ok {
				m++
			}
		}
		fmt.Println(t, m)
		if max < m {
			max = m
		}
	}

	return max
}
