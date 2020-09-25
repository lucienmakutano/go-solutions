package valid_palyndrom_with_removal

// ValidPalyndromWithRemoval creates a palyndrom by removing one char
func ValidPalyndromWithRemoval(str string) bool {
	palyndrom := false

	if len(str) == 0 {
		return palyndrom
	}

	head := 0
	tail := len(str) - 1

	for head <= tail {
		if str[head] != str[tail] {
			return palyndrom
		}

		head++
		tail--
	}

	palyndrom = true

	return palyndrom
}
