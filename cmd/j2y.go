/*
Copyright © 2021 Gianni  Doria gianni.doria@gmail.com
*/

package cmd

import (
	"fmt"

	"io"
	"os"

	"encoding/json"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"
)

// j2yCmd represents the j2y command
var j2yCmd = &cobra.Command{
	Use:   "j2y",
	Short: "Translate Jsont to Yaml",
	Long:  `Translate Jsont to Yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("j2y called")
		json2yaml()
	},
}

func init() {
	j2yCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Source File")
	j2yCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Destination File")

	rootCmd.AddCommand(j2yCmd)

}

func json2yaml() {

	var _internal = &internal{}

	var inputFile string
	var outputFile string

	var reader io.ReadCloser = os.Stdin
	var writer io.WriteCloser = os.Stdout

	defer func() {
		writer.Close()
		reader.Close()
	}()
	cont := 1
	fmt.Printf("%d\n", cont)
	cont++
	if inputFile != "" {
		if in, e := os.OpenFile(inputFile, os.O_RDONLY, 0755); e != nil {
			panic(e)
		} else {
			reader = in
		}
	}
	fmt.Printf("%d\n", cont)
	cont++

	if outputFile != "" {
		if out, e := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644); e == nil {
			writer = out
		}
	}
	fmt.Printf("%d\n", cont)
	cont++

	if err := json.NewDecoder(reader).Decode(_internal); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", cont)
	cont++

	if err := yaml.NewEncoder(writer).Encode(_internal); err != nil {
		panic(err)
	}
	fmt.Printf("Fine %d\n", cont)

}
