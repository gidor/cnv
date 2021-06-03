/*
Copyright © 2021 Gianni  Doria gianni.doria@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// csv2jCmd represents the csv2j command
var csv2jCmd = &cobra.Command{
	Use:   "csv2j",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("csv2j called")
	},
}

func init() {
	rootCmd.AddCommand(csv2jCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csv2jCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csv2jCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
