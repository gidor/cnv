/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

var (
	ver        string = "0.1.1"
	prod       string = "cnv"
	inputFile  string = ""
	outputFile string = ""
	pretty     bool   = false
	htmlescape bool   = false
)

type internal map[string]interface{}

func getcsv(reader io.ReadCloser) internal {
	var _internal = make(internal)

	csvreader := csv.NewReader(reader)
	csvreader.Comma = ','
	csvreader.LazyQuotes = true
	csvreader.FieldsPerRecord = -1
	// csvreader.ReuseRecord = true

	headers, err := csvreader.Read()
	if err == io.EOF {
		panic(errors.New("csv streams end before data"))
	}
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(headers); i++ {
		headers[i] = strings.Trim(headers[i], "\" ")
	}
	// data := make( map[string]interfce{})
	_internal["data"] = make([]internal, 10)

	for {
		record, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		elementMap := make(internal)
		for i := 0; i < len(headers); i += 1 {
			if len(record) > i {
				// field := record[i]
				// flen := len(field)
				// field :=  strings.Trim(field, "\" ")
				// if flen > len (field){

				// }else
				elementMap[headers[i]] = strings.Trim(record[i], "\" ")
			}
		}

		_internal["data"] = append((_internal["data"]).([]internal), elementMap)
		// fmt.Println(record)
	}
	return _internal
}
