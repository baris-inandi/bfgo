package ir_constants

var JSIR = map[string]string{
	"[":               "while(t[p]!=0){",
	"]":               "}",
	"<":               "--p;",
	">":               "++p;",
	"+":               "++t[p];",
	"-":               "--t[p];",
	".":               "o.innerHTML+=String.fromCharCode(t[p]);",
	",":               "a=prompt('Brainfuck: input');for(let i=0;i<a.length;i++){t[p+i]=a[i]}",
	"LEFT_ANGLE_REP":  "p-=%d;",
	"RIGHT_ANGLE_REP": "p+=%d;",
	"PLUS_REP":        "t[p]+=%d;",
	"MINUS_REP":       "t[p]-=%d;",
}
