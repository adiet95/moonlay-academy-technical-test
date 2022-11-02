package libs

import (
	"errors"
	"regexp"
)

func Validation(title, desc string) error {
	if len(title) > 100 && len(title) <= 0 {
		return errors.New("title must be greater than nil character & less than 100")
	}
	alpha := regexp.MustCompile(`^[A-Za-z0-9\s]+[A-Za-z0-9\s]+$(\.0-9+)?`)
	ok := alpha.MatchString(title)
	if !ok {
		return errors.New("title must be alphanumeric")
	}

	if len(desc) > 1000 && len(desc) < 0 {
		return errors.New("title must be greater than nil character & less than 1000")
	}
	return nil
}
