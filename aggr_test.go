// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func BenchmarkAggr(b *testing.B) {
	f, err := os.Open(testDataDir + "test5.in")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	nets, _ := parse(bufio.NewReader(f))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		nets.aggr()
	}
}

func BenchmarkAssim(b *testing.B) {
	f, err := os.Open(testDataDir + "test5.in")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	nets, _ := parse(bufio.NewReader(f))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		nets.assim()
	}
}

func TestAggr(t *testing.T) {
	testFiles, err := filepath.Glob(testDataDir + "*.in")
	if err != nil {
		t.Fatal(err)
	}

	for _, tf := range testFiles {
		t.Logf("Aggregate test: %s", tf)
		f, err := os.Open(tf)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		nets, _ := parse(bufio.NewReader(f))

		nets.assim()
		nets.aggr()

		var b bytes.Buffer
		for _, n := range nets {
			b.WriteString(n.String())
			b.WriteString("\n")
		}

		out, _ := ioutil.ReadFile(strings.TrimSuffix(tf, "in") + "out")

		if !bytes.Equal(b.Bytes(), out) {
			t.Errorf("\nExpected:\n%s\nGot:\n%s\n", out, b.Bytes())
		}
	}
}
