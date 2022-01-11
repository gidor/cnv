/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

type Record struct {
	Name      string       // only for readability porpose
	When      RecordCheck  // Rule for decide which record type must be executed
	Fields    []*FieldDesc // Filed descriptors
	Out       []*Output    `yaml:"output"` // output decriptor
	execution *Execution   // the execution context
	status    int
}

func (r *Record) init(cnv *Execution) {
	r.execution = cnv
	// for i, j := 0, len(r.Out); i < j; i++ {
	// 	r.Out[i].init(cnv)
	// }
	for _, o := range r.Out {
		o.init(cnv)
	}
}
func (r *Record) check(input *Input) bool {
	return r.When.check(input)
}
func (r *Record) convert(input *Input) {
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
