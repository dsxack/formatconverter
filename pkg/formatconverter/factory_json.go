package formatconverter

import (
	"encoding/json"
	"io"
)

var _ EncoderFactory = jsonEncoderDecoderFactory{}
var _ DecoderFactory = jsonEncoderDecoderFactory{}

type jsonEncoderDecoderFactory struct{}

func (j jsonEncoderDecoderFactory) NewDecoder(reader io.Reader) Decoder {
	return json.NewDecoder(reader)
}

func (j jsonEncoderDecoderFactory) NewEncoder(writer io.Writer) Encoder {
	return json.NewEncoder(writer)
}

func (j jsonEncoderDecoderFactory) FormatName() string { return "json" }
