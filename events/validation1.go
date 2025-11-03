package events

import (
	"regexp"
)

func IsValidTitle(title string) bool {
	pattern := "^[a-zA-Z0-9 ]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
