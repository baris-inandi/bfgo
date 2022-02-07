package src

import "strconv"

var initByteBoilerplateRequired = false // if the initByte function will be required in the intermediate code

func addCode(x string, init bool) {
	/*
		func addCode
			appends Go code to `intermediate`,
			if init is true, it will also call initByte function
			before the specified code is appended
			to init a byte only if it is actually needed
			E.g. in operations +-[],.
	*/
	if init {
		initByteBoilerplateRequired = true
		addCode("b(p);", false)
	}
	intermediate += x
}

func transpile(code string) {
	/*
		func transpile
			transpiles brainfuck code to Go code and returns it as a string
	*/

	// variables needed to keep track of the transpiler
	fmtBoilerplateRequired := false     // if fmt needs to be imported in the intermediate code
	byteInBoilerplateRequired := false  // if function byteIn will be required in the intermediate code
	byteOutBoilerplateRequired := false // if function byteOut will be needed in the intermediate code

	// transpilation
	code = code + "/"
	needsInit := true
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
					needsInit = true
					addCode("p-="+rep+";", false)
				case ">":
					needsInit = true
					addCode("p+="+rep+";", false)
				case "+":
					addCode("t[p]+="+rep+";", needsInit)
					needsInit = false
				case "-":
					addCode("t[p]-="+rep+";", needsInit)
					needsInit = false
				case ".":
					addCode("o();", true)
					fmtBoilerplateRequired = true
					byteOutBoilerplateRequired = true
				case ",":
					addCode("i();", false)
					fmtBoilerplateRequired = true
					byteInBoilerplateRequired = true
				case "[":
					addCode("for {if t[p]==byte(0) {break};", false)
				case "]":
					addCode("};", true)
				}
				repeatedCharCounter = 1
			}
		} else {
			initialRepeat = true
		}
		prevChar = char
	}
	/*

		func o  -> byteOut - a utility function that prints out the byte the pointer is currently on as a string
		func i  -> byteIn - a function that gets user input from stdin and writes the input to the tape as bytes
		func b  -> initByte - a function that writes 0 to the current byte if not initialized, called whenever init param of addCode() is true
		var  p  -> pointer (of type uint64)
		var  t  -> tape (of type map[uint64]byte{})

	*/
	boilerplate := "package main;"
	if fmtBoilerplateRequired {
		boilerplate += "import \"fmt\";"
	}
	if initByteBoilerplateRequired {
		boilerplate += "var t = map[uint64]byte{}; var p = uint64(0);"
		boilerplate += "func b(x uint64)byte{if v,ok:=t[x];ok {return v};t[x]=byte(0);return byte(0)};"
	}
	if byteInBoilerplateRequired {
		boilerplate += "func i(){var x byte;fmt.Printf(\"> \");fmt.Scanln(&x);t[p]=x;};"
	}
	if byteOutBoilerplateRequired {
		boilerplate += "func o(){fmt.Printf(string(t[p]))};"
	}
	intermediate = boilerplate + "func main(){" + intermediate + "}"
}
