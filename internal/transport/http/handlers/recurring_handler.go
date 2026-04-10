package handlers

import "example.com/taskservice/internal/usecase/recurring"

type RecurringHandler struct {
	usecase recurring.Usecase
}
