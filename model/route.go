//Done
package model

import (
	"errors"
	"strconv"
)

type RouteRouteType int

type RouteContinuousPickup byte

const (
	RouteContinuousPickupContinuous   RouteContinuousPickup = 0
	RouteContinuousPickupNoContinuous RouteContinuousPickup = 1
	RouteContinuousPickupPhone        RouteContinuousPickup = 2
	RouteContinuousPickupDriver       RouteContinuousPickup = 3
)

type RouteContinuousDropOff byte

const (
	RouteContinuousDropOffContinuous   RouteContinuousDropOff = 0
	RouteContinuousDropOffNoContinuous RouteContinuousDropOff = 1
	RouteContinuousDropOffPhone        RouteContinuousDropOff = 2
	RouteContinuousDropOffDriver       RouteContinuousDropOff = 3
)

type Route struct {
	RouteID           string                 `csv:"route_id"`
	AgencyID          string                 `csv:"agency_id"`
	RouteShortName    string                 `csv:"route_short_name"`
	RouteLongName     string                 `csv:"route_long_name"`
	RouteDesc         string                 `csv:"route_desc"`
	RouteType         RouteRouteType         `csv:"route_type"`
	RouteURL          string                 `csv:"route_url"`
	RouteColor        string                 `csv:"route_color"`
	RouteTextColor    string                 `csv:"route_text_color"`
	RouteSortOrder    uint                   `csv:"route_sort_order"`
	ContinuousPickup  RouteContinuousPickup  `csv:"continuous_pickup,ParseContinuousPickup"`
	ContinuousDropOff RouteContinuousDropOff `csv:"continuous_drop_off,ParseContinuousDropOff"`
}

func (r Route) TableFileName() string {
	return "routes.txt"
}

func (r *Route) ParseContinuousPickup(value string) error {
	if value == "" {
		r.ContinuousPickup = RouteContinuousPickupNoContinuous
		return nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if !(0 <= number && number <= 3) {
		return errors.New("invalid value")
	}

	r.ContinuousPickup = RouteContinuousPickup(number)
	return nil
}

func (r *Route) ParseContinuousDropOff(value string) error {
	if value == "" {
		r.ContinuousDropOff = RouteContinuousDropOffNoContinuous
		return nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if !(0 <= number && number <= 3) {
		return errors.New("invalid value")
	}

	r.ContinuousDropOff = RouteContinuousDropOff(number)
	return nil
}
