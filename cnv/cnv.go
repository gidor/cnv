/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

import (
	"io"
)

type JValue interface{}

type Meta map[string]JValue

type Format int8

const (
	Csv  Format = 0
	Yaml Format = 1
	Json Format = 2
)

type Conversion struct {
	Reader     *io.ReadCloser
	Outdir     string
	Filetype   string
	Cfgname    string
	Cnvformat  Format
	Pretty     bool
	Htm_escape bool
}

type Check struct {
	Field int    `yaml:"field"`
	Name  string `yaml:"name"`
	Value string `yaml:"is"`
}
type RecCheck struct {
	Or  []Check `yaml:"or"`
	And []Check `yaml:"and"`
}

type FieldDesc struct {
	Name      string `yaml:"name"`
	Seq       string `yaml:"seq"`
	Parsetype string `yaml:"parsetype"`
	Len       int16  `yaml:"len"`
	Start     int16  `yaml:"start"`
	Note      string `yaml:"note"`
}

type OutDesc struct {
	Name     string              `yaml:"name"`
	Prefix   string              `yaml:"prefix"`
	Suffix   string              `yaml:"suffix"`
	Mappings []map[string]string `yaml:"mappings"`
}

type RecordDesc struct {
	Name   string
	When   RecCheck
	Fields []FieldDesc
	Out    []OutDesc
}
type CnvDesc struct {
	Name        string
	Extension   string
	Strategy    string `yaml:"type"`
	Delimiter   string
	Quote       string
	Quoted      bool
	Multirecord bool
	Recordstype []RecordDesc
}
type CnvCfg struct {
	Files []CnvDesc `yaml:",flow"`
}
