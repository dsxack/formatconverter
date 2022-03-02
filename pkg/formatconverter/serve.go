package formatconverter

import (
	"bytes"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"net/http"
	"strings"
)

func Serve(writer http.ResponseWriter, request *http.Request) {
	srcBytes, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("error when reading body: %v", err), http.StatusBadRequest)
		return
	}

	mtype := mimetype.Detect(srcBytes)
	decoderFactory, err := NewDecoderFactoryByMimeType(mtype.String())
	if err != nil {
		http.Error(writer, fmt.Sprintf("error when creating decoder: %v", err), http.StatusBadRequest)
		return
	}

	dstFormat := request.URL.Query().Get("dstFormat")
	if dstFormat == "" {
		http.Error(writer, "dstFormat query parameter is required", http.StatusBadRequest)
		return
	}
	dstFormat = strings.ToLower(dstFormat)

	encoderFactory, err := NewEncoderFactoryByFormat(dstFormat)
	if err != nil {
		http.Error(writer, fmt.Sprintf("error when creating encoder: %v", err), http.StatusBadRequest)
		return
	}

	converter, err := NewFormatConverter(encoderFactory, decoderFactory)
	if err != nil {
		http.Error(writer, fmt.Sprintf("error when creating converter: %v", err), http.StatusBadRequest)
		return
	}

	err = converter.Convert(writer, bytes.NewReader(srcBytes))
	if err != nil {
		http.Error(writer, fmt.Sprintf("error when converting: %v", err), http.StatusInternalServerError)
		return
	}
}
