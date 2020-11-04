package core

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists       = errors.New("already exists")
	ErrCouldNotBeEmpty     = errors.New("could not be empty")
	ErrInternalApplication = errors.New("internal application error")
	ErrIsRequired          = errors.New("is required")
	ErrNotFound            = errors.New("not found")
	ErrIsNotValid          = errors.New("is not valid")
	ErrIdCouldNotBeEmpty   = fmt.Errorf("%s %w", "id", ErrCouldNotBeEmpty)
	ErrShouldBeUnique      = errors.New("should be unique")

	FormatToErrFprintf = "fmt.Fprintf fatal: %v\n"
	FormatToErrKitLog  = "kit/log fatal: %v\n"
)
