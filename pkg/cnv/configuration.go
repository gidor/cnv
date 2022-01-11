/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cnv

type Configuration struct {
	Params    map[string]string
	Files     []*Input `yaml:",flow"`
	execution *Execution
}

func (c *Configuration) init(cnv *Execution) {
	c.execution = cnv
	// for _, i := range c.Files {
	// 	i.init(cnv)
	// }
}

func (c *Configuration) parse() {

	for _, i := range c.Files {
		if i.Filetype == c.execution.Filetype {
			i.init(c.execution)
			i.parse()
		}
	}
}
