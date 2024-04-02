// Copyright (C) 2022 baris-inandi
//
// This file is part of bfgo.
//
// bfgo is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// bfgo is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with bfgo.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
	"os"

	bf "github.com/baris-inandi/bfgo/cli"
)

func main() {
	err := bf.Cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
