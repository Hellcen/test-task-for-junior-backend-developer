package recurring

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"example.com/taskservice/internal/domain/recurringrule"
)

func ValidateCreateInput(input CreateInputDTO) (CreateInputDTO, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateInputDTO{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = recurringrule.StatusActive
	}

	if !input.Status.Valid() {
		return CreateInputDTO{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	if !input.RecurrenceType.Valid() {
		return CreateInputDTO{}, fmt.Errorf("%w: invalid recurrence_type", ErrInvalidInput)
	}

	// Валидация конфигурации в зависимости от типа
	if err := validateConfig(input.RecurrenceType, input.Config); err != nil {
		return CreateInputDTO{}, err
	}

	if input.StartDate.IsZero() {
		input.StartDate = time.Now().UTC()
	}

	return input, nil
}

func ValidateUpdateInput(input UpdateInputDTO, existing recurringrule.RecurringRule) (UpdateInputDTO, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		input.Title = existing.Title
	}

	if input.Description == "" {
		input.Description = existing.Description
	}

	if input.Status == "" {
		input.Status = existing.Status
	}

	if !input.Status.Valid() {
		return UpdateInputDTO{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	if input.RecurrenceType == "" {
		input.RecurrenceType = existing.RecurrenceType
	}

	if !input.RecurrenceType.Valid() {
		return UpdateInputDTO{}, fmt.Errorf("%w: invalid recurrence_type", ErrInvalidInput)
	}

	// Если конфиг не передан, используем существующий
	if input.Config == nil {
		input.Config = existing.RecurrenceConfig
	} else {
		if err := validateConfig(input.RecurrenceType, input.Config); err != nil {
			return UpdateInputDTO{}, err
		}
	}

	if input.StartDate.IsZero() {
		input.StartDate = existing.StartDate
	}

	if input.EndDate == nil {
		input.EndDate = existing.EndDate
	}

	return input, nil
}

func validateConfig(recurrenceType recurringrule.RecurrenceType, config json.RawMessage) error {
	switch recurrenceType {
	case recurringrule.TypeDaily:
		var cfg recurringrule.DailyConfig
		if err := json.Unmarshal(config, &cfg); err != nil {
			return fmt.Errorf("%w: invalid daily config: %v", ErrInvalidInput, err)
		}
		if cfg.IntervalDays <= 0 {
			return fmt.Errorf("%w: interval_days must be > 0 for daily recurrence", ErrInvalidInput)
		}

	case recurringrule.TypeMonthly:
		var cfg recurringrule.MonthlyConfig
		if err := json.Unmarshal(config, &cfg); err != nil {
			return fmt.Errorf("%w: invalid monthly config: %v", ErrInvalidInput, err)
		}
		if cfg.DayOfMonth < 1 || cfg.DayOfMonth > 30 {
			return fmt.Errorf("%w: day_of_month must be between 1 and 30", ErrInvalidInput)
		}

	case recurringrule.TypeEvenOdd:
		var cfg recurringrule.EvenOddConfig
		if err := json.Unmarshal(config, &cfg); err != nil {
			return fmt.Errorf("%w: invalid evenodd config: %v", ErrInvalidInput, err)
		}
		if cfg.Parity != "even" && cfg.Parity != "odd" {
			return fmt.Errorf("%w: parity must be 'even' or 'odd'", ErrInvalidInput)
		}

	case recurringrule.TypeSpecific:
		var cfg recurringrule.SpecificConfig
		if err := json.Unmarshal(config, &cfg); err != nil {
			return fmt.Errorf("%w: invalid specific config: %v", ErrInvalidInput, err)
		}
		if len(cfg.Dates) == 0 {
			return fmt.Errorf("%w: dates cannot be empty for specific recurrence", ErrInvalidInput)
		}

	default:
		return fmt.Errorf("%w: unknown recurrence type", ErrInvalidInput)
	}

	return nil
}
