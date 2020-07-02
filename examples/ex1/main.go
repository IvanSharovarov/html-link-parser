package main

import (
	"fmt"
	html_link_parser "github.com/IvanSharovarov/html-link-parser"
	"strings"
)

var html = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(html)
	links, err := html_link_parser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
