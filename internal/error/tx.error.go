package apperror

import "errors"

var ErrFailCommit = errors.New("failed to commit transaction")
var ErrFailRollback = errors.New("failed to rollback transaction")
