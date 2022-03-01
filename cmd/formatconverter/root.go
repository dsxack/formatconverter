package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "formatconverter",
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(convertCmd)
}
