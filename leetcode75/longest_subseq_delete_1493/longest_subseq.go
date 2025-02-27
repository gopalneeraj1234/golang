package longestsubseqdelete1493

func longestSubarray(nums []int) int {
	startIndex := 0
	endIndex := 0
	ones := 0
	zeros := 0
	longestOneLen := 0
	for endIndex < len(nums) {
		if nums[endIndex] == 0 {
			zeros++
		} else {
			ones++
		}
		if zeros == 0 {
			longestOneLen = max(longestOneLen, ones-1)
		} else {
			longestOneLen = max(longestOneLen, ones)
		}

		if zeros > 1 {
			if nums[startIndex] == 0 {
				zeros--
			} else {
				ones--
			}
			startIndex++
		}
		endIndex++
	}
	return longestOneLen
}
