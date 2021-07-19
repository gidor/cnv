/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/gidor/cnv/cnv/csv"
	"github.com/gidor/cnv/cnv/yaml"
	"github.com/spf13/cobra"
)

// csv2yCmd represents the csv2y command
var csv2yCmd = &cobra.Command{
	Use:   "csv2y",
	Short: "Convert csv data to yaml",
	Long: `Convert csv data to yaml
	
	
	.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("csv2y called")
		csv2yaml()
	},
}

func init() {
	csv2yCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	csv2yCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")
	csv2yCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print indent")

	rootCmd.AddCommand(csv2yCmd)

}

func csv2yaml() {

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

	if err := yaml.Save(&data, writer, pretty, false); err != nil {
		panic(err)
	}

}
