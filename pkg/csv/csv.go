/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// "github.com/gidor/cnv/cnv"

func toCsvString(v interface{}) (string, error) {
	switch v := v.(type) {
	case nil:
		return "", nil
		// return "null"
	case bool:
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	// case *big.Int:
	// 	e.w.Write(v.Append(e.buf[:0], 10))
	case string:
		return v, nil
	case []interface{}:
		b, err := json.Marshal(v)
		return string(b), err
	case map[string]interface{}:
		b, err := json.Marshal(v)
		return string(b), err
	default:
		return "", errors.New(fmt.Sprintf("invalid value: %v", v))
	}
}

func fromCsvString(v string) interface{} {
	v = strings.Trim(v, "\" ")
	if b, err := strconv.ParseBool(v); err == nil {
		return b
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	if i, err := strconv.ParseInt(v, 10, 64); err == nil {
		return i
	}
	if u, err := strconv.ParseUint(v, 10, 64); err == nil {
		return u
	}
	return v
}

func Load(reader io.ReadCloser) (map[string]interface{}, error) {
	var data = make(map[string]interface{})

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
	data["data"] = make([]map[string]interface{}, 10)

	for {
		record, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, err
		}

		elementMap := make(map[string]interface{})
		for i := 0; i < len(headers); i += 1 {
			if len(record) > i {
				// elementMap[headers[i]] = strings.Trim(record[i], "\" ")
				elementMap[headers[i]] = fromCsvString(record[i])
			}
		}

		data["data"] = append((data["data"]).([]map[string]interface{}), elementMap)
		// fmt.Println(record)
	}
	return data, nil
}

func Save(data map[string]interface{}, writer io.WriteCloser, root string) error {
	encoder := csv.NewWriter(writer)
	var droot []map[string]interface{} = (data[root]).([]map[string]interface{})

	header := make([]string, 0, len(droot[0]))
	content := make([]string, 0, len(header))
	for k := range droot[0] {
		header = append(header, k)
	}
	encoder.Write(header)

	for i := range droot {
		record := droot[i]
		for j := range header {
			c, _ := toCsvString(record[(header[j])])
			content[j] = c
		}
		encoder.Write(content)

	}

	return nil

}
