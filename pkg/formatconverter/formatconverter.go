package formatconverter

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
)

type FormatConverter struct{}

func New() *FormatConverter {
	return &FormatConverter{}
}

func (converter FormatConverter) Convert(dst io.Writer, src io.Reader) error {
	decoder := json.NewDecoder(src)

	var value interface{}

	err := decoder.Decode(&value)
	if err != nil {
		return fmt.Errorf("error decoding json: %v", err)
	}
	encoder := yaml.NewEncoder(dst)
	err = encoder.Encode(value)
	if err != nil {
		return fmt.Errorf("errror encoding yaml: %v", err)
	}

	return nil
}
