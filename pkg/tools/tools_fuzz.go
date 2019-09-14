// +build gofuzz

package tools

// FuzzParseAnnotationValu tests annotation value parsing.
func FuzzParseAnnotationValue(fuzz []byte) int {
	_, err := ParseAnnotationValue(string(fuzz))
	if err != nil {
		return 0
	}
	return 1
}
