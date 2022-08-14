//Done
package model

import (
	"errors"
	"strconv"
)

type PathwayPathwayMode byte

const (
	PathwayPathwayModeWalkway    PathwayPathwayMode = 1
	PathwayPathwayModeStairs     PathwayPathwayMode = 2
	PathwayPathwayModeTravelator PathwayPathwayMode = 3
	PathwayPathwayModeEscalator  PathwayPathwayMode = 4
	PathwayPathwayModeElevator   PathwayPathwayMode = 5
	PathwayPathwayModeFareGate   PathwayPathwayMode = 6
	PathwayPathwayModeExitGate   PathwayPathwayMode = 7
)

type PathwayIsBidirectional bool

const (
	PathwayIsBidirectionalNo  PathwayIsBidirectional = false
	PathwayIsBidirectionalYes PathwayIsBidirectional = true
)

type Pathway struct {
	PathwayID            string                 `csv:"pathway_id"`
	FromStopID           string                 `csv:"from_stop_id"`
	ToStopID             string                 `csv:"to_stop_id"`
	PathwayMode          PathwayPathwayMode     `csv:"pathway_mode,ParsePathwayMode"`
	IsBidirectional      PathwayIsBidirectional `csv:"is_bidirectional,ParseIsBidirectional"`
	Length               float64                `csv:"length"`
	TraversalTime        uint                   `csv:"traversal_time"`
	StairCount           uint                   `csv:"stair_count"`
	MaxSlope             float64                `csv:"max_slope"`
	MinWidth             float64                `csv:"min_width"`
	SignpostedAs         string                 `csv:"signposted_as"`
	ReversedSignpostedAs string                 `csv:"reversed_signposted_as"`
}

func (p Pathway) TableFileName() string {
	return "pathways.txt"
}

func (p *Pathway) ParsePathwayMode(value string) error {
	number, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	if !(1 <= number && number <= 7) {
		return errors.New("invalid value")
	}

	p.PathwayMode = PathwayPathwayMode(number)
	return nil
}

func (p *Pathway) ParseIsBidirectional(value string) error {
	switch value {
	case "0":
		p.IsBidirectional = PathwayIsBidirectionalNo

	case "1":
		p.IsBidirectional = PathwayIsBidirectionalYes

	default:
		return errors.New("invalid value")
	}

	return nil
}
