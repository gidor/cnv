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
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	goyaml "gopkg.in/yaml.v3"
)

type paramsty struct {
	delimiter  rune
	nulrender  string
	dateformat string
}

var params paramsty

func init() {
	params = paramsty{'|', "N", "2006-01-02"}
}

type Execution struct {
	Reader      *io.ReadCloser                           // input file reader
	Infile      string                                   // input file name
	Outdir      string                                   // output dir i  wich decoded file mast be written
	cfgname     string                                   // Confguration file name
	Filetype    string                                   // file type in input decription use extension by default
	Cnvformat   Encoding                                 // Output File Format Csv, Yaml, Json
	Delimiter   rune                                     // delimiter when Cnvformat is Csv
	lasterr     error                                    // last error uccurred
	cfg         Configuration                            // configuration
	initialized bool                                     // true only if we are ready to start converting
	outwriter   map[string]chan (map[string]interface{}) // named chanel to the output writers

}

func (c *Execution) Cfgname() string {
	return c.cfgname
}

// close all channels so handler will terminate and close created files
func (c *Execution) Close() {
	for k, v := range c.outwriter {
		close(v)
		fmt.Fprintf(os.Stderr, "closing %s", k)
	}
}

func (c *Execution) end() {
	//DBG fmt.Println("ENDING")
	for name, _ := range c.outwriter {
		//DBG fmt.Println("closing", name)
		close(c.outwriter[name])
		delete(c.outwriter, name)
	}
}

func (c *Execution) Write(data map[string]interface{}, on *Output) {
	ch, ok := c.outwriter[on.Name]

	if !ok {
		ch = make(chan (map[string]interface{}))
		c.outwriter[on.Name] = ch
		switch c.Cnvformat {
		case Yaml:
			go yamlWriteHandler(ch, c, on.Prefix, on.Suffix)
		case Json:
			go jsonWriteHandler(ch, c, on.Prefix, on.Suffix)
		case Csv:
			if c.Delimiter == 0 {
				c.Delimiter = ','
			}
			go delWriteHandler(ch, c, on.Prefix, on.Suffix)
		default:
			if c.Delimiter == 0 {
				c.Delimiter = '|'
			}
			go delWriteHandler(ch, c, on.Prefix, on.Suffix)

		}
	}
	ch <- data
}

func (c *Execution) Error() string {
	return c.lasterr.Error()
}

//error setter
func (c *Execution) SetError(err error) {
	c.lasterr = err
}

func (c *Execution) init() {
	c.defaults()
	if c.initialized {
		return
	}
	b, err := ioutil.ReadFile(c.cfgname)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	c.cfg = Configuration{}

	err = goyaml.Unmarshal(b, &c.cfg)
	if err != nil {
		c.lasterr = err
		panic(err)
	}
	if v, ok := c.cfg.Params["delimiter"]; ok {
		params.delimiter = []rune(v)[0]
	} else {
		if c.Delimiter > 0 {
			params.delimiter = c.Delimiter
		}
	}
	if v, ok := c.cfg.Params["dateformat"]; ok {
		params.dateformat = v
	}
	if v, ok := c.cfg.Params["nulrender"]; ok {
		params.nulrender = v
	}

	c.cfg.init(c)
	c.initialized = true
}

func GetConversion(reader *io.ReadCloser) *Execution {
	// c := Execution{}
	c := new(Execution)
	c.Reader = reader
	// c.Outdir = outdir
	// c.cfgname = cfgname
	// c.Cnvformat = cnvformat
	// c.Filetype = filetype
	c.lasterr = nil
	c.outwriter = make(map[string]chan (map[string]interface{}), 10)
	// c.init()
	return c
}

func NewConversion(reader *io.ReadCloser, outdir string, cfgname string, cnvformat Encoding, filetype string) *Execution {
	// c := Execution{}
	c := new(Execution)
	c.Reader = reader
	c.Outdir = outdir
	c.cfgname = cfgname
	c.Cnvformat = cnvformat
	c.Filetype = filetype
	c.lasterr = nil
	c.outwriter = make(map[string]chan (map[string]interface{}), 10)
	c.init()
	return c
}

func (c *Execution) SetCfg(cfg string) {
	c.cfgname = cfg
	if len(c.cfgname) > 0 {
		c.init()
	}
}

func (c *Execution) SetInfile(infile string) {
	c.Infile = infile
}

func (c *Execution) SetOutdir(pat string) {
	c.Outdir = pat
}

func (c *Execution) defaults() {
	if len(c.Infile) == 0 {
		c.Infile = "-"
	}
	if len(c.Outdir) == 0 {
		c.Outdir = path.Dir(c.Infile)
	}

}

func (c *Execution) Execute() {
	c.init()
	fmt.Fprintf(os.Stderr, "start")
	c.cfg.parse()
	c.end()

}

func (c *Execution) Print() {
	// fmt.Printf( ".Reader %s" ,c.Reader)
	fmt.Printf(".Outdir %s", c.Outdir)
	fmt.Printf(".Cfgname %s", c.cfgname)
	fmt.Printf(".Cnvformat %s", c.Cnvformat.AsString())
	if c.lasterr != nil {
		fmt.Printf(".lasterr %s", c.lasterr.Error())
	}
	fmt.Printf("Cfg %#v", c.cfg)
}
