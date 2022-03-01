package formatconverter

import (
	"gopkg.in/yaml.v3"
	"io"
)

var _ EncoderFactory = yamlEncoderDecoderFactory{}
var _ DecoderFactory = yamlEncoderDecoderFactory{}

type yamlEncoderDecoderFactory struct{}

func (j yamlEncoderDecoderFactory) NewDecoder(reader io.Reader) Decoder {
	return yaml.NewDecoder(reader)
}

func (j yamlEncoderDecoderFactory) NewEncoder(writer io.Writer) Encoder {
	return yaml.NewEncoder(writer)
}
