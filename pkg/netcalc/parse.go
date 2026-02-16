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
//	192.0.2.1
//	192.0.2.0/24
//	192.0.2.0 255.255.255.0
//	192.0.2.0/255.255.255.0
//
// It returns a sorted list of Nets.
func Parse(r io.Reader) (nets Nets, err error) {
	netC := make(chan *net.IPNet)
	go func() {
		defer close(netC)
		_, err = ReadFrom(r, netC)
	}()

	for n := range netC {
		nets = append(nets, n)
	}
	sort.Sort(nets)

	return
}

// ReadFrom parses the io.Reader and sends the resulting IPNets to the netC
// channel.  This is a useful construct for concurrently parsing Nets.
func ReadFrom(r io.Reader, netC chan<- *net.IPNet) (int64, error) {
	scanner := bufio.NewScanner(r)
	var i int64
	for ; scanner.Scan(); i++ {
		s := trim(scanner.Text())
		if s == "" {
			continue
		}
		_, n, err := ParseNet(s)
		if err != nil {
			return i, fmt.Errorf("line %d %w", i+1, err)
		}

		netC <- n
	}

	return i, scanner.Err()
}

// ParseNet parses the string s into an IPNet.
func ParseNet(s string) (net.IP, *net.IPNet, error) {
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

// trim removes all comments and leading+trailing white space from a string.
func trim(s string) string {
	if i := strings.IndexAny(s, "#;"); i > -1 {
		return strings.TrimSpace(s[:i])
	}

	return strings.TrimSpace(s)
}
