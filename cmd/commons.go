/*
Copyright © 2021 - 2022 Gianni Doria (gianni.doria@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package cmd

import (
	"io"
	"os"
)

var (
	ver        string = "0.2.3"
	prod       string = "cnv"
	inputFile  string
	outputFile string
	config     string
	delimiter  string
	outputDir  string = "."
	desctype   string = "."
	pretty     bool   = false
	htmlescape bool   = false
)

func openioout(reader *io.ReadCloser, writer *io.WriteCloser) {
	openin(reader)
	// if inputFile != "" {
	// 	if in, err := os.OpenFile(inputFile, os.O_RDONLY, 0755); err != nil {
	// 		panic(err)
	// 	} else {
	// 		*reader = in
	// 	}
	// }

	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			*writer = out
		}
	}

}

func openin(reader *io.ReadCloser) {
	if inputFile != "" {
		if in, err := os.OpenFile(inputFile, os.O_RDONLY, 0755); err != nil {
			panic(err)
		} else {
			*reader = in
		}
	}

}
