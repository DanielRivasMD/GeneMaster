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

	"github.com/spf13/cobra"
)

// bedCmd represents the bed command
var bedCmd = &cobra.Command{
	Use:   "bed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bed called")
	},
}

func init() {
	rootCmd.AddCommand(bedCmd)

	// Here you will define your flags and configuration settings.

////////////////////////////////////////////////////////////////////////////////////////////////////

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bedCmd.PersistentFlags().String("foo", "", "A help for foo")
// read bed file & write
func bedRead(bedFile string) {

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// open an input file, exit on error
	inputFile, readErr := os.Open(bedFile)
	if readErr != nil {
		log.Fatal("Error opening input file : ", readErr)
	}

	// check whether file exists to avoid appending
	if fileExist(outDir + "/" + outFile) {
		os.Remove(outDir + "/" + outFile)
	}

	// scanner.Scan() advances to the next token returning false if an error was encountered
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		records := ρ.Split(scanner.Text(), -1) // second arg -1 means no limits for the number of substrings

		// write
		bedWrite(outDir+"/"+outFile, records)

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// write bed file
func bedWrite(outFile string, records []string) {

	// declare io
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// declare writer
	w := bufio.NewWriter(f)

	// writing
	for i := 0; i < len(records)-1; i++ {
		_, err = w.WriteString(records[i] + ",")
		if err != nil {
			panic(err)
		}
	}

	_, err = w.WriteString(records[len(records)-1] + "\n")
	if err != nil {
		panic(err)
	}

	// flush writer
	w.Flush()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
