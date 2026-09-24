package apperror

import (
	"errors"
	"fmt"
)

var ErrNotEmail = errors.New("invalid email format")

func CreateNotEnoughLengthErr(field string, min int) error {
	return fmt.Errorf("%s have to be atleast %d character", field, min)
}
