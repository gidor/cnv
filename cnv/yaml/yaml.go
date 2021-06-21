/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package yaml

import (
	"io"

	"gopkg.in/yaml.v3"
)

func init() {

}

func Load(reader io.ReadCloser) (map[string]interface{}, error) {

	var data = &map[string]interface{}{}

	err := yaml.NewDecoder(reader).Decode(data)

	return *data, err

}

func Save(data *map[string]interface{}, writer io.WriteCloser, pretty bool, htmlescape bool) error {

	encoder := yaml.NewEncoder(writer)
	if pretty {
		encoder.SetIndent(4)
	} else {
		encoder.SetIndent(2)
	}

	err := encoder.Encode(data)
	return err

}
