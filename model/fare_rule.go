//Done
package model

type FareRule struct {
	FareID        string `csv:"fare_id"`
	RouteID       string `csv:"route_id"`
	OriginID      string `csv:"origin_id"`
	DestinationID string `csv:"destination_id"`
	ContainsID    string `csv:"contains_id"`
}

func (f FareRule) TableFileName() string {
	return "fare_rules.txt"
}
