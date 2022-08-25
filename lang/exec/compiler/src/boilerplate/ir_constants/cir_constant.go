package ir_constants

var CIR = map[string]string{
	"[":               "while(*p){",
	"]":               "};",
	"<":               "--p;",
	">":               "++p;",
	"+":               "++*p;",
	"-":               "--*p;",
	".":               "putc(*p, stdout);",
	",":               "*p=getchar();",
	"LEFT_ANGLE_REP":  "p-=%d;",
	"RIGHT_ANGLE_REP": "p+=%d;",
	"PLUS_REP":        "*p+=%d;",
	"MINUS_REP":       "*p-=%d;",
}
