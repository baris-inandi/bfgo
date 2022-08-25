package strip

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Strip(f string, wd string) {
	stripstdout := &bytes.Buffer{}
	stripstderr := &bytes.Buffer{}
	stripCommand := fmt.Sprintf("strip --strip-unneeded %s", f)
	stripcmd := exec.Command("bash", "-c", stripCommand)
	stripcmd.Stderr = stripstderr
	stripcmd.Stdout = stripstdout
	stripcmd.Dir = wd
	err := stripcmd.Run()
	if err != nil {
		fmt.Println("WARN: Cannot strip binary\n", err)
	}
}
