package src

import (
	"fmt"
	"strconv"
)

const CXX string = `#include <iostream>
#include <unordered_map>
using namespace std;

class Bf
{
public:
    unordered_map<uint_fast16_t, uint_fast8_t> _t =
        unordered_map<uint_fast16_t, uint_fast8_t>();
    uint_fast16_t _p = 0;
    uint_fast8_t c() { return _t[_p]; }          // current
    bool w() { return c() != 0; }                // while, false if current is 0
    void p(uint_fast8_t x) { _t[_p] = c() + x; } // plus, increment
    void m(uint_fast8_t x) { _t[_p] = c() - x; } // minus, decrement
    void l(uint_fast8_t x) { _p -= x; }          // left
    void r(uint_fast8_t x) { _p += x; }          // right
    void o() { printf("%%c", (char)c()); }        // out, print
    void i()
    {
        string x;
        getline(cin, x);
        for (uint_fast8_t i = 0; i < x.length(); i++)
        {
            _t[_p + i] = x[i];
        }
    } // in, read
};

void impl(Bf b) {%s}

int main()
{
    Bf b;
    impl(b);
    return 0;
};
`

func transpile(code string) string {
	// transpiles brainfuck code to c++ code and returns it as a string
	intermediate := ""
	code += "/"
	prevChar := ""
	repeatedCharCounter := 1
	initialRepeat := false
	for _, char := range code {
		char := string(char)
		if initialRepeat {
			if prevChar == char && (prevChar == "+" || prevChar == "-" || char == "<" || char == ">") {
				repeatedCharCounter += 1
			} else {
				rep := strconv.Itoa(repeatedCharCounter)
				switch prevChar {
				case "<":
					intermediate += ("b.l(" + rep + ");")
				case ">":
					intermediate += ("b.r(" + rep + ");")
				case "+":
					intermediate += ("b.p(" + rep + ");")
				case "-":
					intermediate += ("b.m(" + rep + ");")
				case ".":
					intermediate += ("b.o();")
				case ",":
					intermediate += ("b.i();")
				case "[":
					intermediate += ("while(b.w()){")
				case "]":
					intermediate += ("};")
				}
				repeatedCharCounter = 1
			}
		} else {
			initialRepeat = true
		}
		prevChar = char
	}
	intermediate = fmt.Sprintf(CXX, intermediate)
	return intermediate
}
