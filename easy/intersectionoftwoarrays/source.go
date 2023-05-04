package insertationoftwoarrays

import "fmt"

func intersection(nums1, nums2 []int) []int {
	s1 := make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		if _, ok := s1[nums1[i]]; !ok {
			s1[nums1[i]] = struct{}{}
		}
	}

	r := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := s1[nums2[i]]; ok {
			r = append(r, nums2[i])
			delete(s1, nums2[i])
		}
	}

	fmt.Println(s1, r)
	return r
}
