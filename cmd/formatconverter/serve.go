package main

import (
	"github.com/dsxack/formatconverter/pkg/formatconverter"
	"github.com/spf13/cobra"
	"net/http"
)

var serveAddr *string

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("http server listening on %s\n", *serveAddr)
		server := http.Server{
			Addr:    *serveAddr,
			Handler: http.HandlerFunc(formatconverter.Serve),
		}
		return server.ListenAndServe()
	},
}

func init() {
	serveAddr = serveCmd.Flags().StringP("addr", "a", ":8080", "address to listen by http server")
}
