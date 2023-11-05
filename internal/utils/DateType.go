package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalDateTime time.Time

type LocalDate time.Time

func (t *LocalDateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *LocalDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}

func ToLocalDateTime(t time.Time) *LocalDateTime {
	localDateTime := LocalDateTime(t)
	return &localDateTime
}

func ToLocalDate(t time.Time) *LocalDate {
	localDate := LocalDate(t)
	return &localDate
}

func (t LocalDate) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *LocalDate) Scan(value interface{}) error {
	if value == nil {
		*t = LocalDate(time.Time{})
		return nil
	}
	if v, ok := value.(time.Time); ok {
		*t = LocalDate(v)
		return nil
	}
	return fmt.Errorf("unsupported type for LocalDate")
}

func (t LocalDateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *LocalDateTime) Scan(value interface{}) error {
	if value == nil {
		*t = LocalDateTime(time.Time{})
		return nil
	}
	if v, ok := value.(time.Time); ok {
		*t = LocalDateTime(v)
		return nil
	}
	return fmt.Errorf("unsupported type for LocalDateTime")
}
