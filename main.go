// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Network aggregator/summarizer.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	doAggr := flag.Bool("aggr", true, "perform network aggregation")
	doAssim := flag.Bool("assim", true, "perform network assimilation")
	in := flag.String("in", "", "input file")
	flag.Parse()

	var r io.Reader
	if *in == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(*in)
		if err != nil {
			fmt.Printf("Input error: %s\n", err.Error())
			return
		}
		defer f.Close()
		r = bufio.NewReader(f)
	}

	nets, err := parse(r)
	if err != nil {
		fmt.Printf("Parse error: %s", err.Error())
		return
	}

	if *doAssim {
		assim(&nets)
	}

	if *doAggr {
		aggr(&nets)
	}

	for _, n := range nets {
		fmt.Println(n)
	}
}
