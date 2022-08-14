//Done
package model

import "errors"

type TransferTransferType byte

const (
	TransferTransferTypeRecommended     TransferTransferType = 0
	TransferTransferTypeTimed           TransferTransferType = 1
	TransferTransferTypeRequiredMinimum TransferTransferType = 2
	TransferTransferTypeNotPossible     TransferTransferType = 3
)

type Transfer struct {
	FromStopID      string               `csv:"from_stop_id"`
	ToStopID        string               `csv:"to_stop_id"`
	TransferType    TransferTransferType `csv:"transfer_type,ParseTransferType"`
	MinTransferTime uint                 `csv:"min_transfer_time"`
	FromRouteID     string               `csv:"from_route_id"`
	ToRouteID       string               `csv:"to_route_id"`
	FromTripID      string               `csv:"from_trip_id"`
	ToTripID        string               `csv:"to_trip_id"`
}

func (t Transfer) TableFileName() string {
	return "frequencies.txt"
}

func (t *Transfer) ParseTransferType(value string) (err error) {
	switch value {
	case "", "0":
		t.TransferType = TransferTransferTypeRecommended

	case "1":
		t.TransferType = TransferTransferTypeTimed

	case "2":
		t.TransferType = TransferTransferTypeRequiredMinimum

	case "3":
		t.TransferType = TransferTransferTypeNotPossible

	default:
		return errors.New("invalid value")
	}

	return nil
}
