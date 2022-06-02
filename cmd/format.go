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
	"regexp"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
const ρε = `[\t,|*_]` // backticks are used here to contain the expression

var (
	ρ      = regexp.MustCompile(ρε) // declare regex
	header = []string{"chr", "mutStart", "mutEnd", "wild", "alternative", "CAtype", "mutation", "patientID", "motifStart", "motifEnd", "motifID", "precisionMxID", "scoreFIMO", "pValue", "strandSense"}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: `Format ` + chalk.Yellow.Color("bed") + ` files.`,
	Long:  `Format ` + chalk.Yellow.Color("bed") + ` files.`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// execute logic
		bedReadReg(bedFile, header)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	bedCmd.AddCommand(formatCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////
