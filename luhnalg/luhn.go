package luhnalg

import (
	"unicode"
)

type (
	luhnAlgorithm struct{}

	Manager interface {
		Validate(data string) bool
	}
)

func New() Manager {
	return &luhnAlgorithm{}
}

func (la *luhnAlgorithm) Validate(data string) bool {
	return la.isValid(data)
}

func (la *luhnAlgorithm) isValid(data string) bool {
	source := []rune(data)
	total := 0
	double := false

	for i := len(source) - 1; i > -1; i-- {
		if unicode.IsDigit(source[i]) {
			digit := int(source[i] - '0')

			if double {
				digit *= 2
			}

			double = !double

			if digit >= 10 {
				digit = digit - 9
			}

			total += digit
			continue
		}

		if unicode.IsSpace(source[i]) {
			continue
		}

		return false
	}

	result := total % 10

	return result == 0
}
