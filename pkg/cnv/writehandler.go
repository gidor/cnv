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

package cnv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"time"

	"os"
	"path"

	goyaml "gopkg.in/yaml.v3"
)

type paramsty struct {
	delimiter  rune
	nulrender  string
	dateformat string
}

var params paramsty

func init() {
	params = paramsty{'|', "N", "2006-01-02"}
}

// {'|',"N","2006-01-02"}

func yamlWriteHandler(ch chan (map[string]interface{}), cnv *Execution, prefix string, suffix string) {
	outputFile := path.Join(cnv.Outdir, prefix+path.Base(cnv.Infile)+suffix+".txt")
	var writer *os.File
	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}
	defer writer.Close()
	encoder := goyaml.NewEncoder(writer)

	for {
		m, ok := <-ch
		if ok {
			err := encoder.Encode(m)
			if err != nil {
				cnv.SetError(err)
			}
		} else {
			return
		}
	}
}

func jsonWriteHandler(ch chan (map[string]interface{}), cnv *Execution, prefix string, suffix string) {

	outputFile := path.Join(cnv.Outdir, prefix+path.Base(cnv.Infile)+suffix+".txt")
	var writer *os.File
	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	for {
		m, ok := <-ch
		if ok {
			err := encoder.Encode(m)
			if err != nil {
				cnv.SetError(err)
			}
		} else {
			return
		}
	}
}
func decode(v interface{}) string {

	switch v.(type) {
	case int, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case string:
		return v.(string)
	case bool:
		return fmt.Sprintf("%t", v)
	case nil:
		return params.nulrender
	case time.Time:
		t := v.(time.Time)
		return t.Format(params.dateformat)
	default:
		return ""
	}
}

func delWriteHandler(ch chan (map[string]interface{}), cnv *Execution, prefix string, suffix string) {
	// DBG
	fmt.Println("Init delmited writee ", prefix)
	outputFile := path.Join(cnv.Outdir, prefix+path.Base(cnv.Infile)+suffix+".txt")
	var writer *os.File
	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}
	defer writer.Close()
	encoder := csv.NewWriter(writer)
	encoder.Comma = params.delimiter // cnv.Delimiter
	encoder.UseCRLF = true
	for {
		m, ok := <-ch
		if ok {
			// DBG
			fmt.Println("Write record ", prefix)

			if vals, found := m[__values__].([]interface{}); found {
				wvals := make([]string, len(vals))
				for i, v := range vals {
					wvals[i] = decode(v)
				}
				err := encoder.Write(wvals)
				encoder.Flush()
				if err != nil {
					cnv.SetError(err)
				}
			}
		} else {
			// DBG
			fmt.Println("chanel closed", prefix)
			encoder.Flush()
			return
		}
	}
}
