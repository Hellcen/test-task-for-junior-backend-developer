package recurringrule

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"example.com/taskservice/internal/domain/task"
	"time"
)

var ErrNotFound = errors.New("recurringrule rule not found")

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

type Parity string

const (
	ParityEven Parity = "even"
	ParityOdd  Parity = "odd"
)

func (p Parity) Valid() bool {
	switch p {
	case ParityEven, ParityOdd:
		return true
	}
	return false
}

type JSONDates []time.Time

func (j JSONDates) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

type RecurringRule struct {
	Id          uint64      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      task.Status `json:"status"`

	// Fields for periodicity
	RecurrenceType          RecurrenceType `json:"recurrence_type,omitempty"`
	RecurrenceInterval      *int           `json:"recurrence_interval,omitempty"`       // For daily: every N days
	RecurrenceDayOfMonth    *int           `json:"recurrence_day_of_month,omitempty"`   // For monthly: day of the month
	RecurrenceParity        *Parity        `json:"recurrence_parity,omitempty"`         // For evenodd
	RecurrenceSpecificDates []time.Time    `json:"recurrence_specific_dates,omitempty"` // For specific

	RecurrenceStartDate *time.Time `json:"recurrence_start_date,omitempty"` // Start date for tasks
	RecurrenceEndDate   *time.Time `json:"recurrence_end_date,omitempty"`   // End date for tasks
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
