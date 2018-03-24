// Copyright (c) 2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

const testDataDir = "testdata/"

func TestParse(t *testing.T) {
	testFiles, err := filepath.Glob(testDataDir + "*.in")
	if err != nil {
		t.Fatal(err)
	}

	for _, tf := range testFiles {
		t.Logf("Parse test: %s", tf)
		f, err := os.Open(tf)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		_, err = parse(bufio.NewReader(f))
		if err != nil {
			t.Error(err)
		}
	}
}
