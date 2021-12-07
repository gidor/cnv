/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package delimted

import (
	"fmt"
	"io/ioutil"

	"github.com/gidor/cnv/cnv"
	goyaml "gopkg.in/yaml.v3"
)

func Delimited(conversion *cnv.Conversion) {

	b, err := ioutil.ReadFile(conversion.Cfgname)
	if err != nil {
		panic(err)
	}
	var cfg cnv.CnvCfg

	err = goyaml.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}

	// err := mapstructure.Decode(files, &cfg)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// }

	fmt.Printf("%#v", cfg)
	fmt.Printf("%+v", cfg)

	// fmt.Printf("%+v", files)
}
