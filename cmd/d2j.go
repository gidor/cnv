/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)

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
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gidor/cnv/cnv"
	"github.com/gidor/cnv/cnv/delimted"
	"github.com/spf13/cobra"
)

// d2jCmd represents the d2j command
var d2jCmd = &cobra.Command{
	Use:   "d2j",
	Short: "delimited file to jsons ",
	Long: `convert from delimited file described in configuration 
to a set of json files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d2j called")
		d2j()
	},
}

func init() {

	d2jCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	d2jCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Destination Dir")
	d2jCmd.Flags().StringVarP(&desctype, "type", "t", "", "Destination Dir")

	rootCmd.AddCommand(d2jCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// d2jCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// d2jCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func d2j() {

	var reader io.ReadCloser = os.Stdin

	defer func() {
		reader.Close()
	}()

	openin(&reader)
	if desctype == "" {
		if inputFile != "" {
			desctype = path.Ext(inputFile)
		} else {
			desctype = "--"
		}
	}

	par := cnv.NewConversion(&reader, outputDir, "cnv.yaml", cnv.Yaml, desctype)

	delimted.Delimited(par)

}
