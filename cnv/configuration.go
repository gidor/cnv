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

// get fixed length filed from string usngi 1 based index
func field_at(record string, start int, length int) (string, error) {
	var err error
	if start > len(record) {
		err = errors.New("start > of record length ")
		return "", err
	}
	start = start - 1
	end := start + length

	if start < 0 {
		err = errors.New("start < 1 ")
		return "", err
	}
	if (end) > len(record) {
		err = errors.New("start + len > of record length ")
		return record[start:], err
	}
	ret := record[start:end]
	return ret, err
}

// filed at  using zero based index
func field_at_0b(record string, start int, length int) (string, error) {
	var err error
	if start > len(record) {
		err = errors.New("start > of record length ")
		return "", err
	}
	end := start + length

	if start < 0 {
		err = errors.New("start < 0 ")
		return "", err
	}
	if end > len(record) {
		err = errors.New("start + len > of record length ")
		return record[start:], err
	}
	ret := record[start:end]
	return ret, err
}

// get fileds from tokens slice using 1 based index
func field_seq(fields []string, seq int) (string, error) {
	var err error
	if seq > len(fields) {
		err = errors.New("seq > num fileds")
		return "", err
	}
	seq = seq - 1
	if seq < 0 {
		err = errors.New("seq < 1")
		return "", err
	}
	ret := fields[seq]
	return ret, err
}

type Input struct {
	Name        string
	Filetype    string `yaml:"extension"`
	Strategy    string `yaml:"type"`
	Delimiter   string
	Quote       string
	Quoted      bool
	Multirecord bool
	Recordstype []*Record
	execution   *Execution
	delimited   bool
	line        string
	tokens      []string
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

func (i *Input) parse() {
	if i.execution == nil {
		i.execution.lasterr = errors.New("cannot parese if execution is not set")
		return
	}
	scanner := bufio.NewScanner(*i.execution.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		if i.delimited {
			i.tokens = strings.Split(i.line, i.Delimiter)
		}
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

type Configuration struct {
	Params    map[string]string
	Files     []*Input `yaml:",flow"`
	execution *Execution
}

func (c *Configuration) init(cnv *Execution) {
	c.execution = cnv
	// for _, i := range c.Files {
	// 	i.init(cnv)
	// }
}

func (c *Configuration) parse() {

	for _, i := range c.Files {
		if i.Filetype == c.execution.Filetype {
			i.init(c.execution)
			i.parse()
		}
	}
}
