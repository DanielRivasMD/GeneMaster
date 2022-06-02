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
	"github.com/ttacon/chalk"

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
	Short: "A gene to rule them all.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

	` + chalk.Green.Color("gene") + chalk.Magenta.Color(` is a robot for automation of
	gene operations.

	`) + chalk.Green.Color("gene") + ` creates a convenient command line interface
	to handle genomic files.
	`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	if ε := rootCmd.Execute(); ε != nil {
		fmt.Println(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobra.OnInitialize(initConfig)

	// persistent flags
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
