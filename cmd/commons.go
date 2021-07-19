/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/

package cmd

import (
	"io"
	"os"
)

var (
	ver        string = "0.2.3"
	prod       string = "cnv"
	inputFile  string = ""
	outputFile string = ""
	pretty     bool   = false
	htmlescape bool   = false
)

func openioout(reader *io.ReadCloser, writer *io.WriteCloser) {
	if inputFile != "" {
		if in, err := os.OpenFile(inputFile, os.O_RDONLY, 0755); err != nil {
			panic(err)
		} else {
			*reader = in
		}
	}

	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			*writer = out
		}
	}

}
