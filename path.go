package main

import (
	"path"
	"path/filepath"
	"fmt"
)

func main() {
	fmt.Println(path.Base("/a/a/a/b")) //b

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}


	fmt.Println(path.Dir("/a/b/c"))// /a/b
	fmt.Println(path.Dir("/a/b/c/"))// /a/b/c

	fmt.Println(path.Ext("/a/b/c"))// empty
	fmt.Println(path.Ext("/a/b/c/bar.css")) //.css

	fmt.Println(path.IsAbs("../dev/../null"))// false
	fmt.Println(path.IsAbs("/dev/null"))//true

	fmt.Println(path.Join("a","b","c")) // a/b/c
	fmt.Println(path.Join("/a","../b","//////c")) // /b/c

	fmt.Println(path.Split("static/myfile.css"))

	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
}