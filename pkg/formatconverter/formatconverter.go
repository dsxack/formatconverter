package formatconverter

import (
	"fmt"
	"io"
)

type FormatConverter struct {
	encoderFactory EncoderFactory
	decoderFactory DecoderFactory
}

func NewConverter(encoderFactory EncoderFactory, decoderFactory DecoderFactory) (*FormatConverter, error) {
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
