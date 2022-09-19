# Brainfuck

brainfuck-go is an overengineered brainfuck toolkit written in Go.

## Features

The toolkit includes the following features:

- Native Compiler
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

#### Native

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
  - > use `-o-performance` or `-O`

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

## Cli Flags

```plaintext
--run, -r                  Immediately run binary after compilation (default: false)
--output value, -o value   Specify output binary
--compile-only, -C         Only compile, do not output a binary (default: false)
--clang                    Use clang instead of default gcc (default: false)
--jvm                      Compile to JVM bytecode (default: false)
--js                       Compile to JavaScript (default: false)
--o-compile, -F            Disable optimizations and use fast compiler: fast compile time, slow execution (default: false)
--o-balanced, -B           Minimal optimizations for balanced compile time and performance, default behavior (default: false)
--o-performance, -O        Enable optimizations: fast execution, slow compile time (default: false)
--interpret                Interpret file instead of compiling (default: false)
--repl                     Start a read-eval-print loop (default: false)
--c-compiler-flags value   Pass arbitrary flags to the compiler (gcc, clang or javac)
--c-tape-size value        Integer to specify length of brainfuck tape (default: 30000)
--c-tape-init value        Integer value used to initialize all elements in brainfuck tape (default: 0)
--c-cell-type value        Type used for brainfuck tape in intermediate representation (default: "int")
--d-dump-ir                Dump intermediate representation (default: false)
--d-keep-temp              Do not remove temporary IR files (default: false)
--d-print-ir-filepath      Dump temporary IR filepath, use -d-keep-temp to keep them from being deleted (default: false)
--d-print-compile-command  Print C IR compiler command (default: false)
--verbose, -v              Print verbose output (default: false)
--debug, -d                Produce debug output, overrides -o (default: false)
--time, -t                 Prints out execution time before exiting (default: false)
--fmt                      Format code (omits comments) (default: false)
--minify                   Minify code (default: false)
--help, -h                 show help (default: false)
```

## Benchmark

The following is a benchmark of `examples/mandelbrot.bf`

| Optimization Level | -F      | -B         | -O         |
| ------------------ | ------- | ---------- | ---------- |
| Native (arm64)     | 8 secs  | 580 millis | 370 millis |
| Native (x64)       | 16 secs | 710 millis | 440 millis |
| JVM                | 22 secs | 13 secs    | 13 secs    |
| JavaScript         | 35 secs | 19 secs    | 5 secs     |

> Native arm64 using entry level M2 MacBook Air
> Native x64 using Ryzen 5 3600
> JavaScript using Google Chrome 106.0.5245.0 dev
