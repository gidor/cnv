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
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cnv",
	Short: "format translation",
	Long: `
     _________ _   __  
    / ___/ __ \ | / /  
   / /__/ / / / |/ /   
   \___/_/ /_/|___/    

	Conversions between data formats.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// viper.SetConfigName("cnv")  // name of config file (without extension)
	// viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath(".")    // optionally look for config in the working directory
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $cnv.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	// home, err := homedir.Dir()
	// 	// cobra.CheckErr(err)
	// 	// Search config in home directory with name ".cnv" (without extension).
	// 	// viper.AddConfigPath(home)
	// 	// viper.SetConfigName(".cnv.yaml")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// } else {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		fmt.Fprintln(os.Stderr, "No config file found")
	// 		// Config file not found; ignore error if desired
	// 	} else {
	// 		fmt.Fprintln(os.Stderr, "Error:", ok)
	// 		// Config file was found but another error was produced
	// 	}
	// }

}
