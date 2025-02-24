package finddiffarr2215

import "iter"

type Set struct {
	elements map[int]struct{}
}

func (s *Set) Items() iter.Seq[int] {
	return func(yield func(int) bool) {
		for num := range s.elements {
			if !yield(num) {
				return
			}
		}
	}
}

func (s *Set) contains(num int) bool {
	_, found := s.elements[num]
	return found
}

func (s *Set) append(num int) bool {
	keyExists := false
	if s.contains(num) {
		keyExists = false
	}
	s.elements[num] = struct{}{}
	return keyExists
}

func createSetFromSlice(nums []int) *Set {
	set := &Set{
		elements: make(map[int]struct{}),
	}
	for _, num := range nums {
		set.append(num)
	}
	return set
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Set := createSetFromSlice(nums1)
	nums2Set := createSetFromSlice(nums2)

	retSlice := make([][]int, 2)
	for num1 := range nums1Set.Items() {
		if !nums2Set.contains(num1) {
			retSlice[0] = append(retSlice[0], num1)
		}

	}

	for num2 := range nums2Set.Items() {
		if !nums1Set.contains(num2) {
			retSlice[1] = append(retSlice[1], num2)
		}
	}
	return retSlice
}
