package raytracer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColourMultiplication(t *testing.T) {

	cases := []struct {
		input    Colour
		factor   Colour
		expected Colour
	}{
		{
			input:    Colour{R: 0.2, G: 0, B: 0.4},
			factor:   Colour{R: 2, G: 5, B: 3},
			expected: Colour{R: 0.4, G: 0, B: 1.0},
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("colour multiply #%d", i),
			func(t *testing.T) {
				actual := c.input.Multiply(c.factor)
				assert.InDelta(t, c.expected.R, actual.R, 0.00001)
				assert.InDelta(t, c.expected.G, actual.G, 0.00001)
				assert.InDelta(t, c.expected.B, actual.B, 0.00001)
			},
		)
	}
}

func TestColourMultiplicationN(t *testing.T) {

	cases := []struct {
		input    Colour
		factor   float64
		expected Colour
	}{
		{
			input:    Colour{R: 0.2, G: 0, B: 0.4},
			factor:   2,
			expected: Colour{R: 0.4, G: 0, B: 0.8},
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("colour multiplyN #%d", i),
			func(t *testing.T) {
				actual := c.input.MultiplyN(c.factor)
				assert.InDelta(t, c.expected.R, actual.R, 0.00001)
				assert.InDelta(t, c.expected.G, actual.G, 0.00001)
				assert.InDelta(t, c.expected.B, actual.B, 0.00001)
			},
		)
	}
}

func TestColourAddition(t *testing.T) {

	cases := []struct {
		input    Colour
		other    Colour
		expected Colour
	}{
		{
			input:    Colour{R: 0.2, G: 0, B: 0.4},
			other:    Colour{R: 0.1, G: 0, B: 0.7},
			expected: Colour{R: 0.3, G: 0, B: 1.0},
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("colour addition #%d", i),
			func(t *testing.T) {
				actual := c.input.Add(c.other)
				assert.InDelta(t, c.expected.R, actual.R, 0.00001)
				assert.InDelta(t, c.expected.G, actual.G, 0.00001)
				assert.InDelta(t, c.expected.B, actual.B, 0.00001)
			},
		)
	}
}
