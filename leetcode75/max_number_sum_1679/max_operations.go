package maxnumbersum1679

type Set struct {
	elements map[int]struct{}
}

func (s Set) remove(num int) {
	delete(s.elements, num)
}

func (s Set) size() int {
	return len(s.elements)
}

// Removes an element that is not equal to the given element
func (s Set) removeNotEqualElement(num int) (bool, int) {
	for key := range s.elements {
		if key != num {
			delete(s.elements, key)
			return true, key
		}
	}
	return false, -1
}

func (s *Set) contains(element int) bool {
	_, found := s.elements[element]
	return found
}

func (s *Set) append(num int) {
	s.elements[num] = struct{}{}
}

func createSet() *Set {
	set := &Set{
		elements: make(map[int]struct{}),
	}
	return set
}

// Create a set with numbers
// For each number in nums remove them from set if there is a match that leads to 0
func maxOperations(nums []int, k int) int {
	numMap := make(map[int]Set)
	for i, num := range nums {
		currVals, found := numMap[num]
		if !found {
			currVals := createSet()
			currVals.append(i)
			numMap[num] = *currVals
		} else {
			currVals.append(i)
		}
	}

	numOps := 0
	removedIndexes := createSet()

	for i, num := range nums {
		matchedIndexSet, found := numMap[k-num]

		if !removedIndexes.contains(i) && found {
			removed, removedIndex := matchedIndexSet.removeNotEqualElement(i)
			if removed {
				removedIndexes.append(removedIndex)
				if matchedIndexSet.size() == 0 {
					delete(numMap, k-num)
				}
				currIndexSet := numMap[num]
				currIndexSet.remove(i)
				if currIndexSet.size() == 0 {
					delete(numMap, num)
				}
				removedIndexes.append(i)
				numOps++
			}
		}
	}
	return numOps
}
