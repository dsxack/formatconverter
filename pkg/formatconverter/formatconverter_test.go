package formatconverter

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFormatConverter_Convert(t *testing.T) {
	tests := []struct {
		name           string
		encoderFactory EncoderFactory
		decoderFactory DecoderFactory
		src            string
		wantDst        string
		wantErr        error
	}{
		{
			name:           "success decode from json to yaml",
			encoderFactory: yamlEncoderDecoderFactory{},
			decoderFactory: jsonEncoderDecoderFactory{},
			src:            `[{"testKey": "testValue"},{"testKey2": "testValue2"}]`,
			wantDst: `- testKey: testValue
- testKey2: testValue2
`,
		},
		{
			name:           "success decode from yaml to json",
			encoderFactory: jsonEncoderDecoderFactory{},
			decoderFactory: yamlEncoderDecoderFactory{},
			src: `- testKey: testValue
- testKey2: testValue2
`,
			wantDst: `[{"testKey":"testValue"},{"testKey2":"testValue2"}]` + "\n",
		},
		{
			name:           "return error when decode error occurs",
			encoderFactory: yamlEncoderDecoderFactory{},
			decoderFactory: jsonEncoderDecoderFactory{},
			src:            `[{"testKey`,
			wantErr:        errors.New("decode: unexpected EOF"),
		},
		{
			name:           "return error when encode error occurs",
			encoderFactory: jsonEncoderDecoderFactory{},
			decoderFactory: yamlEncoderDecoderFactory{},
			src:            `- 1.0: value`,
			wantErr:        errors.New("encode: json: unsupported type: map[interface {}]interface {}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter, err := NewConverter(tt.encoderFactory, tt.decoderFactory)
			require.NoError(t, err)

			dst := &bytes.Buffer{}
			err = converter.Convert(dst, strings.NewReader(tt.src))
			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantDst, dst.String())
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name           string
		encoderFactory EncoderFactory
		decoderFactory DecoderFactory
		wantErr        bool
	}{
		{
			name:           "success create converter when different encoder and decoder passed",
			encoderFactory: jsonEncoderDecoderFactory{},
			decoderFactory: yamlEncoderDecoderFactory{},
			wantErr:        false,
		},
		{
			name:           "error create converter when same encoder and decoder passed",
			encoderFactory: jsonEncoderDecoderFactory{},
			decoderFactory: jsonEncoderDecoderFactory{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewConverter(tt.encoderFactory, tt.decoderFactory)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
