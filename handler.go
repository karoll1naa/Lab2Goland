package lab2

import (
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}
	result, err := PostfixToInfix(string(data))
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(result))
	if err != nil {
		return err
	}
	return nil
}
