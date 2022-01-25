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

import "strings"

type JValue interface{}

type Meta map[string]JValue

type Encoding int

const (
	Delimited Encoding = iota // delimited data file
	Fixlength                 // Fixed lenghr data file
	Named                     // Named properties
	Csv                       // Csv
	Yaml                      // Yaml
	Json                      // Json
	Toml                      // Toml
	Unknown                   // Toml
)

var extensions = map[Encoding][]string{
	Csv:       {"csv"},
	Yaml:      {"yaml", "yml"},
	Json:      {"json"},
	Delimited: {"dat"},
	Fixlength: {"dat"},
	Named:     {"dat"},
}

var names = map[Encoding]string{
	Delimited: "Delimited",
	Fixlength: "Fixlength",
	Named:     "Named",
	Csv:       "Csv",
	Yaml:      "Yaml",
	Json:      "Json",
}

func Strategy(name string) Encoding {
	name = strings.ToLower(name)
	for k, v := range names {
		v = strings.ToLower(v)
		if v == name {
			return k
		}
	}
	return Unknown
}

type WriteHandler func(ch chan (map[string]interface{}), cnv Execution, prefix string, suffix string)

func (c Encoding) AsString() string {
	name, ok := names[c]
	if ok {
		return name
	}
	return "Unknown"
}

const (
	__sequence__ = "__sequence__"
	__values__   = "__values__"
)
