/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

//a Mapping description
type Mapping struct {
	Name  string // the ouput filed name
	Alias string // the input fileld name
}

// type RecordMapping []Mapping

type Output struct {
	Name      string      `yaml:"name"`     // only for readability porpose
	Prefix    string      `yaml:"prefix"`   // prefix to use in output file
	Suffix    string      `yaml:"suffix"`   // suffix in putput file
	Mappings  [][]Mapping `yaml:"mappings"` // mapping array
	execution *Execution  // the current execution context
}

// initialize the execution context
func (o *Output) init(cnv *Execution) {
	o.execution = cnv
}

func (o *Output) write(data map[string]interface{}) {
	for _, mapping := range o.Mappings {
		mdata := make(map[string]interface{}, 100)
		mdata[__sequence__] = make([]string, len(mapping))
		mdata[__values__] = make([]interface{}, len(mapping))

		for i, m := range mapping {
			mdata[__sequence__].([]string)[i] = m.Name
			if v, ok := data[m.Alias]; ok {
				mdata[m.Name] = v
				mdata[__values__].([]interface{})[i] = v

			} else {
				mdata[m.Name] = nil
				mdata[__values__].([]interface{})[i] = nil
			}
		}
		o.execution.Write(mdata, o)
	}
}
