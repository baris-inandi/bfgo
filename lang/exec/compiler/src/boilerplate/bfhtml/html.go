package bfhtml

import (
	"fmt"
	"os"
	"strings"
)

const BOILERPLATE = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="Description of your page">
  <script defer src="/%s"></script>
  <title>Brainfuck</title>
</head>
<body></body>
</html>
`

func GenerateHTMLForJSFile(f string) {
	content := fmt.Sprintf(BOILERPLATE, f)
	htmlFile := strings.ReplaceAll(f, ".js", "")
	htmlFile += ".html"
	err := os.WriteFile(htmlFile, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
