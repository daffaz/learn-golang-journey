package main

import (
	"fmt"
	"htmlparser"
	"strings"
)

var exampleHTML string = `
<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
		<a href="/page-two">A link to a second page</a>
	</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := htmlparser.Parser(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
