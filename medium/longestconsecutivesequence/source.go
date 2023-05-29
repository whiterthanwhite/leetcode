package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}

	mapNums := make(map[int]int)
	uf := &WeightedQuickUnionUF{
		id:     make([]int, n),
		sz:     make([]int, n),
		maxSeq: 1,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}

	for i := 0; i < n; i++ {
		if _, ok := mapNums[nums[i]]; ok {
			continue
		}
		if pos, ok := mapNums[nums[i]+1]; ok {
			uf.Union(i, pos)
		}
		if pos, ok := mapNums[nums[i]-1]; ok {
			uf.Union(i, pos)
		}
		mapNums[nums[i]] = i
	}

	return uf.maxSeq
}

type WeightedQuickUnionUF struct {
	id     []int
	sz     []int
	maxSeq int
}

func (uf *WeightedQuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *WeightedQuickUnionUF) Union(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
		if uf.sz[j] > uf.maxSeq {
			uf.maxSeq = uf.sz[j]
		}
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
		if uf.sz[i] > uf.maxSeq {
			uf.maxSeq = uf.sz[i]
		}
	}
}
