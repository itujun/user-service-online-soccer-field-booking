package errors

import "errors"

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(append(GeneralErrors[:], UserErrors[:]...))

	for _, item := range allErrors {
		if errors.Is(err, item) {
			return true
		}
	}
	return false
}
