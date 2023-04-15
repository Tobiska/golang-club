package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type unwrapError struct {
	error
}

func (e unwrapError) Unwrap() error {
	return e.error
}

type UserError struct {
	unwrapError
}

type SystemError struct {
	unwrapError
}

type InternalError struct {
	unwrapError
}

func WrapUserError(err error) error {
	if err == nil {
		return nil
	}
	return &UserError{
		unwrapError: unwrapError{fmt.Errorf("user error: %w", err)},
	}
}

func IsUserError(err error) bool {
	var errCast *UserError
	return errors.As(err, &errCast)
}

func WrapSystemError(err error) error {
	if err == nil {
		return err
	}
	return &SystemError{
		unwrapError: unwrapError{errors.WithStack(fmt.Errorf("system error: %w", err))},
	}
}

func IsSystemError(err error) bool {
	var errCast *SystemError
	return errors.As(err, &errCast)
}

func WrapInternalError(err error) error {
	if err == nil {
		return err
	}
	return &InternalError{
		unwrapError: unwrapError{errors.WithStack(fmt.Errorf("internal error: %w", err))},
	}
}

func IsInternalError(err error) bool {
	var errCast *InternalError
	return errors.As(err, &errCast)
}

func IsInheritInternalError(err error) bool {
	return IsInternalError(err)
}

func IsInheritSystemError(err error) bool {
	return IsSystemError(err)
}

func IsInheritUserError(err error) bool {
	return IsUserError(err) || IsInheritSystemError(err) || IsInheritInternalError(err)
}

func main() {
	userError := WrapUserError(errors.New("error"))

	fmt.Println(IsUserError(fmt.Errorf("%w", userError)))
}
