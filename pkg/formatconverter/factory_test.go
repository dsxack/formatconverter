package formatconverter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEncoderFactoryByFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     EncoderFactory
		wantErr  bool
	}{
		{
			name:     "success return encoder factory when passed filename with supported extension",
			filename: "test.json",
			want:     jsonEncoderDecoderFactory{},
			wantErr:  false,
		},
		{
			name:     "return error when passed filename with unsupported extension",
			filename: "test.unsupported",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoderFactoryByFilename(tt.filename)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewDecoderFactoryByFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     EncoderFactory
		wantErr  bool
	}{
		{
			name:     "success return decoder factory when passed filename with supported extension",
			filename: "test.json",
			want:     jsonEncoderDecoderFactory{},
			wantErr:  false,
		},
		{
			name:     "return error when passed filename with unsupported extension",
			filename: "test.unsupported",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecoderFactoryByFilename(tt.filename)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
