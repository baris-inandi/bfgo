# Brainfuck

brainfuck-go is an overengineered brainfuck toolkit written in Go.

## Features

The toolkit includes the following features:

- Compiler (binary)
- JVM Compiler
- JavaScript Compiler
- Interpreter
- REPL
- Formatter
- Minifier

### Compiler

The compiler can compile for three different targets:

- Binary for native execution
- JVM bytecode
- JavaScript for running in the browser

#### Machine Code

**Example:** `$ brainfuck examples/hello.bf`  
Creates a binary `hello`. Run with `./hello`.

#### JVM Compiler

**Example:** `$ brainfuck -jvm examples/hello.bf`  
Generates a JVM classfile `Hello.class`. Run with `java Hello`.

#### JavaScript Compiler

**Example:** `$ brainfuck -js examples/hello.bf`  
Generates a JavaScript file `hello.js` and an HTML file `hello.html`.  
Run your favorite HTTP server and load `hello.html` in the browser.
The output will replace the document `<body>`.

### The Optimizer

All compile targets can be compiled with **the optimizer**. The optimizer options are:

- `-F`: Fast
  - Uses fast IR generation
  - Results in fast compile times
  - Causes Slow execution
  - > use `-o-compile` or `-F`
- `-B`: Balanced
  - Default behaviour
  - Applies some optimizations
  - Balance between -F and -O
  - > use `-o-balanced` or `-B`
- `-O`: Optimized
  - Uses the full optimizer
    - Dead code elimination
    - Canonicalization
  - Smaller binary size
    - Also performs binary stripping
  - Causes Slow compile times
  - Results in very fast execution
  - > use `-o-performance`

Executes given brainfuck file.
There is still room for improvement when it comes to performance. Feel free to submit a PR.

> use `-interpret`

#### REPL

The REPL is a command line interface for the interpreter.
It can be used to execute brainfuck interactively.

> use `-repl`

### bffmt

Brainfuck formatter bundled with `brainfuck-go`.  
> Warning: bffmt currently omits all comments. Feel free to submit a PR for support for comments.  
> use `-fmt`

Example formatted snippet from `examples/fibonacci.bf`:

```brainfuck
  [
    +++++
    [>++++++++<-]> .
    <++++++
    [>--------<-]+<<<
  ]
```

#### Minifier

bffmt can also minify brainfuck code, leaving only valid characters, minimizing file size.
