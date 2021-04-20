// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package netcalc

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sort"
	"strings"
)

// Parse parses single addresses or networks formatted as IPv4/6 addresses,
// IPv4/6 CIDR, or an IPv4 address and a dot-decimal subnet mask, like:
//
//      192.0.2.1
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
			return nets, fmt.Errorf("line %d %w", i, err)
		}
		nets = append(nets, n)
	}

	sort.Sort(nets)

	return nets, nil
}

// parseNet parses the string s into an IPNet.
func parseNet(s string) (net.IP, *net.IPNet, error) {
	if strings.Count(s, ".") == 6 {
		return parseDD(s)
	} else if strings.Contains(s, "/") {
		return net.ParseCIDR(s)
	} else {
		return parseIP(s)
	}
}

// parseIP parses the string s formatted as a single IPv4/6 address.
func parseIP(s string) (net.IP, *net.IPNet, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, nil, &net.ParseError{Type: "ip address", Text: s}
	}

	if ip.To4() != nil {
		ip = ip[len(ip)-net.IPv4len:]
	}

	return ip, &net.IPNet{IP: ip, Mask: net.CIDRMask(len(ip)*8, len(ip)*8)}, nil
}
