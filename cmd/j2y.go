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

	"github.com/gidor/cnv/pkg/json"
	"github.com/gidor/cnv/pkg/yaml"
	"github.com/spf13/cobra"
)

// j2yCmd represents the j2y command
var j2yCmd = &cobra.Command{
	Use:   "j2y",
	Short: "Convert Json to Yaml",
	Long:  `Convert Json to Yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("j2y called")
		json2yaml()
	},
}

func init() {
	j2yCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	j2yCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")
	j2yCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print indent")

	rootCmd.AddCommand(j2yCmd)

}

func json2yaml() {

	var reader io.ReadCloser = os.Stdin
	var writer io.WriteCloser = os.Stdout

	defer func() {
		writer.Close()
		reader.Close()
	}()

	openioout(&reader, &writer)

	data, err := json.Load(reader)
	if err != nil {
		panic(err)
	}
	err = yaml.Save(&data, writer, true, false)
	if err != nil {
		panic(err)
	}

}
