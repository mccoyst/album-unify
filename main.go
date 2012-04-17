// Copyright © Steve McCoy.
// Licensed under the MIT License.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	names := flag.Args()

	p := prefix(names)
	if p == "" {
		fmt.Fprintln(os.Stderr, "The names must share a non-empty prefix.")
		os.Exit(1)
	}

	err := os.Mkdir(p, os.ModeDir|0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making %q: %v\n", p, err)
		os.Exit(1)
	}

	for i := range names {
		err = moveFiles(i+1, p, names[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error moving files from %q: %v\n",
				names[i], err)
			os.Exit(1)
		}

		err = os.RemoveAll(names[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting %q: %v\n",
				names[i], err)
			os.Exit(1)
		}
	}
}

func prefix(names []string) string {
	if len(names) == 0 {
		return ""
	}

	p := names[0]
	e := len(p)
	for _, n := range names[1:len(names)] {
		for i, b := range []byte(n) {
			if b != p[i] {
				e = i
				continue
			}
		}
	}

	p = p[0:e]
	return trimJunk(p)
}

func moveFiles(id int, dest, srcname string) error {
	src, err := os.Open(srcname)
	if err != nil {
		return err
	}
	defer src.Close()

	files, err := src.Readdirnames(0)
	if err != nil {
		return err
	}

	i := strconv.Itoa(id)
	for _, f := range files {
		s := filepath.Join(srcname, f)
		d := filepath.Join(dest, i+"-"+f)
		err = os.Rename(s, d)
		if err != nil {
			return err
		}
	}

	return nil
}

// trimJunk returns s without typical garbage like disc labels.
func trimJunk(s string) string {
	typicalJunk := " [Disc "
	if strings.HasSuffix(s, typicalJunk) {
		return s[0 : len(s)-len(typicalJunk)]
	}
	return strings.TrimSpace(s)
}
