// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package netcalc performs calculations against IP networks.
//
// It can parse networks formatted as IPv4/6 CIDR or an IPv4 address and
// a dot-decimal subnet mask, and assimilate or aggregate them.
package netcalc

import (
	"bytes"
	"net"
)

// Nets is a sorted slice of IPNet's.  If this is populated by means other
// than Parse then the caller is responsible for sorting.
type Nets []*net.IPNet

func (n Nets) String() string {
	buf := &bytes.Buffer{}

	for _, net := range n {
		buf.WriteString(net.String())
		buf.WriteRune('\n')
	}

	return buf.String()
}

// sort.Interface implementation.
func (n Nets) Len() int           { return len(n) }
func (n Nets) Less(i, j int) bool { return Compare(*n[i], *n[j]) < 0 }
func (n Nets) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// Compare returns an integer comparing two IPNet's lexicographically. The
// result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare(a, b net.IPNet) int {
	return bytes.Compare(append(a.IP, a.Mask...), append(b.IP, b.Mask...))
}

// Diff finds the differences between two slices of networks.  It returns a
// slice of added networks that don't exist in a but do in b, and a slice of
// deleted networks that exist in a but don't in b.
func Diff(a, b Nets) (Nets, Nets) {
	var added, deleted Nets

	for i, j := 0, 0; i < len(a) || j < len(b); {
		// If we hit the end of the first list before the second then
		// anything else in the second are adds.
		if i == len(a) {
			added = append(added, b[j:]...)
			break
		}

		// Similarly, if we hit the end of the second list before the
		// first then anything else in the first are deletes.
		if j == len(b) {
			deleted = append(deleted, a[i:]...)
			break
		}

		// Lexicographic order tells us if an add or delete took place.
		c := Compare(*a[i], *b[j])
		if c < 0 {
			deleted = append(deleted, a[i])
			i++
		} else if c > 0 {
			added = append(added, b[j])
			j++
		} else {
			i++
			j++
		}
	}

	return added, deleted
}
