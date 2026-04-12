package recurringrule

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

var ErrNotFound = errors.New("recurring rule not found")

type Status string

const (
	StatusActive   Status = "active"
	StatusPaused   Status = "paused"
	StatusArchived Status = "archived"
)

func (s Status) Valid() bool {
	switch s {
	case StatusActive, StatusPaused, StatusArchived:
		return true
	}
	return false
}

type RecurrenceType string

const (
	TypeDaily    RecurrenceType = "daily"
	TypeMonthly  RecurrenceType = "monthly"
	TypeEvenOdd  RecurrenceType = "evenodd"
	TypeSpecific RecurrenceType = "specific"
)

func (t RecurrenceType) Valid() bool {
	switch t {
	case TypeDaily, TypeMonthly, TypeEvenOdd, TypeSpecific:
		return true
	}
	return false
}

type DailyConfig struct {
	IntervalDays int `json:"interval_days"`
}

type MonthlyConfig struct {
	DayOfMonth int `json:"day_of_month"`
}

type EvenOddConfig struct {
	Parity string `json:"parity"` // "even" или "odd"
}

type SpecificConfig struct {
	Dates []time.Time `json:"dates"`
}

// JSONRawMessage для работы с JSONB
type JSONRawMessage json.RawMessage

func (j *JSONRawMessage) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	*j = append((*j)[0:0], bytes...)
	return nil
}

func (j JSONRawMessage) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return []byte(j), nil
}

func (j JSONRawMessage) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSONRawMessage) UnmarshalJSON(data []byte) error {
	if j == nil {
		return nil
	}
	*j = append((*j)[0:0], data...)
	return nil
}

type RecurringRule struct {
	ID               int64           `json:"id"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	Status           Status          `json:"status"`
	RecurrenceType   RecurrenceType  `json:"recurrence_type"`
	RecurrenceConfig json.RawMessage `json:"recurrence_config"`
	StartDate        time.Time       `json:"start_date"`
	EndDate          *time.Time      `json:"end_date,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

// Helper методы для получения конфигурации
func (r *RecurringRule) GetDailyConfig() (*DailyConfig, error) {
	if r.RecurrenceType != TypeDaily {
		return nil, errors.New("not a daily rule")
	}
	var config DailyConfig
	if err := json.Unmarshal(r.RecurrenceConfig, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *RecurringRule) GetMonthlyConfig() (*MonthlyConfig, error) {
	if r.RecurrenceType != TypeMonthly {
		return nil, errors.New("not a monthly rule")
	}
	var config MonthlyConfig
	if err := json.Unmarshal(r.RecurrenceConfig, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *RecurringRule) GetEvenOddConfig() (*EvenOddConfig, error) {
	if r.RecurrenceType != TypeEvenOdd {
		return nil, errors.New("not an evenodd rule")
	}
	var config EvenOddConfig
	if err := json.Unmarshal(r.RecurrenceConfig, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *RecurringRule) GetSpecificConfig() (*SpecificConfig, error) {
	if r.RecurrenceType != TypeSpecific {
		return nil, errors.New("not a specific rule")
	}
	var config SpecificConfig
	if err := json.Unmarshal(r.RecurrenceConfig, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
