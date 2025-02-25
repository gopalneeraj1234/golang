package asteroid735

func resolveCollision(asteroidStack *[]int) {
	if len(*asteroidStack) <= 1 {
		return
	}
	lastAsteriod := (*asteroidStack)[len(*asteroidStack)-1]
	if lastAsteriod > 0 {
		return
	}
	prevAsteroid := (*asteroidStack)[len(*asteroidStack)-2]
	for prevAsteroid > 0 && lastAsteriod < 0 {
		if prevAsteroid == (-1)*lastAsteriod {
			*asteroidStack = (*asteroidStack)[:len(*asteroidStack)-2]
		} else if prevAsteroid < (-1)*lastAsteriod {
			(*asteroidStack)[len(*asteroidStack)-2] = lastAsteriod
			*asteroidStack = (*asteroidStack)[:len(*asteroidStack)-1]
		} else {
			*asteroidStack = (*asteroidStack)[:len(*asteroidStack)-1]
		}
		if len(*asteroidStack) <= 1 {
			return
		}
		prevAsteroid = (*asteroidStack)[len(*asteroidStack)-2]
		lastAsteriod = (*asteroidStack)[len(*asteroidStack)-1]
	}
}

func asteroidCollision(asteroids []int) []int {
	asteroidStack := []int{}

	for _, asteroid := range asteroids {
		asteroidStack = append(asteroidStack, asteroid)
		resolveCollision(&asteroidStack)
	}
	return asteroidStack
}
