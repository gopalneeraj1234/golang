package containermostwater11

func maxArea(height []int) int {
	leftPtr := 0
	rightPtr := len(height) - 1

	containerArea := 0
	for leftPtr < rightPtr {
		minHeight := min(height[leftPtr], height[rightPtr])
		width := rightPtr - leftPtr
		containerArea = max(width*minHeight, containerArea)
		if height[leftPtr] < height[rightPtr] {
			leftPtr++
		} else if height[leftPtr] > height[rightPtr] {
			rightPtr--
		} else {
			leftPtr++
			rightPtr--
		}
	}
	return containerArea
}
