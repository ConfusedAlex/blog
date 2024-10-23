package cmd

import (
	"codeberg.org/confusedalex/website/internal/routes"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	serve = &cobra.Command{
		Use:   "serve",
		Short: "serve the blog",
		Run: func(cmd *cobra.Command, args []string) {
			router := routes.NewRouter()

			log.Fatal(http.ListenAndServe(":3000", router))
		},
	}
)

func init() {
	rootCmd.AddCommand(serve)
}
