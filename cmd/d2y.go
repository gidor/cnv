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

var ()

// d2yCmd represents the d2y command
var d2yCmd = &cobra.Command{
	Use:   "d2y",
	Short: "delimited file to yaml",
	Long: `convert from delimited file described in configuration 
to a set of yaml files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d2y called")
		d2y()
	},
}

func init() {

	d2yCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	d2yCmd.Flags().StringVarP(&outputDir, "outdir", "o", "", "Destination Dir")
	d2yCmd.Flags().StringVarP(&desctype, "type", "t", "", "Destination Dir")
	d2yCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print indent")
	d2yCmd.Flags().BoolVar(&htmlescape, "html-escape", false, "Html Escape strings")

	rootCmd.AddCommand(d2yCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// d2yCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// d2yCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func d2y() {

	var reader io.ReadCloser = os.Stdin

	defer func() {
		reader.Close()
	}()

	if desctype == "" {
		if inputFile != "" {
			desctype = path.Ext(inputFile)
		} else {
			desctype = "--"
		}
	}

	openin(&reader)

	var par cnv.Conversion

	par.Reader = &reader
	par.Outdir = outputDir
	par.Cnvformat = cnv.Yaml
	par.Filetype = desctype
	par.Pretty = pretty
	par.Htm_escape = htmlescape
	par.Cfgname = "cnv.yaml"

	delimted.Delimited(&par)

}
