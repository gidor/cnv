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

import "fmt"

type Record struct {
	Name      string       // only for readability porpose
	When      RecordCheck  // Rule for decide which record type must be executed
	Fields    []*FieldDesc // Filed descriptors
	Out       []*Output    `yaml:"output"` // output decriptor
	execution *Execution   // the execution context
	status    int
}

func (r *Record) init(cnv *Execution) {
	// DBUG
	fmt.Println("init record ", r.Name)

	r.execution = cnv
	// for i, j := 0, len(r.Out); i < j; i++ {
	// 	r.Out[i].init(cnv)
	// }
	for _, o := range r.Out {
		o.init(cnv)
	}
}

func (r *Record) check(input *InputFile) bool {
	return r.When.check(input)
}

func (r *Record) convert(input *InputFile) {
	// DBUG
	fmt.Println("converting record ", r.Name)

	// TODO
	r.status = 0
	data := make(map[string]interface{}, 200)
	for _, f := range r.Fields {
		if input.delimited {
			f.parse_delimited(input.tokens, data)
		} else {
			f.parse_fixed(input.line, data)
		}
	}
	for _, o := range r.Out {
		o.write(data)
	}
}
