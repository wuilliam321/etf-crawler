package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	expected := Output{
		Quality:   "A 95",
		PER:       "23.90",
		ExpRatio:  "0.03%",
		Issuer:    "Vanguard",
		Yield10Y:  "12.70%",
		Yield5Y:   "14.55%",
		Yield1Y:   "33.53%",
		DistYield: "1.34%",
		NHoldings: "505",
		AUM:       "$424.79B",
		Top10:     "30.81%",
	}
	actual := parse()
	require.Equal(t, expected, actual)
}
