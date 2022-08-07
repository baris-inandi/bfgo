// Copyright (C) 2022 baris-inandi
//
// This file is part of brainfuck.
//
// brainfuck is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.
//
// brainfuck is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with brainfuck.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
	"os"

	bf "github.com/baris-inandi/brainfuck-go/src"
)

func main() {
	err := bf.Cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
