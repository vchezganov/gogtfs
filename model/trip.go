//Done
package model

import "errors"

type TripDirectionID byte

const (
	TripDirectionIDOutbound TripDirectionID = 0
	TripDirectionIDInbound  TripDirectionID = 1
)

type TripWheelchairAccessible byte

const (
	TripWheelchairAccessibleNoInformation TripWheelchairAccessible = 0
	TripWheelchairAccessibleYes           TripWheelchairAccessible = 1
	TripWheelchairAccessibleNo            TripWheelchairAccessible = 2
)

type TripBikesAllowed byte

const (
	TripBikesAllowedNoInformation TripBikesAllowed = 0
	TripBikesAllowedYes           TripBikesAllowed = 1
	TripBikesAllowedNo            TripBikesAllowed = 2
)

type Trip struct {
	RouteID              string                   `csv:"route_id"`
	ServiceID            string                   `csv:"service_id"`
	TripID               string                   `csv:"trip_id"`
	TripHeadsign         string                   `csv:"trip_headsign"`
	TripShortName        string                   `csv:"trip_short_name"`
	DirectionID          TripDirectionID          `csv:"direction_id,ParseDirectionID"`
	BlockID              string                   `csv:"block_id"`
	ShapeID              string                   `csv:"shape_id"`
	WheelchairAccessible TripWheelchairAccessible `csv:"wheelchair_accessible,ParseWheelchairAccessible"`
	BikesAllowed         TripBikesAllowed         `csv:"bikes_allowed,ParseBikesAllowed"`
}

func (t Trip) TableFileName() string {
	return "trips.txt"
}

func (t *Trip) ParseDirectionID(value string) error {
	switch value {
	case "0":
		t.DirectionID = TripDirectionIDOutbound

	case "1":
		t.DirectionID = TripDirectionIDInbound

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (t *Trip) ParseWheelchairAccessible(value string) error {
	switch value {
	case "", "0":
		t.WheelchairAccessible = TripWheelchairAccessibleNoInformation

	case "1":
		t.WheelchairAccessible = TripWheelchairAccessibleYes

	case "2":
		t.WheelchairAccessible = TripWheelchairAccessibleNo

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (t *Trip) ParseBikesAllowed(value string) error {
	switch value {
	case "", "0":
		t.BikesAllowed = TripBikesAllowedNoInformation

	case "1":
		t.BikesAllowed = TripBikesAllowedYes

	case "2":
		t.BikesAllowed = TripBikesAllowedNo

	default:
		return errors.New("invalid value")
	}

	return nil
}
