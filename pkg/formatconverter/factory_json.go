package formatconverter

import (
	"encoding/json"
	"io"
)

var _ EncoderFactory = jsonEncoderDecoderFactory{}
var _ DecoderFactory = jsonEncoderDecoderFactory{}

// jsonEncoderDecoderFactory implements DecoderFactory, EncoderFactory
// to create decoders and encoders for JSON.
type jsonEncoderDecoderFactory struct{}

// NewDecoder implements DecoderFactory.
func (j jsonEncoderDecoderFactory) NewDecoder(reader io.Reader) Decoder {
	return json.NewDecoder(reader)
}

// NewEncoder implements EncoderFactory.
func (j jsonEncoderDecoderFactory) NewEncoder(writer io.Writer) Encoder {
	return json.NewEncoder(writer)
}

// FormatName implements DecoderFactory, EncoderFactory.
func (j jsonEncoderDecoderFactory) FormatName() string { return "json" }
