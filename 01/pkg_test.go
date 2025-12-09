package main

import "testing"

func TestDial(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name string
			v    int
			exp  int
		}{
			{
				name: "ValidValue",
				v:    5,
				exp:  5,
			},
			{
				name: "TooBig",
				v:    305,
				exp:  5,
			},
			{
				name: "TooSmall",
				v:    -305,
				exp:  5,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				dv := NewDialValue(tt.v)
				if tt.exp != int(dv) {
					t.Fatalf("Expected %v, got %v", tt.exp, dv)
				}
			})
		}
	})

	t.Run("Add", func(t *testing.T) {
		dialValue := NewDialValue(50)
		tests := []struct {
			name      string
			v         int
			exp       int
			expClicks int
		}{
			{
				name:      "InRange",
				v:         5,
				exp:       55,
				expClicks: 0,
			},
			{
				name:      "PositiveOverflow",
				v:         55,
				exp:       5,
				expClicks: 1,
			},
			{
				name:      "NegativeOverflow",
				v:         -55,
				exp:       95,
				expClicks: 1,
			},
			{
				name:      "MassiveOverflow",
				v:         500,
				exp:       50,
				expClicks: 5,
			},
			{
				name:      "MassiveUnderflow",
				v:         -500,
				exp:       50,
				expClicks: 5,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				dv, clicks := dialValue.Add(tt.v)
				if tt.exp != int(dv) {
					t.Fatalf("Expected %v, got %v", tt.exp, dv)
				}
				if tt.expClicks != clicks {
					t.Fatalf("Expected %d clicks, got %d", tt.expClicks, clicks)
				}
			})
		}
	})

	t.Run("EdgeCases", func(t *testing.T) {
		tests := []struct {
			name      string
			start     int
			v         int
			exp       int
			expClicks int
		}{
			{
				name:      "EndAtZero",
				start:     48,
				v:         52,
				exp:       0,
				expClicks: 1,
			},
			{
				name:      "EndAtZeroNegative",
				start:     5,
				v:         -5,
				exp:       0,
				expClicks: 1,
			},
			{
				name:      "MultipleRotationsEndAtZero",
				start:     50,
				v:         250,
				exp:       0,
				expClicks: 3,
			},
			{
				name:      "MultipleNegativeRotationsEndAtZero",
				start:     50,
				v:         -250,
				exp:       0,
				expClicks: 3,
			},
			{
				name:      "FromZero",
				start:     0,
				v:         5,
				exp:       5,
				expClicks: 0,
			},
			{
				name:      "FromZeroNegative",
				start:     0,
				v:         -5,
				exp:       95,
				expClicks: 0,
			},
			{
				name:      "FromZeroMultipleRotations",
				start:     0,
				v:         1050,
				exp:       50,
				expClicks: 10,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				dialValue := NewDialValue(tt.start)
				dv, clicks := dialValue.Add(tt.v)
				if tt.exp != int(dv) {
					t.Fatalf("Expected position %v, got %v", tt.exp, dv)
				}
				if tt.expClicks != clicks {
					t.Fatalf("Expected %d clicks, got %d", tt.expClicks, clicks)
				}
			})
		}
	})
}
