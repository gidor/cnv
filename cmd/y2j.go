/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cmd

import (
	"io"
	"os"

	"github.com/gidor/cnv/cnv/json"
	"github.com/gidor/cnv/cnv/yaml"

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

	if inputFile != "" {
		if in, err := os.OpenFile(inputFile, os.O_RDONLY, 0755); err != nil {
			panic(err)
		} else {
			reader = in
		}
	}

	if outputFile != "" {
		if out, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			panic(err)
		} else {
			writer = out
		}
	}

	data, err := yaml.Load(reader)
	if err != nil {
		panic(err)
	}
	err = json.Save(&data, writer, pretty, htmlescape)
	if err != nil {
		panic(err)
	}
}
