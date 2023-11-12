package main

import "errors"

var (
	errGetRequest     = errors.New("failed to send GET request:")
	errResponseStatus = errors.New("not OK status")
	errWrite          = errors.New("failed to write to file:")
)
