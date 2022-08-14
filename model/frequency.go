//Done
package model

import "errors"

type FrequencyExactTimes byte

const (
	FrequencyExactTimesFrequencyBased FrequencyExactTimes = 0
	FrequencyExactTimesScheduleBased  FrequencyExactTimes = 1
)

type Frequency struct {
	TripID      string              `csv:"trip_id"`
	StartTime   GTFSTime            `csv:"start_time,ParseStartTime"`
	EndTime     GTFSTime            `csv:"end_time,ParseEndTime"`
	HeadwaySecs uint                `csv:"headway_secs"`
	ExactTimes  FrequencyExactTimes `csv:"exact_times,ParseExactTimes"`
}

func (f Frequency) TableFileName() string {
	return "frequencies.txt"
}

func (f *Frequency) ParseStartTime(value string) (err error) {
	f.StartTime, err = NewGTFSTime(value)
	return
}

func (f *Frequency) ParseEndTime(value string) (err error) {
	f.EndTime, err = NewGTFSTime(value)
	return
}

func (f *Frequency) ParseExactTimes(value string) (err error) {
	switch value {
	case "", "0":
		f.ExactTimes = FrequencyExactTimesFrequencyBased

	case "1":
		f.ExactTimes = FrequencyExactTimesScheduleBased

	default:
		return errors.New("invalid value")
	}

	return nil
}
