package onfleet

import (
	"encoding/json"
	"fmt"
	"io"
)

type RequestErrorMessage struct {
	Cause any `json:"cause,omitempty"`
	// Error is an internal error code.
	// It is different than the request status code.
	Error int `json:"error,omitempty"`
	// Message is the error messages / description
	Message string `json:"message,omitempty"`
	// RemoteAddress is remote ip
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Request is uuid associated with the request
	Request string `json:"request,omitempty"`
	// StatusCode only present on errors returned for batch task creation
	StatusCode int `json:"statusCode,omitempty"`
}

type RequestError struct {
	// Code is error type e.g. "InvalidArgument"
	Code string `json:"code,omitempty"`
	// Message contains futher details about the error.
	Message RequestErrorMessage `json:"message"`
}

func (err RequestError) Error() string {
	return fmt.Sprintf("%s: \n  Cause: %s\n  Message: %s", err.Code, err.Message.Cause, err.Message.Message)
}

func ParseError(r io.Reader) error {
	var reqError RequestError
	if err := json.NewDecoder(r).Decode(&reqError); err != nil {
		return err
	}
	return reqError
}
