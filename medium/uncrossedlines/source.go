package uncrossedlines

import "fmt"

func maxUncrossedLines(nums1, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	s1 := make([]bag, n1)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if nums1[i] == nums2[j] {
				s1[i].add(j)
			}

		}
	}
	for i, v := range s1 {
		fmt.Print(i, ": ")
		items := v.items()
		for j := len(items) - 1; j >= 0; j-- {
			fmt.Print(items[j], " ")
		}
		fmt.Println()
	}
	return -1
}

type digraph struct {
	v   int
	e   int
	adj []bag
}

func New(v int) *digraph {
	d := &digraph{
		v:   v,
		e:   0,
		adj: make([]bag, v),
	}
	for i := range d.adj {
		d.adj[i] = bag{}
	}
	return d
}

func (d *digraph) addEdge(v, w int) {
	d.adj[v].add(w)
	d.e++
}

type node struct {
	item int
	next *node
}

type bag struct {
	first *node
}

func (b *bag) add(item int) {
	oldFirst := b.first
	b.first = &node{
		item: item,
		next: oldFirst,
	}
}

func (b *bag) items() []int {
	r := make([]int, 0)
	for n := b.first; n != nil; n = n.next {
		r = append(r, n.item)
	}
	return r
}
