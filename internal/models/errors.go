package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrDuplicateEmail = errors.New("models: Duplicate EMail records found")
var ErrInvalidCredentials = errors.New("models: Invalid credentials records found")
