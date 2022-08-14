//Done
package model

import (
	"errors"
)

type FareAttributePaymentMethod byte

const (
	FareAttributePaymentMethodOnBoard     FareAttributePaymentMethod = 0
	FareAttributePaymentMethodBeforeBoard FareAttributePaymentMethod = 1
)

type FareAttributeTransfers int8

const (
	FareAttributeTransfersNoTransfers        FareAttributeTransfers = 0
	FareAttributeTransfersOneTransfer        FareAttributeTransfers = 1
	FareAttributeTransfersTwoTransfers       FareAttributeTransfers = 2
	FareAttributeTransfersUnlimitedTransfers FareAttributeTransfers = -1 // TODO: Maybe convert to `string` type
)

type FareAttribute struct {
	FareID           string                     `csv:"fare_id"`
	Price            float64                    `csv:"price"`
	CurrencyType     string                     `csv:"currency_type"`
	PaymentMethod    FareAttributePaymentMethod `csv:"payment_method,ParsePaymentMethod"`
	Transfers        FareAttributeTransfers     `csv:"transfers,ParseTransfers"`
	AgencyID         string                     `csv:"agency_id"`
	TransferDuration uint                       `csv:"transfer_duration"`
}

func (f FareAttribute) TableFileName() string {
	return "fare_attributes.txt"
}

func (f *FareAttribute) ParsePaymentMethod(value string) error {
	switch value {
	case "0":
		f.PaymentMethod = FareAttributePaymentMethodOnBoard

	case "1":
		f.PaymentMethod = FareAttributePaymentMethodBeforeBoard

	default:
		return errors.New("invalid value")
	}

	return nil
}

func (f *FareAttribute) ParseTransfers(value string) error {
	switch value {
	case "0":
		f.Transfers = FareAttributeTransfersNoTransfers

	case "1":
		f.Transfers = FareAttributeTransfersOneTransfer

	case "2":
		f.Transfers = FareAttributeTransfersTwoTransfers

	case "":
		f.Transfers = FareAttributeTransfersUnlimitedTransfers

	default:
		return errors.New("invalid value")
	}

	return nil
}
