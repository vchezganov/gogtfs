package model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGTFSTimeNew(t *testing.T) {
	tests := []struct {
		value  string
		result GTFSTime
	}{
		{"00:00:24", 24},
		{"00:01:00", 60},
		{"00:02:05", 125},
		{"01:00:00", 3600},
		{"01:02:05", 3725},
		{"01:02:15", 3735},
	}

	for _, test := range tests {
		actual, err := NewGTFSTime(test.value)
		require.NoError(t, err)
		require.Equal(t, test.result, actual)
	}
}

func TestGTFSTimeString(t *testing.T) {
	tests := []struct {
		value  GTFSTime
		result string
	}{
		{24, "00:00:24"},
		{60, "00:01:00"},
		{125, "00:02:05"},
		{3600, "01:00:00"},
		{3725, "01:02:05"},
		{3735, "01:02:15"},
		{33333, "09:15:33"},
	}

	for _, test := range tests {
		actual := test.value.String()
		require.Equal(t, test.result, actual)
	}
}
