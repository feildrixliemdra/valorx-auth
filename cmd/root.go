package cmd

import (
	"valorx-auth/cmd/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Start() {
	rootCmd := &cobra.Command{}

	cmd := []*cobra.Command{
		{
			Use:   "serve-http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start()
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
