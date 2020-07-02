package html_link_parser

import "io"

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
