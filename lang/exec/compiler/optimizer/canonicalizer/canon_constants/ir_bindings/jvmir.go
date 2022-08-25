package ir_bindings

var JAVAIR = map[string]string{
	"IR__RESET_BYTE": "*p=0;",
	// %d -> shift
	"IR__ADD_RIGHT": "*(p+%d)+=*p;*p=0;",
	"IR__ADD_LEFT":  "*(p-%d)+=*p;*p=0;",
	"IR__SUB_RIGHT": "*(p+%d)-=*p;*p=0;",
	"IR__SUB_LEFT":  "*(p-%d)-=*p;*p=0;",
	// %d -> shift
	// %d -> constant multiplier
	"IR__MUL_RIGHT": "*(p+%d)+=(*p)*%d;*p=0;",
	"IR__MUL_LEFT":  "*(p-%d)+=(*p)*%d;*p=0;",
}
