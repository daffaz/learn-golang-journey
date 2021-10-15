package htmlparser

import (
	"io"
)

type Link struct {
	Href string
	Text string
}

func Parser(r io.Reader) ([]Link, error) {
	return nil, nil
}
