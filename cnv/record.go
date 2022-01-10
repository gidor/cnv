/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dateregexp, fixeddecregexp *regexp.Regexp

func init() {
	dateregexp = regexp.MustCompile(`^\s*date\((?P<format>.+)\)$`)
	fixeddecregexp = regexp.MustCompile(`^\s*fixed\((?P<format>[0-9]+)\)$`)
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

type RecordCheck struct {
	Or  []Check `yaml:"or"`
	And []Check `yaml:"and"`
}

func (c *RecordCheck) check(i *Input) bool {
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

type Parsingtype uint8

const (
	Unizialized  Parsingtype = iota // zero value
	Integers                        // parse integerner
	Dates                           // date
	Strings                         //string
	Decimal                         // parsed decimal
	Fixeddecimal                    // decimal with no fot end fixed decimal digit
	Booleans                        // boleans value
)

type FieldDesc struct {
	Name      string `yaml:"name"`
	Seq       int    `yaml:"seq"`
	Parsetype string `yaml:"parsetype"`
	Len       int16  `yaml:"len"`
	Start     int16  `yaml:"start"`
	Note      string `yaml:"note"`
	parse     Parsingtype
	div       int
	yyyy      int
	yy        int
	mm        int
	dd        int
}

func (f *FieldDesc) parsedet() {
	if f.parse == Unizialized {
		switch p := strings.ToLower(f.Parsetype); {
		case p == "int", p == "numeric":
			f.parse = Integers
		case p == "dec", p == "real", p == "decimal":
			f.parse = Decimal
		case dateregexp.MatchString(p):
			mm := dateregexp.FindStringSubmatch(p)
			id := dateregexp.SubexpIndex("format")
			format := mm[id]
			f.yyyy, f.yy, f.mm, f.dd = -1, -1, -1, -1
			f.yyyy = int(strings.Index(format, "aaaa"))
			if f.yyyy < 0 {
				f.yyyy = int(strings.Index(format, "yyyy"))
			}
			if f.yyyy < 0 {
				f.yy = int(strings.Index(format, "yy"))
				if f.yy < -1 {
					f.yy = int(strings.Index(format, "aa"))
				}
			}
			f.mm = int(strings.Index(format, "mm"))
			f.dd = int(strings.Index(format, "dd"))
			if f.dd < 0 {
				f.dd = int(strings.Index(format, "gg"))
			}
			f.parse = Dates
		case fixeddecregexp.MatchString(p):
			format := fixeddecregexp.FindStringSubmatch(p)[fixeddecregexp.SubexpIndex("format")]
			d, _ := strconv.ParseInt(format, 10, 32)
			f.div = int(math.Pow10(int(d)))
			f.parse = Fixeddecimal
		default:
			f.parse = Strings
		}
	}
}

func (f *FieldDesc) parsevalue(val string) interface{} {
	f.parsedet()

	switch f.parse {
	case Integers:
		v, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return v
		} else {
			return nil
		}
	case Fixeddecimal:
		v, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return float64(v) / float64(f.div)
		} else {
			return nil
		}
	case Dates:
		var y, m, d int64
		var err error
		vy, _ := field_at_0b(val, int(f.mm), 2)
		m, err = strconv.ParseInt(vy, 10, 32)
		if err != nil {
			return nil
		}
		if f.yyyy > -1 {
			vy, _ = field_at_0b(val, int(f.yyyy), 4)
			y, err = strconv.ParseInt(vy, 10, 32)
		} else {
			vy, _ := field_at_0b(val, int(f.yy), 2)
			y, err = strconv.ParseInt(vy, 10, 32)
			if y < 50 {
				y = 1900 + y
			} else {
				y = 2000 + y
			}
		}
		if err != nil {
			return nil
		}
		vy, _ = field_at_0b(val, int(f.dd), 2)
		d, err = strconv.ParseInt(vy, 10, 32)
		if err != nil {
			return nil
		}
		return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
	case Decimal:
		v, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return v
		} else {
			return nil
		}
	case Booleans:
		v, err := strconv.ParseBool(val)
		if err == nil {
			return v
		} else {
			return nil
		}

	default:
		return val
	}
}

// parse fixed lenghth field
func (f *FieldDesc) parse_fixed(line string, target map[string]interface{}) {
	val, _ := field_at(line, int(f.Start), int(f.Len))
	target[f.Name] = f.parsevalue(val)
}

// parse delimited field
func (f *FieldDesc) parse_delimited(fields []string, target map[string]interface{}) {
	val, _ := field_seq(fields, int(f.Seq))
	target[f.Name] = f.parsevalue(val)
}

type Mapping struct {
	Name  string
	Alias string
}

// type RecordMapping []Mapping

type Output struct {
	Name      string      `yaml:"name"`
	Prefix    string      `yaml:"prefix"`
	Suffix    string      `yaml:"suffix"`
	Mappings  [][]Mapping `yaml:"mappings"`
	execution *Execution
}

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

type Record struct {
	Name      string
	When      RecordCheck
	Fields    []*FieldDesc
	Out       []*Output `yaml:"output"`
	execution *Execution
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
