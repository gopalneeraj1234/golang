package maxaverage643

func findMaxAverage(nums []int, k int) float64 {
	windowStart := 0
	windowEnd := k - 1

	windowSum := getSum(windowStart, windowEnd, nums)
	maxWindowSum := windowSum
	windowStart++
	windowEnd++
	for windowEnd < len(nums) {
		windowSum -= nums[windowStart-1]
		windowSum += nums[windowEnd]
		maxWindowSum = max(maxWindowSum, windowSum)
		windowStart++
		windowEnd++
	}
	return float64(maxWindowSum) / float64(k)
}

func getSum(windowStart, windowEnd int, nums []int) int {
	sum := 0
	for i := windowStart; i <= windowEnd; i++ {
		sum += nums[i]
	}
	return sum
}
