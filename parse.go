// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"sort"
	"strings"
)

type ipNets []*net.IPNet

// parse parses a list of CIDR network strings from an io.Reader and returns a
// sorted slice of net.IPNet's.
func parse(r io.Reader) (ipNets, error) {
	var nets ipNets
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
	// IPv4/6 CIDR format.
	if strings.Count(s, ".") < 6 {
		return net.ParseCIDR(s)
	}

	// IPv4 dotted decimal subnet mask format.
	return parseDD(s)
}
