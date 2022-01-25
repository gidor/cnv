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

// y2jCmd represents the y2j command
var y2jCmd = &cobra.Command{
	Use:   "y2j",
	Short: "Convert Yaml to Json ",
	Long:  `Convert Yaml to Json `,
	Run: func(cmd *cobra.Command, args []string) {
		yaml2json()
	},
}

func init() {
	y2jCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	y2jCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")
	y2jCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print indent")
	y2jCmd.Flags().BoolVar(&htmlescape, "html-escape", false, "Html Escape strings")

	rootCmd.AddCommand(y2jCmd)

}

func yaml2json() {

	var reader io.ReadCloser = os.Stdin
	var writer io.WriteCloser = os.Stdout

	defer func() {
		writer.Close()
		reader.Close()
	}()

	openioout(&reader, &writer)

	data, err := yaml.Load(reader)
	if err != nil {
		panic(err)
	}
	err = json.Save(&data, writer, pretty, htmlescape)
	if err != nil {
		panic(err)
	}
}
