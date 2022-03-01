package main

import (
	"fmt"
	"github.com/dsxack/formatconverter/pkg/formatconverter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var dstFormat *string

var convertCmd = &cobra.Command{
	Use:   "convert <source> <destination>",
	Short: "convert your source file into destination file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			srcPath = args[0]
			dstPath = args[1]
		)

		srcStat, err := os.Stat(srcPath)
		if err != nil {
			return fmt.Errorf("src stat: %v", err)
		}
		dstStat, dstErr := os.Stat(dstPath)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("dst stat: %v", err)
		}

		if dstErr == nil && srcStat.IsDir() != dstStat.IsDir() {
			if srcStat.IsDir() {
				return fmt.Errorf("source is directory, but destination is file")
			}
			return fmt.Errorf("source is file, but destination is directory")
		}

		if srcStat.IsDir() {
			if os.IsNotExist(dstErr) {
				err = os.MkdirAll(dstPath, os.ModePerm)
				if err != nil {
					return fmt.Errorf("make destination dir: %v", err)
				}
			}

			if dstFormat == nil || *dstFormat == "" {
				return fmt.Errorf("dst-format flag is required when source is directory")
			}
			dirWalker, err := formatconverter.NewDirWalker(
				srcPath,
				dstPath,
				*dstFormat,
				log.New(cmd.OutOrStderr(), "", log.LstdFlags),
			)
			if err != nil {
				return fmt.Errorf("create dir walker: %v", err)
			}

			err = dirWalker.Walk()
			if err != nil {
				return fmt.Errorf("walk dir: %v", err)
			}
			return nil
		}

		src, err := os.Open(srcPath)
		if err != nil {
			return fmt.Errorf("opening source file: %s", srcPath)
		}
		dst, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("create destination file: %s: %v", dstPath, err)
		}

		decoderFactory, err := formatconverter.NewDecoderFactoryByFilename(srcPath)
		if err != nil {
			return fmt.Errorf("create decoder factory: %v", err)
		}

		var encoderFactory formatconverter.EncoderFactory
		if dstFormat != nil && *dstFormat != "" {
			encoderFactory, err = formatconverter.NewEncoderFactoryByFormat(*dstFormat)
			if err != nil {
				return fmt.Errorf("create encoder factory: %v", err)
			}
		} else {
			encoderFactory, err = formatconverter.NewEncoderFactoryByFilename(dstPath)
			if err != nil {
				return fmt.Errorf("create encoder factory: %v", err)
			}
		}

		converter, err := formatconverter.NewConverter(encoderFactory, decoderFactory)
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

func init() {
	dstFormat = convertCmd.Flags().StringP("dst-format", "d", "", "set destination format")
}
