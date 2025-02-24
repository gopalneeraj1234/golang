package longestsubseqdelete1493

func longestSubarray(nums []int) int {
	startIndex := 0
	endIndex := 0
	numOnes := 0
	numZeros := 0
	longestOneLen := 0
	for endIndex < len(nums) {
		if nums[endIndex] == 0 {
			numZeros++
		} else {
			numOnes++
		}
		if numZeros == 0 {
			longestOneLen = max(longestOneLen, numOnes-1)
		} else {
			longestOneLen = max(longestOneLen, numOnes)
		}

		if numZeros > 1 {
			if nums[startIndex] == 0 {
				numZeros--
			} else {
				numOnes--
			}
			startIndex++
		}
		endIndex++
	}
	return longestOneLen
}
