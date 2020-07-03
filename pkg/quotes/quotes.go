package quotes

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swagger "github.com/techsimplifier/test-sigma-cpq-api/go-client"
)

// RunGetQuote executes the quote read command
func RunGetQuote() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		endpoint := viper.GetString("endpoint")
		id := viper.GetString("id")
		tpl := viper.GetString("templates")

		// Load the output template
		t := template.Must(template.ParseFiles(fmt.Sprintf("%s/quote.tpl", tpl)))

		config := &swagger.Configuration{
			BasePath:      endpoint,
			DefaultHeader: make(map[string]string),
			UserAgent:     "Swagger-Codegen/1.0.0/go",
		}
		cpqClient := swagger.NewAPIClient(config)
		result, _, err := cpqClient.QuotesApi.ApiQuotesQuoteIdGet(context.Background(), id, nil)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		err = t.Execute(os.Stdout, result)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

	}
}
