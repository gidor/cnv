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

//a Mapping description
type Mapping struct {
	Name  string // the ouput filed name
	Alias string // the input fileld name
}

// type RecordMapping []Mapping

type Output struct {
	Name      string      `yaml:"name"`     // only for readability porpose
	Prefix    string      `yaml:"prefix"`   // prefix to use in output file
	Suffix    string      `yaml:"suffix"`   // suffix in putput file
	Mappings  [][]Mapping `yaml:"mappings"` // mapping array
	execution *Execution  // the current execution context
}

// initialize the execution context
func (o *Output) init(cnv *Execution) {
	o.execution = cnv
}

func (o *Output) write(data map[string]interface{}) {
	// DBG fmt.Println("Writing to ", o.Name)
	for _, mapping := range o.Mappings {
		mdata := make(map[string]interface{}, 100)
		mdata[__sequence__] = make([]string, len(mapping))
		mdata[__values__] = make([]interface{}, len(mapping))

		for i, m := range mapping {
			mdata[__sequence__].([]string)[i] = m.Name
			if v, ok := data[m.Alias]; ok {
				mdata[m.Name] = v
				mdata[__values__].([]interface{})[i] = v

			} else {
				mdata[m.Name] = nil
				mdata[__values__].([]interface{})[i] = nil
			}
		}
		o.execution.Write(mdata, o)
	}
}
