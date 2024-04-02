package generate_out_file

import (
	"path/filepath"
	"strings"

	"github.com/baris-inandi/bfgo/lang"
)

func GenerateOutFile(c lang.Code) string {
	fileIn := c.Filepath
	specifiedName := c.Context.Path("output")

	out := ""
	if specifiedName == "" {
		fileInNameSplit := strings.Split(fileIn, "/")
		fileInName := fileInNameSplit[len(fileInNameSplit)-1]
		fileInNameDotSplit := strings.Split(fileInName, ".")
		out = fileInNameDotSplit[0]
		if c.UsingJS() {
			out += ".js"
		}
	} else {
		out = specifiedName
	}
	return filepath.Join(out)
}
