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

package cmd

import (
	"io"
	"os"

	"github.com/gidor/cnv/pkg/csv"
	"github.com/gidor/cnv/pkg/json"

	"github.com/spf13/cobra"
)

// csv2jCmd represents the csv2j command
var csv2jCmd = &cobra.Command{
	Use:   "csv2j",
	Short: "Convert csv data to json",
	Long:  `Convert csv data to json to be done`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("csv2j called")
		csv2json()
	},
}

func init() {

	csv2jCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	csv2jCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")
	csv2jCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print indent")
	csv2jCmd.Flags().BoolVar(&htmlescape, "html-escape", false, "Html Escape strings")

	rootCmd.AddCommand(csv2jCmd)
}

func csv2json() {

	// var _internal = &internal{}
	// var _internal = make(internal)

	var reader io.ReadCloser = os.Stdin
	var writer io.WriteCloser = os.Stdout

	defer func() {
		writer.Close()
		reader.Close()
	}()

	openioout(&reader, &writer)

	data, err := csv.Load(reader)
	if err != nil {

		panic(err)
	}

	if err := json.Save(&data, writer, pretty, htmlescape); err != nil {
		panic(err)
	}
}
