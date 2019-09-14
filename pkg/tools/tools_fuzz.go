// +build gofuzz

package tools

// FuzzParseAnnotationValue tests annotation value parsing.
func FuzzParseAnnotationValue(fuzz []byte) int {
	_, err := ParseAnnotationValue(string(fuzz))
	if err != nil {
		return 0
	}
	return 1
}

// FuzzParseKVString tests KV string parsing.
func FuzzParseKVString(fuzz []byte) int {
	input, sep, kvsep, ok := decodeKVStringFuzz(fuzz)
	if !ok {
		return -1
	}
	ParseKVStringToMap(input, sep, kvsep)
	return 1
}

func decodeKVStringFuzz(fuzz []byte) (
	input string,
	sep string,
	kvsep string,
	ok bool,
) {
	sep, fuzz, ok = extractFuzzString(fuzz)
	if !ok {
		return
	}
	kvsep, fuzz, ok = extractFuzzString(fuzz)
	if !ok {
		return
	}
	input = string(fuzz)
	ok = true
	return
}

func extractFuzzString(fuzz []byte) (
	value string,
	rest []byte,
	ok bool,
) {
	if len(fuzz) < 2 {
		// Invalid string encoding
		return
	}
	length := int(fuzz[0])
	if length == 0 {
		// Invalid length
		return
	}
	if len(fuzz) < (length + 1) {
		// Insufficient fuzz
		return
	}
	value = string(fuzz[1 : length+1])
	if len(fuzz) == (length + 1) {
		// Consumed all fuzz
		rest = []byte{}
	} else {
		// More fuzz
		rest = fuzz[length+1:]
	}
	ok = true
	return
}
