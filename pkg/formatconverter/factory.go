package formatconverter

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

// Encoder is interface for encoding any value into specific format.
type Encoder interface {
	Encode(interface{}) error
}

// Decoder is interface for decoding any value from specific format.
type Decoder interface {
	Decode(interface{}) error
}

// EncoderFactory is factories interface for creating specific format encoders.
type EncoderFactory interface {
	// NewEncoder accepts io.Writer to encode and returns Encoder for that io.Reader.
	NewEncoder(io.Writer) Encoder
	// FormatName return name of format of decoders that could be created by factory.
	// It needs to check that encoder and decoder have different formats.
	FormatName() string
}

// DecoderFactory is factories interface for creating specific format decoders.
type DecoderFactory interface {
	// NewDecoder accepts io.Reader to decode and returns Decoder for that io.Reader.
	NewDecoder(io.Reader) Decoder
	// FormatName return name of format of decoders that could be created by factory.
	// It needs to check that encoder and decoder have different formats.
	FormatName() string
}

// EncoderDecoderFactory is aggregates two interfaces EncoderFactory and DecoderFactory.
// Needs to store all encoders/decoders in single encodersByFormat map (and also encodersByMimeType map).
type EncoderDecoderFactory interface {
	EncoderFactory
	DecoderFactory
}

var encodersByFormat = map[string]EncoderDecoderFactory{
	"json": jsonEncoderDecoderFactory{},
	"yaml": yamlEncoderDecoderFactory{},
	"yml":  yamlEncoderDecoderFactory{},
}

var encodersByMimeType = map[string]EncoderDecoderFactory{
	"application/json": jsonEncoderDecoderFactory{},
	"application/yaml": yamlEncoderDecoderFactory{},
}

// NewEncoderFactoryByFilename returns specific EncoderFactory for filename.
// It uses file extension to detect file format.
// Returns error if file format is unsupported.
func NewEncoderFactoryByFilename(filename string) (EncoderFactory, error) {
	ext := filepath.Ext(filename)
	format := strings.TrimLeft(ext, ".")
	return NewEncoderFactoryByFormat(format)
}

// NewEncoderFactoryByFormat returns specific EncoderFactory for format.
// Returns error if format is unsupported.
func NewEncoderFactoryByFormat(format string) (EncoderFactory, error) {
	encoder, ok := encodersByFormat[format]
	if !ok {
		return nil, fmt.Errorf("can not encode files with format: %v", format)
	}
	return encoder, nil
}

// NewDecoderFactoryByFilename returns specific DecoderFactory for filename.
// It uses file extension to detect file for
// Returns error if file format is unsupported.
func NewDecoderFactoryByFilename(filename string) (DecoderFactory, error) {
	ext := filepath.Ext(filename)
	format := strings.TrimLeft(ext, ".")
	return NewDecoderFactoryByFormat(format)
}

// NewDecoderFactoryByFormat returns specific DecoderFactory for format.
// Returns error if format is unsupported.
func NewDecoderFactoryByFormat(format string) (DecoderFactory, error) {
	decoder, ok := encodersByFormat[format]
	if !ok {
		return nil, fmt.Errorf("can not decode files with format: %v", format)
	}
	return decoder, nil
}

// NewDecoderFactoryByMimeType returns specific DecoderFactory for mimetype.
// Returns error if mimetype is unsupported.
func NewDecoderFactoryByMimeType(mimetype string) (DecoderFactory, error) {
	encoder, ok := encodersByMimeType[mimetype]
	if !ok {
		return nil, fmt.Errorf("can not decode files with mimetype: %v", mimetype)
	}
	return encoder, nil
}
