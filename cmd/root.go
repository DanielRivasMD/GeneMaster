/*
Copyright © 2022 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	cfgFile string
	inDir   string
	outDir  string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GeneMaster",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if ε := rootCmd.Execute(); ε != nil {
		fmt.Println(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobra.OnInitialize(initConfig)

	// flags
	rootCmd.PersistentFlags().StringVarP(&inDir, "inDir", "I", ".", "Directory where input files are located")
	rootCmd.PersistentFlags().StringVarP(&outDir, "outDir", "O", ".", "Output directory. Creates if not exitst")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, ε := homedir.Dir()
		if ε != nil {
			fmt.Println(ε)
			os.Exit(1)
		}

		// Search config in home directory with name ".GeneMaster" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".GeneMaster")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if ε := viper.ReadInConfig(); ε == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// fileExist checks if a file exists and is not a directory before
// try using it to prevent further εors
func fileExist(filename string) bool {
	info, ε := os.Stat(filename)
	if os.IsNotExist(ε) {
		return false
	}
	return !info.IsDir()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
