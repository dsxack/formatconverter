package main

import (
	"github.com/dsxack/formatconverter/pkg/formatconverter"
	"github.com/spf13/cobra"
	"io"
)

var rootCmd = &cobra.Command{
	Use: "formatconverter",
}

func init() {
	cobra.OnInitialize(initialize)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(convertCmd)
}

var defaultConverter converter

type converter interface {
	Convert(dst io.Writer, src io.Reader) error
}

func initialize() {
	defaultConverter = formatconverter.New()
}
