package main

import (
	"errors"
	"fmt"
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

var (
	ErrValidation = errors.New("Validation failed")
	ErrSignal     = errors.New("Recive signal")
)

func (s *stepErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v ", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}
	return t.step == s.step
}

func (s *stepErr) Unwrap() error {
	return s.cause
}
