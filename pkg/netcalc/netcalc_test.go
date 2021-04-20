package netcalc

import (
	"fmt"
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

	fmt.Printf("ADDED\n%s", added)
	fmt.Printf("DELETED\n%s", deleted)
	// Output:
	// ADDED
	// 10.10.5.0/24
	// 192.168.2.0/24
	// 192.168.3.0/25
	// DELETED
	// 10.10.2.0/24
	// 10.10.4.0/24
}

func ExampleSort() {
	n, _ := Parse(strings.NewReader(
		`10.10.3.0/24
		192.168.1.0/24
		10.10.1.0/24
		10.10.4.0/24
		10.10.2.0/25`))

	sort.Sort(n)

	fmt.Println(n)
	// Output:
	// 10.10.1.0/24
	// 10.10.2.0/25
	// 10.10.3.0/24
	// 10.10.4.0/24
	// 192.168.1.0/24
}
