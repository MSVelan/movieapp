package repository

import "errors"

// ErrNotFound returned when requested record is not found.
var ErrNotFound = errors.New("not found")
