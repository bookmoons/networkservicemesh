// +build gofuzz

package common

import (
	"bufio"
	"bytes"
)

// FuzzParseProcFile tests proc file parsing.
func FuzzParseProcFile(fuzz []byte) int {
	reader := bufio.NewReader(bytes.NewReader(fuzz))
	_, _, err := parseProcFile(reader)
	if err != nil {
		return 0
	}
	return 1
}
