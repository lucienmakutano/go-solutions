package valid_palyndrom_with_removal

// ValidPalyndromWithRemoval creates a palyndrom by removing one char
func ValidPalyndromWithRemoval(str string) bool {

	palyndrom := false
	length := len(str)

	if length == 0 {
		return palyndrom
	} else if str[0] != str[length-1] {
		return palyndrom
	}

	head := 0
	tail := len(str) - 1

	for head <= tail {
		if str[head] != str[tail] {
			str = string(str[:tail]) + string(str[tail+1:])
			return ValidPalyndromWithRemoval(str)
		}

		head++
		tail--
	}

	palyndrom = true

	return palyndrom
}
