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
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	bedFile string
	outFile string
	altFile string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// bedCmd represents the bed command
var bedCmd = &cobra.Command{
	Use:   "bed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(bedCmd)

	// flags
	bedCmd.PersistentFlags().StringVarP(&bedFile, "bed", "b", "", "Bed file")
	bedCmd.PersistentFlags().StringVarP(&outFile, "outfile", "o", "", "Out file")
	bedCmd.PersistentFlags().StringVarP(&outFile, "altfile", "a", "", "Alternative file")

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// read bed file & write
func bedRead(bedFile string) {

	// open an input file, exit on error
	inputFile, ε := os.Open(bedFile)
	if ε != nil {
		log.Fatal("Error opening input file : ", ε)
	}

	// check whether file exists to avoid appending
	if fileExist(outDir + "/" + outFile) {
		os.Remove(outDir + "/" + outFile)
	}

	// scanner.Scan() advances to the next token returning false if an error was encountered
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		// tab separated records
		records := strings.Split(scanner.Text(), "\t")

		// write
		bedWrite(outDir+"/"+outFile, records)

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// read bed file & write
func bedReadReg(bedFile string, header []string) {

	// open an input file, exit on error
	inputFile, ε := os.Open(bedFile)
	if ε != nil {
		log.Fatal("Error opening input file : ", ε)
	}

	// check whether file exists to avoid appending
	if fileExist(outDir + "/" + outFile) {
		os.Remove(outDir + "/" + outFile)
	}

	// write header
	bedWrite(outDir+"/"+outFile, header)

	// scanner.Scan() advances to the next token returning false if an error was encountered
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		// regex separated
		records := ρ.Split(scanner.Text(), -1) // second arg -1 means no limits for the number of substrings

		// write
		switch {
		case len(records) > len(header):
			bedWrite(outDir+"/"+altFile, records)
		case len(records) == len(header):
			bedWrite(outDir+"/"+outFile, records)
		case len(records) < len(header):
			fmt.Println("Records might be absent")
		}

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// write bed file
func bedWrite(outFile string, records []string) {

	// declare io
	ƒ, ε := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if ε != nil {
		panic(ε)
	}

	defer ƒ.Close()

	// declare writer
	ϖ := bufio.NewWriter(ƒ)

	// writing
	for ι := 0; ι < len(records)-1; ι++ { // iterate on records
		_, ε = ϖ.WriteString(records[ι] + ",")
		if ε != nil {
			panic(ε)
		}
	}

	// write last record
	_, ε = ϖ.WriteString(records[len(records)-1] + "\n")
	if ε != nil {
		panic(ε)
	}

	// flush writer
	ϖ.Flush()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
