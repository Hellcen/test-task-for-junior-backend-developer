package handlers

import (
	"errors"
	"example.com/taskservice/internal/usecase/recurring"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type RecurringHandler struct {
	usecase recurring.Usecase
}
