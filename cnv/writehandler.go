/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
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
			// if seq, found := m[__sequence__].([]string); found {
			// 	for i, v := range seq {
			// 		fmt.Println(i, v)
			// 	}
			// }

			if vals, found := m[__values__].([]interface{}); found {
				wvals := make([]string, len(vals))
				for i, v := range vals {
					wvals[i] = decode(v)
				}
				err := encoder.Write(wvals)
				if err != nil {
					cnv.SetError(err)
				}
			}
		} else {
			// chanel closed
			return
		}
	}
}
