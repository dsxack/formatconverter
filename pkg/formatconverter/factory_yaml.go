package formatconverter

import (
	"gopkg.in/yaml.v3"
	"io"
)

var _ EncoderFactory = yamlEncoderDecoderFactory{}
var _ DecoderFactory = yamlEncoderDecoderFactory{}

// yamlEncoderDecoderFactory implements DecoderFactory, EncoderFactory
// to create decoders and encoders for YAML.
type yamlEncoderDecoderFactory struct{}

// NewDecoder implements DecoderFactory.
func (j yamlEncoderDecoderFactory) NewDecoder(reader io.Reader) Decoder {
	return yaml.NewDecoder(reader)
}

// NewEncoder implements EncoderFactory.
func (j yamlEncoderDecoderFactory) NewEncoder(writer io.Writer) Encoder {
	return yaml.NewEncoder(writer)
}

// FormatName implements DecoderFactory, EncoderFactory.
func (j yamlEncoderDecoderFactory) FormatName() string { return "yaml" }
