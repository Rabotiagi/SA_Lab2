package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}


func (ch *ComputeHandler) Compute() error {
	data := make([]byte, 1024)
	count, err1 := ch.Input.Read(data)
	if err1 != nil {
		return err1
	}

	res, err2 := PostfixToInfix(string(data[:count]))
	if err2 != nil {
		return fmt.Errorf("invalid input")
	}

	_, err3 := ch.Output.Write([]byte(res + "\n"))
	if err3 != nil {
		return err3
	}

	return nil
}