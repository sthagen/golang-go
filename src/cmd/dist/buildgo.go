// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/*
 * Helpers for building cmd/go and cmd/cgo.
 */

// mkzdefaultcc writes zdefaultcc.go:
//
//	package main
//	const defaultCC = <defaultcc>
//	const defaultCXX = <defaultcxx>
//	const defaultPkgConfig = <defaultpkgconfig>
//
// It is invoked to write cmd/go/internal/cfg/zdefaultcc.go
// but we also write cmd/cgo/zdefaultcc.go
func mkzdefaultcc(dir, file string) {
	if strings.Contains(file, filepath.FromSlash("go/internal/cfg")) {
		var buf strings.Builder
		fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
		fmt.Fprintln(&buf)
		fmt.Fprintf(&buf, "package cfg\n")
		fmt.Fprintln(&buf)
		fmt.Fprintf(&buf, "const DefaultPkgConfig = `%s`\n", defaultpkgconfig)
		buf.WriteString(defaultCCFunc("DefaultCC", defaultcc))
		buf.WriteString(defaultCCFunc("DefaultCXX", defaultcxx))
		writefile(buf.String(), file, writeSkipSame)
		return
	}

	var buf strings.Builder
	fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "package main\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "const defaultPkgConfig = `%s`\n", defaultpkgconfig)
	buf.WriteString(defaultCCFunc("defaultCC", defaultcc))
	buf.WriteString(defaultCCFunc("defaultCXX", defaultcxx))
	writefile(buf.String(), file, writeSkipSame)
}

func defaultCCFunc(name string, defaultcc map[string]string) string {
	var buf strings.Builder

	fmt.Fprintf(&buf, "func %s(goos, goarch string) string {\n", name)
	fmt.Fprintf(&buf, "\tswitch goos+`/`+goarch {\n")
	var keys []string
	for k := range defaultcc {
		if k != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&buf, "\tcase %s:\n\t\treturn %s\n", quote(k), quote(defaultcc[k]))
	}
	fmt.Fprintf(&buf, "\t}\n")
	if cc := defaultcc[""]; cc != "" {
		fmt.Fprintf(&buf, "\treturn %s\n", quote(cc))
	} else {
		clang, gcc := "clang", "gcc"
		if strings.HasSuffix(name, "CXX") {
			clang, gcc = "clang++", "g++"
		}
		fmt.Fprintf(&buf, "\tswitch goos {\n")
		fmt.Fprintf(&buf, "\tcase ")
		for i, os := range clangos {
			if i > 0 {
				fmt.Fprintf(&buf, ", ")
			}
			fmt.Fprintf(&buf, "%s", quote(os))
		}
		fmt.Fprintf(&buf, ":\n")
		fmt.Fprintf(&buf, "\t\treturn %s\n", quote(clang))
		fmt.Fprintf(&buf, "\t}\n")
		fmt.Fprintf(&buf, "\treturn %s\n", quote(gcc))
	}
	fmt.Fprintf(&buf, "}\n")

	return buf.String()
}

// mkzcgo writes zcgo.go for the go/build package:
//
//	package build
//	const defaultCGO_ENABLED = <CGO_ENABLED>
//
// It is invoked to write go/build/zcgo.go.
func mkzcgo(dir, file string) {
	var buf strings.Builder
	fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "package build\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "const defaultCGO_ENABLED = %s\n", quote(os.Getenv("CGO_ENABLED")))

	writefile(buf.String(), file, writeSkipSame)
}

// mktzdata src/time/tzdata/zzipdata.go:
//
//	package tzdata
//	const zipdata = "PK..."
func mktzdata(dir, file string) {
	zip := readfile(filepath.Join(dir, "../../../lib/time/zoneinfo.zip"))

	var buf strings.Builder
	fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "package tzdata\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "const zipdata = %s\n", quote(zip))

	writefile(buf.String(), file, writeSkipSame)
}

// quote is like strconv.Quote but simpler and has output
// that does not depend on the exact Go bootstrap version.
func quote(s string) string {
	const hex = "0123456789abcdef"
	var out strings.Builder
	out.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 0x20 <= c && c <= 0x7E && c != '"' && c != '\\' {
			out.WriteByte(c)
		} else {
			out.WriteByte('\\')
			out.WriteByte('x')
			out.WriteByte(hex[c>>4])
			out.WriteByte(hex[c&0xf])
		}
	}
	out.WriteByte('"')
	return out.String()
}
