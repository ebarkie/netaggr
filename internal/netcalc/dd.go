// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package netcalc

import (
	"net"
	"strings"
)

// parseDD parses the string s as an IPv4 address and subnet mask in
// quad-dotted notation, like:
//
//	192.0.2.0 255.255.255.0
//	192.0.2.0/255.255.255.0
func parseDD(s string) (net.IP, *net.IPNet, error) {
	i := strings.IndexAny(s, "/ ")
	if i < 0 {
		return nil, nil, &net.ParseError{Type: "network address", Text: s}
	}

	ip := net.ParseIP(s[:i])
	m := net.IPMask(net.ParseIP(strings.TrimLeft(s[i+1:], " ")).To4())
	if _, size := m.Size(); ip == nil || size != 32 {
		return nil, nil, &net.ParseError{Type: "network address", Text: s}
	}

	return ip, &net.IPNet{IP: ip.Mask(m), Mask: m}, nil
}

// DD returns the IP network n as a string formatted as an IPv4 address and
// a subnet mask in quad-dotted notation.
func DD(n net.IPNet) string {
	if len(n.IP) != 4 {
		return n.String()
	}

	return n.IP.String() + " " + net.IP(n.Mask).String()
}
