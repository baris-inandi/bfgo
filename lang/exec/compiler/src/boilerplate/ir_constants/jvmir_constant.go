package ir_constants

var JAVAIR = map[string]string{
	"[":               "while(t[p]!=0){",
	"]":               "}",
	"<":               "--p;",
	">":               "++p;",
	"+":               "++t[p];",
	"-":               "--t[p];",
	".":               "System.out.print((char)t[p]);",
	",":               "t[p]=System.in.read();",
	"LEFT_ANGLE_REP":  "p-=%d;",
	"RIGHT_ANGLE_REP": "p+=%d;",
	"PLUS_REP":        "t[p]+=%d;",
	"MINUS_REP":       "t[p]-=%d;",
}
