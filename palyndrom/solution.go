package palyndrom

// ValidPalyndromWithRemoval creates a palyndrom by removing one char
func ValidPalyndromWithRemoval(str string) bool {

	length := len(str)

	if length == 0 {
		return false
	} else if str[0] != str[length-1] {
		return false
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

	return true
}
