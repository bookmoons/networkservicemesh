// +build gofuzz

package caddyfile

import (
	"io/ioutil"
	"os"
)

// FuzzParseCaddyfile tests caddyfile parsing.
func FuzzParseCaddyfile(fuzz []byte) int {
	file, err := ioutil.TempFile("", "FuzzParseCaddyfile")
	if err != nil {
		panic(err)
	}
	path := file.Name()
	defer os.Remove(path)
	_, err = file.Write(fuzz)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	caddyfile := NewCaddyfile(path)
	caddyfile.String()
	return 1
}
