package internal

// isNumber checks if a given ch
// is a number
func isNumber(ch string) bool {

	if len(ch) == 0 {
		return false
	}

	for _, ch := range ch {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}
