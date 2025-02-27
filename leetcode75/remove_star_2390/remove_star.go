package removestar2390

func removeStars(s string) string {
	var charStack []rune
	for _, ch := range s {
		if ch == '*' {
			charStack = charStack[0 : len(charStack)-1] //pop
		} else {
			charStack = append(charStack, ch)
		}
	}
	return string(charStack)
}
