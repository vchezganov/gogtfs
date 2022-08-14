package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Modelable interface {
	TableFileName() string
}

func parseDate(value string) (time.Time, error) {
	return time.Parse("02012006", value)
}

// GTFSTime is number of seconds from midnight
// 00:00:53 = 53
// 09:15:33 = 33333 (9 * 3600 + 15 * 60 + 33)
type GTFSTime int

func (s GTFSTime) IsEmpty() bool {
	return s < 0
}

func (s GTFSTime) String() string {
	hours := s / (60 * 60)
	minutes := (s % (60 * 60)) / 60
	seconds := s % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func NewEmptyGTFSTime() GTFSTime {
	return -1
}

func NewGTFSTime(value string) (GTFSTime, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 3 {
		return 0, errors.New("invalid value")
	}

	var result GTFSTime
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return 0, err
		}

		result = 60*result + GTFSTime(number)
	}

	return result, nil
}
