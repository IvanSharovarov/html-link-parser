package html_link_parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, p string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(p, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, p+" ")
	}
}
