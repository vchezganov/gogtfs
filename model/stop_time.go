//Done
package model

import "errors"

type StopTimePickupType byte

const (
	StopTimePickupTypeScheduled StopTimePickupType = 0
	StopTimePickupTypeNoPickup  StopTimePickupType = 1
	StopTimePickupTypeMustPhone StopTimePickupType = 2
	StopTimePickupTypeDriver    StopTimePickupType = 3
)

type StopTimeDropOffType byte

const (
	StopTimeDropOffTypeScheduled StopTimeDropOffType = 0
	StopTimeDropOffTypeNoDropOff StopTimeDropOffType = 1
	StopTimeDropOffTypeMustPhone StopTimeDropOffType = 2
	StopTimeDropOffTypeDriver    StopTimeDropOffType = 3
)

type StopTimeContinuousPickup byte

const (
	StopTimeContinuousPickupContinuous   StopTimeContinuousPickup = 0
	StopTimeContinuousPickupNoContinuous StopTimeContinuousPickup = 1
	StopTimeContinuousPickupPhone        StopTimeContinuousPickup = 2
	StopTimeContinuousPickupDriver       StopTimeContinuousPickup = 3
)

type StopTimeContinuousDropOff byte

const (
	StopTimeContinuousDropOffContinuous   StopTimeContinuousDropOff = 0
	StopTimeContinuousDropOffNoContinuous StopTimeContinuousDropOff = 1
	StopTimeContinuousDropOffPhone        StopTimeContinuousDropOff = 2
	StopTimeContinuousDropOffDriver       StopTimeContinuousDropOff = 3
)

type StopTimeTimepoint byte

const (
	StopTimeTimepointApproximate StopTimeTimepoint = 0
	StopTimeTimepointExact       StopTimeTimepoint = 1
)

type StopTime struct {
	TripID            string                    `csv:"trip_id"`
	ArrivalTime       GTFSTime                  `csv:"arrival_time,ParseArrivalTime"`
	DepartureTime     GTFSTime                  `csv:"departure_time,ParseDepartureTime"`
	StopID            string                    `csv:"stop_id"`
	StopSequence      uint                      `csv:"stop_sequence"`
	StopHeadsign      string                    `csv:"stop_headsign"`
	PickupType        StopTimePickupType        `csv:"pickup_type,ParsePickupType"`
	DropOffType       StopTimeDropOffType       `csv:"drop_off_type,ParseDropOffType"`
	ContinuousPickup  StopTimeContinuousPickup  `csv:"continuous_pickup,ParseContinuousPickup"`
	ContinuousDropOff StopTimeContinuousDropOff `csv:"continuous_drop_off,ParseContinuousDropOff"`
	ShapeDistTraveled float64                   `csv:"shape_dist_traveled"`
	Timepoint         StopTimeTimepoint         `csv:"timepoint,ParseTimepoint"`
}

func (s StopTime) TableFileName() string {
	return "stop_times.txt"
}

func (s *StopTime) ParseArrivalTime(value string) (err error) {
	if value == "" {
		s.ArrivalTime = NewEmptyGTFSTime()
		return nil
	}

	s.ArrivalTime, err = NewGTFSTime(value)
	return
}

func (s *StopTime) ParseDepartureTime(value string) (err error) {
	if value == "" {
		s.ArrivalTime = NewEmptyGTFSTime()
		return nil
	}

	s.ArrivalTime, err = NewGTFSTime(value)
	return
}

func (s *StopTime) ParsePickupType(value string) (err error) {
	switch value {
	case "", "0":
		s.PickupType = StopTimePickupTypeScheduled

	case "1":
		s.PickupType = StopTimePickupTypeNoPickup

	case "2":
		s.PickupType = StopTimePickupTypeMustPhone

	case "3":
		s.PickupType = StopTimePickupTypeDriver

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (s *StopTime) ParseDropOffType(value string) (err error) {
	switch value {
	case "", "0":
		s.DropOffType = StopTimeDropOffTypeScheduled

	case "1":
		s.DropOffType = StopTimeDropOffTypeNoDropOff

	case "2":
		s.DropOffType = StopTimeDropOffTypeMustPhone

	case "3":
		s.DropOffType = StopTimeDropOffTypeDriver

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (s *StopTime) ParseContinuousPickup(value string) (err error) {
	switch value {
	case "0":
		s.ContinuousPickup = StopTimeContinuousPickupContinuous

	case "", "1":
		s.ContinuousPickup = StopTimeContinuousPickupNoContinuous

	case "2":
		s.ContinuousPickup = StopTimeContinuousPickupPhone

	case "3":
		s.ContinuousPickup = StopTimeContinuousPickupDriver

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (s *StopTime) ParseContinuousDropOff(value string) (err error) {
	switch value {
	case "0":
		s.ContinuousDropOff = StopTimeContinuousDropOffContinuous

	case "", "1":
		s.ContinuousDropOff = StopTimeContinuousDropOffNoContinuous

	case "2":
		s.ContinuousDropOff = StopTimeContinuousDropOffPhone

	case "3":
		s.ContinuousDropOff = StopTimeContinuousDropOffDriver

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (s *StopTime) ParseTimepoint(value string) (err error) {
	switch value {
	case "0":
		s.Timepoint = StopTimeTimepointApproximate

	case "", "1":
		s.Timepoint = StopTimeTimepointExact

	default:
		return errors.New("invalid value")
	}

	return nil
}
