package insertationoftwoarrays

import "fmt"

func intersection(nums1, nums2 []int) []int {
	s1 := make(map[int]struct{})
	s2 := make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		if _, ok := s1[nums1[i]]; !ok {
			s1[nums1[i]] = struct{}{}
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := s2[nums2[i]]; !ok {
			s2[nums2[i]] = struct{}{}
		}
	}

	r := make([]int, 0)
	for k := range s1 {
		if _, ok := s2[k]; ok {
			r = append(r, k)
		}
	}

	fmt.Println(s1, s2)
	return r
}
