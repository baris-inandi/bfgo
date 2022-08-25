package lang

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var verboseOutBuffer = ""

func (c *Code) VerboseOut(s ...interface{}) {
	if c.Context.Bool("debug") || c.Context.Bool("verbose") {
		elapsedTime := time.Since(c.startTime).Milliseconds()
		out := fmt.Sprint(s...) + fmt.Sprintf(" [%dms]", elapsedTime)
		verboseOutBuffer += out + "\n"
		if c.Context.Bool("verbose") {
			fmt.Println(out)
		}
	}
}

func (c *Code) AddDebugFile(filename string, content string) {
	if c.IsDebugging {
		c.DebugFiles[filename] = content
	}
}

func (c *Code) WriteDebugFiles(basepath string) {
	c.AddDebugFile("VERBOSE-OUT", verboseOutBuffer)
	for k, v := range c.DebugFiles {
		err := os.WriteFile(
			filepath.Join(basepath, k),
			[]byte(v), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
