package netcalc

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

func ExampleDiff() {
	a, _ := Parse(strings.NewReader(
		`10.10.1.0/24
		10.10.2.0/24
		10.10.3.0/24
		10.10.4.0/24
		192.168.1.0/24`))

	b, _ := Parse(strings.NewReader(
		`10.10.1.0/24
		10.10.3.0/24
		10.10.5.0/24
		192.168.1.0/24
		192.168.2.0/24
		192.168.3.0/25`))

	added, deleted := Diff(a, b)

	fmt.Printf("Added %s", added)
	fmt.Printf("Deleted %s", deleted)
	// Output:
	// Added 3 network(s):
	// 	10.10.5.0/24
	// 	192.168.2.0/24
	// 	192.168.3.0/25
	// Deleted 2 network(s):
	// 	10.10.2.0/24
	// 	10.10.4.0/24
}

func ExampleSort() {
	nets := Nets{
		&net.IPNet{IP: net.IP{10, 10, 3, 0}, Mask: net.CIDRMask(24, 32)},
		&net.IPNet{IP: net.IP{192, 168, 1, 0}, Mask: net.CIDRMask(24, 32)},
		&net.IPNet{IP: net.IP{10, 10, 1, 0}, Mask: net.CIDRMask(24, 32)},
		&net.IPNet{IP: net.IP{10, 10, 4, 0}, Mask: net.CIDRMask(24, 32)},
		&net.IPNet{IP: net.IP{10, 10, 2, 0}, Mask: net.CIDRMask(25, 32)},
	}

	sort.Sort(nets)

	fmt.Println(nets)
	// Output:
	// 5 network(s):
	// 	10.10.1.0/24
	// 	10.10.2.0/25
	// 	10.10.3.0/24
	// 	10.10.4.0/24
	// 	192.168.1.0/24
}
