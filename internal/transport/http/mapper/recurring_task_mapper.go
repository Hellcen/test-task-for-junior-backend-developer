package mapper

import (
	"errors"
	"example.com/taskservice/internal/domain/recurringrule"
	"example.com/taskservice/internal/transport/http/handlers"
	"fmt"
	"strings"
	"time"
)

func ValidateCreateInputDTO(input handlers.CreateInputDTO) (handlers.CreateInputDTO, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return handlers.CreateInputDTO{}, errors.New("title is required")
	}

	if input.Status == "" {
		input.Status = recurringrule.StatusActive
	}

	if !input.Status.Valid() {
		return handlers.CreateInputDTO{}, fmt.Errorf("%w: invalid status", handlers.ErrInvalidInput)
	}

	if !input.RecurrenceType.Valid() {
		return handlers.CreateInputDTO{}, fmt.Errorf("%w: invalid recurrence_type", handlers.ErrInvalidInput)
	}

	switch input.RecurrenceType {
	case recurringrule.TypeDaily:
		if input.RecurrenceInterval == nil || *input.RecurrenceInterval <= 0 {
			return handlers.CreateInputDTO{}, fmt.Errorf("%w: interval_days must be > 0 for daily recurrence", handlers.ErrInvalidInput)
		}
	case recurringrule.TypeMonthly:
		if input.RecurrenceDayOfMonth == nil || *input.RecurrenceDayOfMonth < 1 || *input.RecurrenceDayOfMonth > 30 {
			return handlers.CreateInputDTO{}, fmt.Errorf("%w: day_of_month must be between 1 and 30", handlers.ErrInvalidInput)
		}
	case recurringrule.TypeEvenOdd:
		if input.RecurrenceParity == nil || !input.RecurrenceParity.Valid() {
			return handlers.CreateInputDTO{}, fmt.Errorf("%w: parity must be 'even' or 'odd'", handlers.ErrInvalidInput)
		}
	case recurringrule.TypeSpecific:
		if len(input.RecurrenceSpecificDates) == 0 {
			return handlers.CreateInputDTO{}, fmt.Errorf("%w: specific_dates cannot be empty", handlers.ErrInvalidInput)
		}
	}

	if input.RecurrenceStartDate.IsZero() {
		input.RecurrenceStartDate = time.Now().UTC()
	}

	return input, nil
}
