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

	"encoding/json"

	goyaml "gopkg.in/yaml.v3"
)

type WriteHandler func(ch chan (map[string]interface{}), cnv Conversion, prefix string, suffix string)

func yamlWriteHandler(ch chan (map[string]interface{}), cnv *Conversion, prefix string, suffix string) {
	outputFile := path.Join(cnv.Outdir, prefix+path.Base(cnv.Infile)+suffix+".txt")
	var writer *os.File
	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}
	defer writer.Close()
	encoder := goyaml.NewEncoder(writer)

	for {
		m, ok := <-ch
		if ok {
			err := encoder.Encode(m)
			if err != nil {
				cnv.SetError(err)
			}
		} else {
			return
		}
	}
}

func jsonWriteHandler(ch chan (map[string]interface{}), cnv *Conversion, prefix string, suffix string) {
	outputFile := path.Join(cnv.Outdir, prefix+path.Base(cnv.Infile)+suffix+".txt")
	var writer *os.File
	if outputFile != "" {
		if out, err := os.Create(outputFile); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}
	defer writer.Close()
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	for {
		m, ok := <-ch
		if ok {
			err := encoder.Encode(m)
			if err != nil {
				cnv.SetError(err)
			}
		} else {
			return
		}
	}
}

type Conversion struct {
	Reader    *io.ReadCloser                           // input file reader
	Infile    string                                   // input file name
	Outdir    string                                   // output dir i  wich decoded file mast be written
	Cfgname   string                                   // Confguration file name
	Filetype  string                                   // file type in input decription use extension by default
	Cnvformat Format                                   // Output File Format Csv, Yaml, Json
	lasterr   error                                    // last error uccurred
	cfg       Configuration                            // configuration
	outwriter map[string]chan (map[string]interface{}) // named chanel to the output writers

}

// close all channels so handler will terminate and close created files
func (c *Conversion) Close() {
	for k, v := range c.outwriter {
		close(v)
		fmt.Fprintf(os.Stderr, "closing %s", k)
	}
}

func (c *Conversion) Write(data map[string]interface{}, on Output) {
	ch, ok := c.outwriter[on.Name]

	if !ok {
		ch = make(chan (map[string]interface{}))
		c.outwriter[on.Name] = ch
		if c.Cnvformat == Yaml {
			go yamlWriteHandler(ch, c, on.Prefix, on.Suffix)
		} else {
			go jsonWriteHandler(ch, c, on.Prefix, on.Suffix)
		}
	}
	ch <- data
}

func (c *Conversion) Error() string {
	return c.lasterr.Error()
}

//error setter
func (c *Conversion) SetError(err error) {
	c.lasterr = err
}

func (c *Conversion) init() {
	b, err := ioutil.ReadFile(c.Cfgname)
	if err != nil {
		panic(err)
	}
	c.cfg = Configuration{}

	err = goyaml.Unmarshal(b, &c.cfg)
	if err != nil {
		c.lasterr = err
		panic(err)
	}
	c.cfg.init(c)
}

func NewConversion(reader *io.ReadCloser, outdir string, cfgname string, cnvformat Format, filetype string) *Conversion {
	// c := Conversion{}
	c := new(Conversion)
	c.Reader = reader
	c.Outdir = outdir
	c.Cfgname = cfgname
	c.Cnvformat = cnvformat
	c.Filetype = filetype
	c.lasterr = nil
	c.init()
	c.outwriter = make(map[string]chan (map[string]interface{}), 10)
	return c
}

func (c *Conversion) Init() {
	c.init()
}

func (c *Conversion) Print() {
	// fmt.Printf( ".Reader %s" ,c.Reader)
	fmt.Printf(".Outdir %s", c.Outdir)
	fmt.Printf(".Cfgname %s", c.Cfgname)
	fmt.Printf(".Cnvformat %s", c.Cnvformat.AsString())
	fmt.Printf(".lasterr %s", c.lasterr.Error())
	fmt.Printf("Cfg %#v", c.cfg)
}
