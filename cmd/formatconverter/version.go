package main

import "github.com/spf13/cobra"

var Version = "development"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show information about version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("formatconverter version: %v\n", Version)
	},
}
