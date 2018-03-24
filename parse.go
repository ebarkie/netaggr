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

// parse parses a list of CIDR network strings from an io.Reader and returns a
// sorted slice of net.IPNet's.
func parse(r io.Reader) ([]*net.IPNet, error) {
	var nets []*net.IPNet

	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		_, n, err := parseNet(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, fmt.Errorf("line %d %s", i, err)
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

	// Traditional IPv4 netmask format.
	i := strings.IndexAny(s, "/ ")
	if i < 0 {
		return nil, nil, &net.ParseError{Type: "network address", Text: s}
	}

	ip := net.ParseIP(s[:i])
	m := net.IPMask(net.ParseIP(s[i+1:]).To4())
	if _, size := m.Size(); ip == nil || size != 32 {
		return nil, nil, &net.ParseError{Type: "network address", Text: s}
	}

	return ip, &net.IPNet{IP: ip.Mask(m), Mask: m}, nil
}
