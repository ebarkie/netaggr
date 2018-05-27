// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package netcalc performs calculations against IP networks.
//
// It can parse a list of IPv4/6 CIDR networks or IPv4 addresses and subnet
// masks in quad-dotted notation and assimilate or aggregate them.
package netcalc

import "net"

// Nets is a sorted slice of IPNet's.  If this is populated by means other
// than Parse then the caller is responsible for sorting.
type Nets []*net.IPNet
