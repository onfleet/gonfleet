package util

import (
	"encoding/json"
	"fmt"
	"io"
)

type RequestErrorMessage struct {
	// Error is an internal error code.
	// It is different than the request status code.
	Error int `json:"error"`
	// Message is the error messages / description
	Message string `json:"message"`
	// RemoteAddress is remote ip
	RemoteAddress string `json:"remoteAddress"`
	// Request is uuid associated with the request
	Request string `json:"request"`
}

type RequestError struct {
	// Code is error type e.g. "InvalidArgument"
	Code string `json:"code"`
	// Message contains futher details about the error.
	Message RequestErrorMessage `json:"message"`
}

func (err RequestError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message.Message)
}

func ReadRequestError(r io.Reader) error {
	var reqError RequestError
	if err := json.NewDecoder(r).Decode(&reqError); err != nil {
		return err
	}
	return reqError
}
