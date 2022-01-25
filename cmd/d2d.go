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
	"path"
	"strings"

	"github.com/gidor/cnv/pkg/cnv"
	"github.com/spf13/cobra"
)

// d2dCmd represents the d2d command
var d2dCmd = &cobra.Command{
	Use:   "d2d",
	Short: "Delimited 2 delimited",
	Long: `convert from delimited multi records file described in configuration 
	to a set of delimited files.`,
	Run: func(cmd *cobra.Command, args []string) {
		d2d()
	},
}

func init() {
	d2dCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	d2dCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Destination Dir")
	d2dCmd.Flags().StringVarP(&desctype, "type", "t", "", "Type when using standard input")
	d2dCmd.Flags().StringVarP(&config, "config", "c", "", "Config File")
	d2dCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "Delimiter char")

	rootCmd.AddCommand(d2dCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// d2dCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// d2dCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func d2d() {

	var reader io.ReadCloser = os.Stdin

	defer func() {
		reader.Close()
	}()

	openin(&reader)
	par := cnv.GetConversion(&reader)
	if desctype == "" {
		if inputFile != "" {
			par.Filetype = strings.
				Trim(path.Ext(inputFile), ".")
		} else {
			par.Filetype = "--"
		}
	} else {
		par.Filetype = desctype
	}

	if len(delimiter) == 0 {
		par.Delimiter = '|'
	} else {
		par.Delimiter = []rune(delimiter)[0]
	}

	par.Infile = inputFile
	par.Outdir = outputDir

	if len(config) > 0 {
		par.SetCfg(config)
	} else {
		par.SetCfg("cnv.yaml")

	}

	par.Execute()
}
