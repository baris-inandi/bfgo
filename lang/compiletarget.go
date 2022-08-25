package lang

func (c *Code) UseJVM() {
	c.CompileTarget = "java"
}

func (c *Code) UsingJVM() bool {
	return c.CompileTarget == "java"
}

func (c *Code) UseJS() {
	c.CompileTarget = "js"
}

func (c *Code) UsingJS() bool {
	return c.CompileTarget == "js"
}
