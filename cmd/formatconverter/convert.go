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
			return fmt.Errorf("error opening source file: %s", srcPath)
		}
		dst, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("error create destination file: %s", dstPath)
		}

		decoderFactory, err := formatconverter.NewDecoderFactoryByFilename(srcPath)
		if err != nil {
			return err
		}
		encoderFactory, err := formatconverter.NewEncoderFactoryByFilename(dstPath)
		if err != nil {
			return err
		}

		converter := formatconverter.New(encoderFactory, decoderFactory)
		err = converter.Convert(dst, src)
		if err != nil {
			return fmt.Errorf("error convert %s into %s: %v", srcPath, dstPath, err)
		}

		return nil
	},
}
