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

	if inputFile != "" {
		if in, e := os.OpenFile(inputFile, os.O_RDONLY, 0755); e != nil {
			panic(e)
		} else {
			reader = in
		}
	}

	if outputFile != "" {
		if out, e := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644); e == nil {
			writer = out
		}
	}
	data, err := json.Load(reader)
	if err != nil {
		panic(err)
	}
	err = yaml.Save(&data, writer, true, false)
	if err != nil {
		panic(err)
	}

}
