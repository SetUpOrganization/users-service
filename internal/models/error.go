package models

import "errors"

type HTTPError struct {
	Status  int
	Message string
}

func (e *HTTPError) Err() error {
	return errors.New(e.Message)
}
