package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	bf "github.com/baris-inandi/brainfuck/src"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"compile", "b"},
				Usage:   "Compiles given brainfuck code to binary",
				Action: func(c *cli.Context) error {
					bf.Compile("mandelbrot.bf", "hello.out", false)
					return nil
				},
			},
			{
				Name:    "interpret",
				Aliases: []string{"inter", "i"},
				Usage:   "Interprets given brainfuck code",
				Action: func(c *cli.Context) error {
					// bf.Interpret("hello.bf", "hello.out", false)
					return nil
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Compiles code and immediately runs binary",
				Action: func(c *cli.Context) error {
					bf.Compile("hello.bf", "", true)
					return nil
				},
			},
			{
				Name:  "repl",
				Usage: "Starts brainfuck interactive shell",
				Action: func(c *cli.Context) error {
					bf.Repl()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
