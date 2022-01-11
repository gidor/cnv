/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

import (
	"math"
	"strconv"
	"strings"
	"time"
)

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
	Name      string      `yaml:"name"`      // name
	Seq       int         `yaml:"seq"`       // sequence
	Parsetype string      `yaml:"parsetype"` // a Parsingtype
	Len       int16       `yaml:"len"`       // field length
	Start     int16       `yaml:"start"`     // field start when fixed
	Note      string      `yaml:"note"`      // decription
	parse     Parsingtype // Parsingtype
	div       int         //
	yyyy      int         // 4 digit year
	yy        int         // 2 digit year
	mm        int         // 2 digit month
	dd        int         // 2 digit day
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
