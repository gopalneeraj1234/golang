package findaltitude1732

func largestAltitude(gain []int) int {
	maxAltitude := 0

	for i, altGain := range gain {
		if i == 0 {
			gain[0] = 0 + altGain
		} else {
			gain[i] = gain[i-1] + altGain
		}
		maxAltitude = max(maxAltitude, gain[i])
	}
	return maxAltitude
}
