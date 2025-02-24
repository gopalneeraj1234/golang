package findpivot724

func pivotIndex(nums []int) int {
	prefixSumArr := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			prefixSumArr[0] = num
		} else {
			prefixSumArr[i] = prefixSumArr[i-1] + num
		}
	}

	//For the Zeroth index
	leftSum := 0
	rightSum := prefixSumArr[len(nums)-1] - prefixSumArr[0]
	currIndex := 0

	for leftSum != rightSum && currIndex+1 < len(nums) {
		currIndex++
		leftSum = prefixSumArr[currIndex-1]
		rightSum = prefixSumArr[len(nums)-1] - prefixSumArr[currIndex]
	}
	if leftSum == rightSum {
		return currIndex
	} else {
		return -1
	}
}
