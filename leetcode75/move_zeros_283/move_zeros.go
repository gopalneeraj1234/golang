package movezeros283

func moveZeroes(nums []int) {
	zeroPtr := 0
	nonZeroPtr := 0

	for zeroPtr < len(nums) && nonZeroPtr < len(nums) {
		for zeroPtr < len(nums) {
			if nums[zeroPtr] == 0 {
				break
			}
			zeroPtr++
		}

		for nonZeroPtr < len(nums) {
			if nums[nonZeroPtr] != 0 {
				break
			}
			nonZeroPtr++
		}
		if nonZeroPtr >= len(nums) || zeroPtr >= len(nums) {
			break
		}
		if zeroPtr < nonZeroPtr {
			swap(nums, zeroPtr, nonZeroPtr)
		} else {
			nonZeroPtr++
		}
	}
}

func swap(nums []int, zeroPtr, nonZeroPtr int) {
	temp := nums[zeroPtr]
	nums[zeroPtr] = nums[nonZeroPtr]
	nums[nonZeroPtr] = temp
}
