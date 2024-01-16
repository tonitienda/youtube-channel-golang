package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstDigit(t *testing.T) {
	examples := map[string]int{
		"1":          1,
		"1a":         1,
		"a1":         1,
		"2":          2,
		"avbsads2":   2,
		"2sadasdasd": 2,
		"23":         2,
		"a2fg3vf5":   2,
		"2sdad4":     2,
	}

	for in, out := range examples {
		t.Run("First Digit of "+in, func(t *testing.T) {
			assert.Equal(t, out, getFirstDigit(in))
		})
	}

}

func TestLastDigit(t *testing.T) {
	examples := map[string]int{
		"1":          1,
		"1a":         1,
		"a1":         1,
		"2":          2,
		"avbsads2":   2,
		"2sadasdasd": 2,
		"23":         3,
		"a2fg3vf5":   5,
		"2sdad4":     4,
	}

	for in, out := range examples {
		t.Run("Last Digit of "+in, func(t *testing.T) {
			assert.Equal(t, out, getLastDigit(in))
		})
	}

}

func TestCalculateTotal(t *testing.T) {
	assert.Equal(t, 142, calculateTotal(part1Text))
}
