package boilerplate

const JAVA_IR_BOILERPLATE string = `class %s {
    public static void main(String[] args) throws Exception {
        %s[] t = new %s[%d];
        %s p = 0;
        // ir %s
    }
}
`
