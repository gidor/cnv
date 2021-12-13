/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

type JValue interface{}

type Meta map[string]JValue

type Format int8

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

type Encoding int8

const (
	Delimited Encoding = iota
	Fixlength Encoding = iota
)

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
