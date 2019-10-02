package util

import (
	"golang.org/x/crypto/ssh/terminal"
)

func GetDimensions() (int, int, error) {
	return terminal.GetSize(0)
}
