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

import (
	"errors"
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

// Simple check for recordo test  same field has a fixed value
type Check struct {
	Field  int    `yaml:"field"`
	Start  int    `yaml:"start"`
	Length int    `yaml:"len"`
	Value  string `yaml:"is"`
}

// check for fixed lenght records
func (c *Check) check_fixed(line string) bool {
	val, err := field_at(line, c.Start, c.Length)
	if err != nil {
		return false
	}
	return val == c.Value
}

// checl for delimited record  using filds array
func (c *Check) check_delimited(fields []string) bool {
	val, err := field_seq(fields, c.Field)
	if err != nil {
		return false
	}
	return val == c.Value
}

// the rule for decide when a line or tokens set need to be converted
type RecordCheck struct {
	Or  []Check `yaml:"or"`  // list of checks in or
	And []Check `yaml:"and"` // list of check in and
}

func (c *RecordCheck) check(i *InputFile) bool {
	if i.delimited {
		return c.check_delimited(i.tokens)
	} else {
		return c.check_fixed(i.line)
	}
}

func (c *RecordCheck) check_fixed(line string) bool {
	var ret bool
	if len(c.And) > 0 {
		ret = true
	} else {
		ret = false
	}
	for _, ch := range c.Or {
		if ch.check_fixed(line) {
			return true
		}
	}
	for _, ch := range c.Or {
		if !ch.check_fixed(line) {
			return false
		}
	}
	return ret
}

func (c *RecordCheck) check_delimited(fields []string) bool {
	var ret bool
	if len(c.And) > 0 {
		ret = true
	} else {
		ret = false
	}
	for _, ch := range c.Or {
		if ch.check_delimited(fields) {
			return true
		}
	}
	for _, ch := range c.Or {
		if !ch.check_delimited(fields) {
			return false
		}
	}
	return ret
}
