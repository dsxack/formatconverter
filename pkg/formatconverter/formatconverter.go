package formatconverter

import (
	"fmt"
	"io"
)

// FormatConverter converts io.Reader content into another format and saves it to io.Writer.
type FormatConverter struct {
	encoderFactory EncoderFactory
	decoderFactory DecoderFactory
}

// NewFormatConverter creates new FormatConverter.
// encoderFactory is EncoderFactory that needs to create Encoder for destination io.Writer.
// decoderFactory is DecoderFactory that needs to create Decoder for source io.Reader.
// Returns error if encoderFactory and decoderFactory have same formats.
func NewFormatConverter(encoderFactory EncoderFactory, decoderFactory DecoderFactory) (*FormatConverter, error) {
	if encoderFactory.FormatName() == decoderFactory.FormatName() {
		return nil, fmt.Errorf("source and destination have same formats: %v", decoderFactory.FormatName())
	}

	return &FormatConverter{
		encoderFactory: encoderFactory,
		decoderFactory: decoderFactory,
	}, nil
}

func (converter FormatConverter) Convert(dst io.Writer, src io.Reader) error {
	decoder := converter.decoderFactory.NewDecoder(src)

	var value interface{}

	err := decoder.Decode(&value)
	if err != nil {
		return fmt.Errorf("decode: %v", err)
	}
	encoder := converter.encoderFactory.NewEncoder(dst)
	err = encoder.Encode(value)
	if err != nil {
		return fmt.Errorf("encode: %v", err)
	}

	return nil
}
