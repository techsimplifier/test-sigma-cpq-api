package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// quotesCmd represents the quotes command
var quotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Use the /api/quotes endpoint",
	Long:  `Exposes the /api/quotes endpoint`,
}

func init() {
	rootCmd.AddCommand(quotesCmd)

	quotesCmd.PersistentFlags().StringP("endpoint", "", "https://localhost", "Endpoint for the Sigma CPQ swagger API")
	quotesCmd.PersistentFlags().StringP("templates", "", "./templates", "Path to the templates used for rendering api results")

	// Bind flags from the command line to the viper framework
	if err := viper.BindPFlags(quotesCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

}
