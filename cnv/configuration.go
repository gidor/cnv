/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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

type Output struct {
	Name       string              `yaml:"name"`
	Prefix     string              `yaml:"prefix"`
	Suffix     string              `yaml:"suffix"`
	Mappings   []map[string]string `yaml:"mappings"`
	conversion *Conversion
}

func (o *Output) init(cnv *Conversion) {
	o.conversion = cnv
}

type Record struct {
	Name       string
	When       RecCheck
	Fields     []FieldDesc
	Out        []Output
	conversion *Conversion
}

func (r *Record) init(cnv *Conversion) {
	r.conversion = cnv
	for _, o := range r.Out {
		o.init(cnv)
	}
}

type Input struct {
	Name        string
	Filetype    string `yaml:"extension"`
	Strategy    string `yaml:"type"`
	Delimiter   string
	Quote       string
	Quoted      bool
	Multirecord bool
	Recordstype []Record
	conversion  *Conversion
	line string
	tokens []string
}

func (i *Input) init(cnv *Conversion) {
	i.conversion = cnv
	for _, r := range i.Recordstype {
		r.init(cnv)
	}
}

func (i *Input) parse() {
	if i.conversion == nil {
		i.conversion.lasterr = errors.New("cannot parese if conversion is not set")
		return
	}
	scanner := bufio.NewScanner(*i.conversion.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		if i.Delimiter{
			i.tokens= i.line
		} 
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	scanne

}

type Configuration struct {
	Files      []Input `yaml:",flow"`
	conversion *Conversion
}

func (c *Configuration) init(cnv *Conversion) {
	c.conversion = cnv
	// for _, i := range c.Files {
	// 	i.init(cnv)
	// }
}

func (c *Configuration) parse() {

	for _, i := range c.Files {
		if i.Filetype == c.conversion.Filetype {
			i.init(c.conversion)
			i.parse()
		}
	}
}
