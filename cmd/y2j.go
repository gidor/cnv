/*
Copyright © 2021 Gianni  Doria gianni.doria@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// y2jCmd represents the y2j command
var y2jCmd = &cobra.Command{
	Use:   "y2j",
	Short: "yaml to json translation",
	Long: `yaml to json 
	translation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("y2j called")
		yaml2json()
	},
}

func init() {
	y2jCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	y2jCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")

	rootCmd.AddCommand(y2jCmd)

}

func yaml2json() {

	var _internal = &internal{}

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

	if err := yaml.NewDecoder(reader).Decode(_internal); err != nil {
		panic(err)
	}

	if err := json.NewEncoder(writer).Encode(_internal); err != nil {
		panic(err)
	}
}
