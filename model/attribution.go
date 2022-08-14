//Done
package model

import "errors"

type AttributionIsProducer bool

const (
	AttributionIsProducerNoRole  AttributionIsProducer = false
	AttributionIsProducerHasRole AttributionIsProducer = true
)

type AttributionIsOperator bool

const (
	AttributionIsOperatorNoRole  AttributionIsOperator = false
	AttributionIsOperatorHasRole AttributionIsOperator = true
)

type AttributionIsAuthority bool

const (
	AttributionIsAuthorityNoRole  AttributionIsAuthority = false
	AttributionIsAuthorityHasRole AttributionIsAuthority = true
)

type Attribution struct {
	AttributionID    string                 `csv:"attribution_id"`
	AgencyID         string                 `csv:"agency_id"`
	RouteID          string                 `csv:"route_id"`
	TripID           string                 `csv:"trip_id"`
	OrganizationName string                 `csv:"organization_name"`
	IsProducer       AttributionIsProducer  `csv:"is_producer,ParseIsProducer"`
	IsOperator       AttributionIsOperator  `csv:"is_operator,ParseIsOperator"`
	IsAuthority      AttributionIsAuthority `csv:"is_authority,ParseIsAuthority"`
	AttributionURL   string                 `csv:"attribution_url"`
	AttributionEmail string                 `csv:"attribution_email"`
	AttributionPhone string                 `csv:"attribution_phone"`
}

func (a Attribution) TableFileName() string {
	return "attributions.txt"
}

func (a *Attribution) ParseIsProducer(value string) error {
	switch value {
	case "", "0":
		a.IsProducer = AttributionIsProducerNoRole

	case "1":
		a.IsProducer = AttributionIsProducerHasRole

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (a *Attribution) ParseIsOperator(value string) error {
	switch value {
	case "", "0":
		a.IsOperator = AttributionIsOperatorNoRole

	case "1":
		a.IsOperator = AttributionIsOperatorHasRole

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (a *Attribution) ParseIsAuthority(value string) error {
	switch value {
	case "", "0":
		a.IsAuthority = AttributionIsAuthorityNoRole

	case "1":
		a.IsAuthority = AttributionIsAuthorityHasRole

	default:
		return errors.New("invalid value")
	}

	return nil
}
