/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Input describe an input data file
type Input struct {
	Name        string     // only for readability porpose
	Filetype    string     `yaml:"extension"` // file extension my be used  when  we need decide which input apply to a file
	Strategy    string     `yaml:"type"`      // encoding  formate and strategy see Encoding
	Delimiter   string     // optional when starategy is delimited the delimiter
	Quote       string     // optionaly quting character
	Quoted      bool       // flag true if quoted
	Multirecord bool       // flag true if multiple record type are present
	Recordstype []*Record  // a slice of record descrtiption
	execution   *Execution // the current execution context
	delimited   bool       // flag true if delimited
	line        string     // last line read
	tokens      []string   // la tokens read from line
}

func (i *Input) init(cnv *Execution) {
	i.Delimiter = strings.Trim(i.Delimiter, " ")
	if len(i.Delimiter) > 0 {
		i.delimited = true
	}
	i.execution = cnv
	for _, r := range i.Recordstype {
		r.init(cnv)
	}
}
func (i *Input) parse_yaml() {

}
func (i *Input) parse_delimited() {
	scanner := bufio.NewScanner(*i.execution.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		i.tokens = strings.Split(i.line, i.Delimiter)
		for _, rt := range i.Recordstype {
			if rt.check(i) {
				rt.convert(i)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		i.execution.lasterr = err
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}
func (i *Input) parse_fixed() {
	scanner := bufio.NewScanner(*i.execution.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		for _, rt := range i.Recordstype {
			if rt.check(i) {
				rt.convert(i)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		i.execution.lasterr = err
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}

func (i *Input) parse() {
	if i.execution == nil {
		i.execution.lasterr = errors.New("cannot parese if execution is not set")
		return
	}
	switch enc := Strategy(i.Strategy); enc {
	case Yaml:
		i.parse_yaml()
	case Delimited:
		i.parse_delimited()
	case Fixlength:
		i.parse_fixed()
	}
}
