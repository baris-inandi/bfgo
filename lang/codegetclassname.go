package lang

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func toValidClassName(in string) string {
	split1 := strings.Split(in, "/")
	split1out := split1[len(split1)-1]
	split2 := strings.Split(split1out, "\\")
	s := split2[len(split2)-1]
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	valid := reg.ReplaceAllString(s, "")
	if valid == "" {
		fmt.Println("Invalid JVM class name")
		os.Exit(1)
	}
	caser := cases.Title(language.English)
	return caser.String(valid)
}

func (c *Code) GetClassName() string {
	specified := c.Context.String("output")
	if specified != "" {
		return toValidClassName(specified)
	} else {
		return toValidClassName(c.Filepath)
	}
}
