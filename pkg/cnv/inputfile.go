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
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gidor/cnv/pkg/yaml"
)

// InputFile describe an input data file
type InputFile struct {
	Name        string     // only for readability porpose
	Filetype    string     `yaml:"extension"` // file extension my be used  when  we need decide which input apply to a file
	Strategy    string     `yaml:"type"`      // encoding  formate and strategy see Encoding
	Delimiter   string     // optional only when starategy is delimited the delimiter
	Quote       string     // optionaly quoting character only when starategy is delimited
	Quoted      bool       // flag true if quoted
	Multirecord bool       // flag true if multiple record type are present
	Recordstype []*Record  // a slice of record descrtiption
	execution   *Execution // the current execution context
	delimited   bool       // flag true if delimited
	line        string     // last line read
	tokens      []string   // la tokens read from line
	// data        map[string]interface{} // data when yaml or json
}

func (i *InputFile) init(cnv *Execution) {
	// DBG fmt.Println("init file ", i.Name)

	i.Delimiter = strings.Trim(i.Delimiter, " ")
	if len(i.Delimiter) > 0 {
		i.delimited = true
	}
	i.execution = cnv
	for _, r := range i.Recordstype {
		r.init(cnv)
	}
}

func (i *InputFile) parse_yaml() {
	data, err := yaml.Load(*i.execution.Reader)
	if err == nil {
		for _, rt := range i.Recordstype {
			for _, o := range rt.Out {
				o.write(data)
			}
		}
	} else {
		i.execution.lasterr = err
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	if strings.HasPrefix(i.Filetype, ".") {
		i.Filetype = strings.Replace(i.Filetype, ".", "", 1)
	}

}

func (i *InputFile) parse_delimited() {
	scanner := bufio.NewScanner(*i.execution.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		i.tokens = strings.Split(i.line, i.Delimiter)
		for _, rt := range i.Recordstype {
			if i.Multirecord {
				if rt.check(i) {
					rt.convert(i)
				}
			} else {
				rt.convert(i)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		i.execution.lasterr = err
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}
func (i *InputFile) parse_fixed() {
	scanner := bufio.NewScanner(*i.execution.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i.line = scanner.Text()
		for _, rt := range i.Recordstype {
			if i.Multirecord {
				if rt.check(i) {
					rt.convert(i)
				}
			} else {
				rt.convert(i)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		i.execution.lasterr = err
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}

func (i *InputFile) parse() {
	// DBG fmt.Println("Parsing file ", i.Name)

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
