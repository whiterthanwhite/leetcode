package countpairsofsimilarstrings

func similarPairs(words []string) int {
	pairs := 0
	newWords := make([]string, len(words))
	n := 0
	for _, word := range words {
		c := &charST{
			keys: make([]rune, len(word)),
		}
		for _, r := range word {
			c.put(r)
		}
		newWords[n] = c.getString()
		n++
	}
	var w1, w2 string
	n = len(newWords) - 1
	for i := 0; i < n; i++ {
		w1 = newWords[i]
		for j := n; j > i; j-- {
			w2 = newWords[j]
			if w1 == w2 {
				pairs++
			}
		}
	}

	return pairs
}

type charST struct {
	keys []rune
	n    int
}

func (c *charST) getString() string {
	if c.n == 0 {
		return ""
	}
	return string(c.keys[:c.n])
}

func (c *charST) put(key rune) {
	i := c.rank(key, 0, c.n-1)
	if i < c.n && c.keys[i] == key {
		return
	}
	for j := c.n; j > i; j-- {
		c.keys[j] = c.keys[j-1]
	}
	c.keys[i] = key
	c.n++
}

func (c *charST) rank(key rune, lo, hi int) int {
	if hi < lo {
		return lo
	}
	mid := lo + (hi-lo)/2
	if key < c.keys[mid] {
		return c.rank(key, lo, mid-1)
	} else if key > c.keys[mid] {
		return c.rank(key, mid+1, hi)
	} else {
		return mid
	}
}
