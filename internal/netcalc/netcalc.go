// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package netcalc performs calculations against IP networks.
//
// It can parse networks formatted as IPv4/6 CIDR or an IPv4 address and
// a dot-decimal subnet mask, and assimilate or aggregate them.
package netcalc

import "net"

// Nets is a sorted slice of IPNet's.  If this is populated by means other
// than Parse then the caller is responsible for sorting.
type Nets []*net.IPNet
