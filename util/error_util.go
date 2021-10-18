package util

import "strings"

func IsTooManyRequestError(err error) bool {
	if strings.Contains(strings.ToLower(err.Error()), "too many") {
		return true
	}
	return false
}
