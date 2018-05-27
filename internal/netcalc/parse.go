// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package netcalc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"sort"
	"strings"
)

// Parse parses a list of IPv4/6 CIDR networks or IPv4 addresses and subnet
// masks in quad-dotted notation, like:
//
//	192.0.2.0/24
//	192.0.2.0 255.255.255.0
//	192.0.2.0/255.255.255.0
//
// It returns a sorted list of Nets.
func Parse(r io.Reader) (Nets, error) {
	var nets Nets
	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		_, n, err := parseNet(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nets, fmt.Errorf("line %d %s", i, err)
		}
		nets = append(nets, n)
	}

	sort.Slice(nets, func(i, j int) bool {
		return bytes.Compare(
			append(nets[i].IP, nets[i].Mask...),
			append(nets[j].IP, nets[j].Mask...)) < 0
	})

	return nets, nil
}

func parseNet(s string) (net.IP, *net.IPNet, error) {
	// IPv4/6 CIDR notation.
	if strings.Count(s, ".") < 6 {
		return net.ParseCIDR(s)
	}

	// IPv4 address and subnet mask in quad-dotted notation.
	return parseDD(s)
}
