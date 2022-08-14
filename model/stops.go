//Done
package model

import (
	"errors"
	"strconv"
)

type StopLocationType byte

const (
	StopLocationTypeStopOrPlatform StopLocationType = 0
	StopLocationTypeStation        StopLocationType = 1
	StopLocationTypeEntranceExit   StopLocationType = 2
	StopLocationTypeGenericNone    StopLocationType = 3
	StopLocationTypeBoardingArea   StopLocationType = 4
)

type StopWheelchairBoarding byte

const (
	StopWheelchairBoardingInherit    StopWheelchairBoarding = 0
	StopWheelchairBoardingAllowed    StopWheelchairBoarding = 1
	StopWheelchairBoardingNotAllowed StopWheelchairBoarding = 2
)

type Stop struct {
	StopID             string                 `csv:"stop_id"`
	StopCode           string                 `csv:"stop_code"`
	StopName           string                 `csv:"stop_name"`
	StopDesc           string                 `csv:"stop_desc"`
	StopLat            float64                `csv:"stop_lat"`
	StopLon            float64                `csv:"stop_lon"`
	ZoneID             string                 `csv:"zone_id"`
	StopURL            string                 `csv:"stop_url"`
	LocationType       StopLocationType       `csv:"location_type,ParseLocationType"`
	ParentStation      string                 `csv:"parent_station"`
	StopTimezone       string                 `csv:"stop_timezone"`
	WheelchairBoarding StopWheelchairBoarding `csv:"wheelchair_boarding,ParseWheelchairBoarding"`
	LevelID            string                 `csv:"level_id"`
	PlatformCode       string                 `csv:"platform_code"`
}

func (s Stop) TableFileName() string {
	return "stops.txt"
}

func (s *Stop) ParseLocationType(value string) error {
	if value == "" {
		s.LocationType = StopLocationTypeStopOrPlatform
		return nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if !(0 <= number && number <= 4) {
		return errors.New("invalid value")
	}

	s.LocationType = StopLocationType(number)
	return nil
}

func (s *Stop) ParseWheelchairBoarding(value string) error {
	if value == "" {
		s.WheelchairBoarding = StopWheelchairBoardingInherit
		return nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if !(0 <= number && number <= 2) {
		return errors.New("invalid value")
	}

	s.WheelchairBoarding = StopWheelchairBoarding(number)
	return nil
}
