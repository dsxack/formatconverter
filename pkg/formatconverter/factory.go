package formatconverter

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Encoder interface {
	Encode(interface{}) error
}

type Decoder interface {
	Decode(interface{}) error
}

type EncoderFactory interface {
	NewEncoder(io.Writer) Encoder
	FormatName() string
}

type DecoderFactory interface {
	NewDecoder(io.Reader) Decoder
	FormatName() string
}

type EncoderDecoderFactory interface {
	EncoderFactory
	DecoderFactory
}

var encodersMap = map[string]EncoderDecoderFactory{
	"json": jsonEncoderDecoderFactory{},
	"yaml": yamlEncoderDecoderFactory{},
	"yml":  yamlEncoderDecoderFactory{},
}

func NewEncoderFactoryByFilename(filename string) (EncoderFactory, error) {
	ext := filepath.Ext(filename)
	format := strings.TrimLeft(ext, ".")
	return NewEncoderFactoryByFormat(format)
}

func NewEncoderFactoryByFormat(format string) (EncoderFactory, error) {
	encoder, ok := encodersMap[format]
	if !ok {
		return nil, fmt.Errorf("can not encode files with format: %v", format)
	}
	return encoder, nil
}

func NewDecoderFactoryByFilename(filename string) (DecoderFactory, error) {
	ext := filepath.Ext(filename)
	format := strings.TrimLeft(ext, ".")
	return NewDecoderFactoryByFormat(format)
}

func NewDecoderFactoryByFormat(format string) (DecoderFactory, error) {
	decoder, ok := encodersMap[format]
	if !ok {
		return nil, fmt.Errorf("can not decode files with format: %v", format)
	}
	return decoder, nil
}
