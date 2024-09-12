package utils

import (
	"strconv"
)

type NumberFormatter interface {
	ToTwoDigits(int) string
}

type numberFormatter struct{}

func NewNumberFormatter() NumberFormatter {
	return numberFormatter{}
}

func (nf numberFormatter) ToTwoDigits(input int) string {
	number := strconv.Itoa(input % 100)
	if input < 10 {
		number = "0" + number
	}
	return number
}
