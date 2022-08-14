//Done
package model

import (
	"errors"
	"time"
)

type CalendarService bool

const (
	CalendarServiceNotAvailable CalendarService = false
	CalendarServiceAvailable    CalendarService = true
)

type Calendar struct {
	ServiceID string          `csv:"service_id"`
	Monday    CalendarService `csv:"monday,ParseMonday"`
	Tuesday   CalendarService `csv:"tuesday,ParseTuesday"`
	Wednesday CalendarService `csv:"wednesday,ParseWednesday"`
	Thursday  CalendarService `csv:"thursday,ParseThursday"`
	Friday    CalendarService `csv:"friday,ParseFriday"`
	Saturday  CalendarService `csv:"saturday,ParseSaturday"`
	Sunday    CalendarService `csv:"sunday,ParseSunday"`
	StartDate time.Time       `csv:"start_date,ParseStartDate"`
	EndDate   time.Time       `csv:"end_date,ParseEndDate"`
}

func (c Calendar) TableFileName() string {
	return "calendar.txt"
}

func (c *Calendar) parseService(value string) (CalendarService, error) {
	switch value {
	case "0":
		return CalendarServiceNotAvailable, nil

	case "1":
		return CalendarServiceAvailable, nil

	default:
		return CalendarServiceNotAvailable, errors.New("invalid value")
	}
}

func (c *Calendar) ParseMonday(value string) (err error) {
	c.Monday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseTuesday(value string) (err error) {
	c.Tuesday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseWednesday(value string) (err error) {
	c.Wednesday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseThursday(value string) (err error) {
	c.Thursday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseFriday(value string) (err error) {
	c.Friday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseSaturday(value string) (err error) {
	c.Saturday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseSunday(value string) (err error) {
	c.Sunday, err = c.parseService(value)
	return
}

func (c *Calendar) ParseStartDate(value string) (err error) {
	c.StartDate, err = parseDate(value)
	return
}

func (c *Calendar) ParseEndDate(value string) (err error) {
	c.EndDate, err = parseDate(value)
	return
}
