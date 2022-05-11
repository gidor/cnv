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

package delimted

import (
	"github.com/gidor/cnv/pkg/cnv"
)

func Delimited(execution *cnv.Execution) {

	// err := mapstructure.Decode(files, &cfg)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// }
	execution.Print()

	// fmt.Printf("%+v", files)
}