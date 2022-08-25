package canonicalizer

// addition
const BF__ADD_LEFT = "[-<+>]"
const BF__ADD_LEFT_ALT = "[<+>-]"
const BF__ADD_RIGHT = "[->+<]"
const BF__ADD_RIGHT_ALT = "[>+<-]"

// subtraction
const BF__SUB_LEFT = "[-<->]"
const BF__SUB_LEFT_ALT = "[<->-]"
const BF__SUB_RIGHT = "[->-<]"
const BF__SUB_RIGHT_ALT = "[>-<-]"

// multiplication
// %s -> multiplier * "+"
const BF__MUL_LEFT = "[-<%s>]"
const BF__MUL_LEFT_ALT = "[<%s>-]"
const BF__MUL_RIGHT = "[->%s<]"
const BF__MUL_RIGHT_ALT = "[>%s<-]"

// %d -> shift
const IR__ADD_RIGHT = "*(p+%d)+=*p;*p=0;"
const IR__ADD_LEFT = "*(p-%d)+=*p;*p=0;"
const IR__SUB_RIGHT = "*(p+%d)-=*p;*p=0;"
const IR__SUB_LEFT = "*(p-%d)-=*p;*p=0;"

// %d -> shift
// %d -> constant multiplier
const IR__MUL_RIGHT = "*(p+%d)+=(*p)*%d;*p=0;"
const IR__MUL_LEFT = "*(p-%d)+=(*p)*%d;*p=0;"
