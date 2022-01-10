/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

type JValue interface{}

type Meta map[string]JValue

type Format uint8

const (
	Csv  Format = iota // Csv
	Yaml Format = iota // yaml
	Json Format = iota // Json
)

func (c Format) AsString() string {
	switch c {
	case Csv:
		return "Csv"
	case Yaml:
		return "Yaml"
	case Json:
		return "Json"
	default:
		return "Unkwon"
	}
}

type Encoding uint8

const (
	Delimited Encoding = iota
	Fixlength Encoding = iota
	Named     Encoding = iota
)

type WriteHandler func(ch chan (map[string]interface{}), cnv Execution, prefix string, suffix string)

func (c Encoding) AsString() string {
	switch c {
	case Delimited:
		return "Delimited"
	case Fixlength:
		return "Fixlength"
	default:
		return "Unkwon"
	}
}

const (
	__sequence__ = "__sequence__"
	__values__   = "__values__"
)
