// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package netcalc

import (
	"bytes"
	"math"
	"net"
)

// decrIP decrements an IP address by one.
func decrIP(ip net.IP) net.IP {
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)
	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]--
		if newIP[i] < math.MaxUint8 {
			break
		}
	}

	return newIP
}

// decrPrefix decrements a prefix length by one.
func decrPrefix(mask net.IPMask) net.IPMask {
	newMask := make(net.IPMask, len(mask))
	copy(newMask, mask)
	for i := len(newMask) - 1; i >= 0; i-- {
		if newMask[i] > 0 {
			newMask[i] = newMask[i] << 1
			break
		}
	}

	return newMask
}

// Aggr aggregates networks by joining adjacent networks to form larger
// networks.
func (nets *Nets) Aggr() {
	// The slice of IPNet's are sorted so iterate and if the current net
	// decremented by one is in the previous net (the broadcast address,
	// actually) then they are adjacent.  If the masks/prefixes are the
	// same and decrementing the prefix of the current net does not change
	// it's IP then they can be combined.
	for i := 0; i < len(*nets)-1; {
		if (*nets)[i].Contains(decrIP((*nets)[i+1].IP)) &&
			bytes.Equal((*nets)[i].Mask, (*nets)[i+1].Mask) &&
			(*nets)[i].IP.Equal((*nets)[i].IP.Mask(decrPrefix((*nets)[i].Mask))) {
			(*nets)[i].Mask = decrPrefix((*nets)[i].Mask)
			*nets = append((*nets)[:i+1], (*nets)[i+2:]...)
			// If this isn't the first network then decrement the index by 1
			// to see if the new prefix allows for additional combines.
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
}

// Assim assimilates networks by removing smaller networks that are inside
// larger networks.
func (nets *Nets) Assim() {
	// The slice of IPNet's are sorted so simply iterate and check if the
	// current is in the previous IPNet.
	for i := 1; i < len(*nets); {
		if (*nets)[i-1].Contains((*nets)[i].IP) {
			*nets = append((*nets)[:i], (*nets)[i+1:]...)
		} else {
			i++
		}
	}
}
