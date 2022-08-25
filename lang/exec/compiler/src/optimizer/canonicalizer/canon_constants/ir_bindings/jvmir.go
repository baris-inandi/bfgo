package ir_bindings

var JAVAIR = map[string]string{
	"IR__RESET_BYTE": "t[p]=0;",
	// %d -> shift
	"IR__ADD_RIGHT": "t[p+%d]+=t[p];t[p]=0;",
	"IR__ADD_LEFT":  "t[p-%d]+=t[p];t[p]=0;",
	"IR__SUB_RIGHT": "t[p+%d]-=t[p];t[p]=0;",
	"IR__SUB_LEFT":  "t[p-%d]-=t[p];t[p]=0;",
	// %d -> shift
	// %d -> constant multiplier
	"IR__MUL_RIGHT": "t[p+%d]+=(t[p])*%d;t[p]=0;",
	"IR__MUL_LEFT":  "t[p-%d]+=(t[p])*%d;t[p]=0;",
}
