package main

import (
	"fmt"
	"io"
)

type fakeConverter struct{}

func (converter fakeConverter) Convert(_ io.Writer, src io.Reader) error {
	bytes, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
