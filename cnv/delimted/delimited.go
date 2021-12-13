/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package delimted

import (
	"github.com/gidor/cnv/cnv"
)

func Delimited(conversion *cnv.Conversion) {

	// err := mapstructure.Decode(files, &cfg)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// }
	conversion.Print()

	// fmt.Printf("%+v", files)
}
