//Done
package model

import (
	"errors"
	"time"
)

type CalendarDateException byte

const (
	CalendarDateExceptionAdded   CalendarDateException = 1
	CalendarDateExceptionRemoved CalendarDateException = 2
)

type CalendarDate struct {
	ServiceID     string                `csv:"service_id"`
	Date          time.Time             `csv:"date,ParseDate"`
	ExceptionType CalendarDateException `csv:"exception_type,ParseExceptionType"`
}

func (c CalendarDate) TableFileName() string {
	return "calendar_dates.txt"
}

func (c *CalendarDate) ParseDate(value string) (err error) {
	c.Date, err = parseDate(value)
	return
}

func (c *CalendarDate) ParseExceptionType(value string) error {
	switch value {
	case "1":
		c.ExceptionType = CalendarDateExceptionAdded

	case "2":
		c.ExceptionType = CalendarDateExceptionRemoved

	default:
		return errors.New("invalid value")
	}

	return nil
}
