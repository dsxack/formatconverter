package formatconverter

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type DirWalker struct {
	srcPath        string
	dstPath        string
	encoderFactory EncoderFactory
	logger         Logger
	dstFormat      string
}

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
		walker.printError(relativePath, err)
		return nil
	}

	converter, err := NewConverter(walker.encoderFactory, decoderFactory)
	if err != nil {
		walker.printError(relativePath, err)
		return nil
	}

	src, err := os.Open(path)
	if err != nil {
		walker.printError(relativePath, err)
		return nil
	}

	dst, err := os.Create(filepath.Join(walker.dstPath, walker.dstFilename(path)))
	if err != nil {
		walker.printError(relativePath, err)
		return nil
	}

	err = converter.Convert(dst, src)
	if err != nil {
		walker.printError(relativePath, err)
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
