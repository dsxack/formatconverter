package formatconverter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// DirWalker walks through source directory and converts all files into destination directory
// with saving structure of source directory.
type DirWalker struct {
	srcPath        string
	dstPath        string
	encoderFactory EncoderFactory
	logger         Logger
	dstFormat      string
}

// NewDirWalker creates a new DirWalker.
// srcPath is path to source directory.
// dstPath is path to destination directory
// dstFormat is format for encoder.
// logger is any implementation of Logger for logging success and fail converting messages.
func NewDirWalker(srcPath, dstPath, dstFormat string, logger Logger) (*DirWalker, error) {
	encoderFactory, err := NewEncoderFactoryByFormat(dstFormat)
	if err != nil {
		return nil, err
	}

	return &DirWalker{
		srcPath:        srcPath,
		dstPath:        dstPath,
		dstFormat:      dstFormat,
		encoderFactory: encoderFactory,
		logger:         logger,
	}, nil
}

func (walker *DirWalker) Walk() error {
	return filepath.Walk(walker.srcPath, walker.walk)
}

func (walker *DirWalker) walk(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	relativePath, err := filepath.Rel(walker.srcPath, path)
	if err != nil {
		return err
	}

	decoderFactory, err := NewDecoderFactoryByFilename(path)
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("create decoder: %v", err))
		return nil
	}

	converter, err := NewFormatConverter(walker.encoderFactory, decoderFactory)
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("create converter: %v", err))
		return nil
	}

	src, err := os.Open(path)
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("open source file: %v", err))
		return nil
	}

	dirName, fileName := filepath.Split(relativePath)
	dstDirName := filepath.Join(walker.dstPath, dirName)

	err = os.MkdirAll(dstDirName, os.ModePerm)
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("make destination dir: %v", err))
		return nil
	}

	dst, err := os.Create(filepath.Join(dstDirName, walker.dstFilename(fileName)))
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("create destination file: %v", err))
		return nil
	}

	err = converter.Convert(dst, src)
	if err != nil {
		walker.printError(relativePath, fmt.Errorf("convert: %v", err))
		return nil
	}

	walker.printSuccess(relativePath)

	return nil
}

func (walker DirWalker) dstFilename(srcFilePath string) string {
	fileName := filepath.Base(srcFilePath)
	fileName = strings.TrimRight(fileName, filepath.Ext(fileName))
	return fileName + "." + walker.dstFormat
}

func (walker DirWalker) printError(relativePath string, err error) {
	walker.logger.Printf("%v: error: %v\n", relativePath, err)
}

func (walker DirWalker) printSuccess(relativePath string) {
	walker.logger.Printf("%v: success\n", relativePath)
}
