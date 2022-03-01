package main

import (
	"fmt"
	"github.com/dsxack/formatconverter/pkg/formatconverter"
	"github.com/spf13/cobra"
	"os"
)

var convertCmd = &cobra.Command{
	Use:   "convert <source> <destination>",
	Short: "convert your source file into destination file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			srcPath = args[0]
			dstPath = args[1]
		)

		src, err := os.Open(srcPath)
		if err != nil {
			return fmt.Errorf("opening source file: %s", srcPath)
		}
		dst, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("create destination file: %s", dstPath)
		}

		decoderFactory, err := formatconverter.NewDecoderFactoryByFilename(srcPath)
		if err != nil {
			return fmt.Errorf("create decoder factory: %v", err)
		}
		encoderFactory, err := formatconverter.NewEncoderFactoryByFilename(dstPath)
		if err != nil {
			return fmt.Errorf("create encoder factory: %v", err)
		}

		converter, err := formatconverter.New(encoderFactory, decoderFactory)
		if err != nil {
			return fmt.Errorf("create converter: %v", err)
		}

		err = converter.Convert(dst, src)
		if err != nil {
			return fmt.Errorf("convert %s into %s: %v", srcPath, dstPath, err)
		}

		return nil
	},
}
