/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
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

type Execution struct {
	Reader      *io.ReadCloser                           // input file reader
	Infile      string                                   // input file name
	Outdir      string                                   // output dir i  wich decoded file mast be written
	cfgname     string                                   // Confguration file name
	Filetype    string                                   // file type in input decription use extension by default
	Cnvformat   Format                                   // Output File Format Csv, Yaml, Json
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
	}
	if c.Delimiter > 0 {
		params.delimiter = c.Delimiter
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

func NewConversion(reader *io.ReadCloser, outdir string, cfgname string, cnvformat Format, filetype string) *Execution {
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
	fmt.Fprintf(os.Stderr, "startig conversione")
	c.cfg.parse()

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
