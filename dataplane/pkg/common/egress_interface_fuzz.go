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

// FuzzParseGatewayIP tests gateway IP address parsing.
func FuzzParseGatewayIP(fuzz []byte) int {
	if len(fuzz) == 0 {
		return -1
	}
	result := parseGatewayIP(string(fuzz))
	if result.IsUnspecified() {
		return 0
	}
	return 1
}
