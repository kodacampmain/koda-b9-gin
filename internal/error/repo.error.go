package apperror

import "errors"

var ErrNoData = errors.New("no data")
var ErrAlreadyExists = errors.New("user already exists")
var ErrNoRowAffected = errors.New("no row affected")
var ErrInvalidUsernamePassword = errors.New("invalid username or password")
var ErrEmptyUsernamePassword = errors.New("username or password cannot be empty")
