// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Network aggregator/summarizer.
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/ebarkie/netaggr/pkg/netcalc"
)

func main() {
	doAggr := flag.Bool("aggr", true, "perform network aggregation")
	doAssim := flag.Bool("assim", true, "perform network assimilation")
	in := flag.String("in", "", "input file")
	notation := flag.String("notation", "cidr", "output notation: \"cidr\" or \"dd\"")
	flag.Parse()

	var s func(n net.IPNet) string
	switch strings.ToLower(*notation) {
	case "cidr":
		s = func(n net.IPNet) string { return n.String() }
	case "dd":
		s = netcalc.DD
	default:
		fmt.Fprintf(os.Stderr, "Invalid output notation: %s\n", *notation)
		os.Exit(1)
	}

	var r io.Reader
	if *in == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(*in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Input error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	}

	nets, err := netcalc.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %s\n", err)
		os.Exit(1)
	}

	if *doAssim {
		nets.Assim()
	}

	if *doAggr {
		nets.Aggr()
	}

	for _, n := range nets {
		fmt.Println(s(*n))
	}
}
