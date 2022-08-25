package lang

func (c *Code) UseJVM() {
	c.compileTarget = "jvm"
}

func (c *Code) UsingJVM() bool {
	return c.compileTarget == "jvm"
}
