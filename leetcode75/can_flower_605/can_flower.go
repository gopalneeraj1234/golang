package canflower605

const (
	EMPTY     = 0
	NON_EMPTY = 1
	PLACE_POT = 1
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	prev := EMPTY
	next := EMPTY
	for i, element := range flowerbed {
		if n == 0 {
			break
		}

		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}
		if element == EMPTY && prev == EMPTY && next == EMPTY {
			flowerbed[i] = PLACE_POT
			n--
		}
		prev = flowerbed[i]
	}
	return n == 0
}
