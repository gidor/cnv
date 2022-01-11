/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

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

var extension = map[Encoding][]string{
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
	for k, v := range names {
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
