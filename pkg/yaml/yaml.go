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

package yaml

import (
	"io"

	"gopkg.in/yaml.v3"
)

func init() {

}

func Load(reader io.Reader) (map[string]interface{}, error) {

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
