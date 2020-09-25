package valid_palyndrom_with_removal

// ValidPalyndromWithRemoval creates a palyndrom by removing one char
func ValidPalyndromWithRemoval(str string) (palyndrom bool) {

	if len(str) == 0 {
		return
	}

	head := 0
	tail := len(str) - 1

	for head <= tail {
		if str[head] != str[tail] {
			return
		}

		head++
		tail--
	}

	return
}
