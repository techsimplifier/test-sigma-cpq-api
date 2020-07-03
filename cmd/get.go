package cmd

import (
	"log"

	"github.com/techsimplifier/test-sigma-cpq-api/pkg/quotes"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Calls GET /api/quotes/{id} in the CPQ",
	Long:  `Calls GET /api/quotes/{id} in the CPQ`,
	Run:   quotes.RunGetQuote(),
}

func init() {
	quotesCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "", "", "The quote id represented as a GUID like 4add8c86-5431-4d5c-b481-7479bf515db0")

	// Bind flags from the command line to the viper framework
	if err := viper.BindPFlags(getCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}
