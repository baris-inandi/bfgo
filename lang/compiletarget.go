package lang

func (c *Code) UseJVM() {
	c.CompileTarget = "java"
}

func (c *Code) UsingJVM() bool {
	return c.CompileTarget == "java"
}
