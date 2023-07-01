package cmd

import "errors"

var (
	ErrInvalidPath            = errors.New("the given directory path is not valid")
	ErrUnrecognizedSubcommand = errors.New("unrecognized subcommand")
)
