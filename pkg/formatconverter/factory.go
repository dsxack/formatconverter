package formatconverter

import (
	"fmt"
	"io"
	"path/filepath"
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
	".json": jsonEncoderDecoderFactory{},
	".yaml": yamlEncoderDecoderFactory{},
	".yml":  yamlEncoderDecoderFactory{},
}

func NewEncoderFactoryByFilename(filename string) (EncoderFactory, error) {
	ext := filepath.Ext(filename)
	encoder, ok := encodersMap[ext]
	if !ok {
		return nil, fmt.Errorf("can not encode files with ext: %v", ext)
	}
	return encoder, nil
}

func NewDecoderFactoryByFilename(filename string) (DecoderFactory, error) {
	ext := filepath.Ext(filename)
	decoder, ok := encodersMap[ext]
	if !ok {
		return nil, fmt.Errorf("can not decode files with ext: %v", ext)
	}
	return decoder, nil
}
