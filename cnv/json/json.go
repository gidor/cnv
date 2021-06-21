/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package json

import (
	"encoding/json"
	"io"
)

func init() {

}

func Load(reader io.ReadCloser) (map[string]interface{}, error) {

	// var data = make(map[string]interface{})
	var data = &map[string]interface{}{}

	err := json.NewDecoder(reader).Decode(data)

	return *data, err

}

func Save(data *map[string]interface{}, writer io.WriteCloser, pretty bool, htmlescape bool) error {

	encoder := json.NewEncoder(writer)
	if pretty {
		encoder.SetIndent("", "  ")
	} else {
		encoder.SetIndent("", "")
	}
	// set escape html
	encoder.SetEscapeHTML(htmlescape)

	err := encoder.Encode(data)
	return err

}
