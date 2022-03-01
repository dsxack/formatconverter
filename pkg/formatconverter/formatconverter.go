package formatconverter

import (
	"fmt"
	"io"
)

type FormatConverter struct {
	encoderFactory EncoderFactory
	decoderFactory DecoderFactory
}

func New(encoderFactory EncoderFactory, decoderFactory DecoderFactory) *FormatConverter {
	return &FormatConverter{
		encoderFactory: encoderFactory,
		decoderFactory: decoderFactory,
	}
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
