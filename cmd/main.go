package main

import (
	"fmt"

	_ "embed"

	cg "github.com/Shivam010/crispy-giggle/v3"
	test "github.com/Shivam010/crispy-giggle/v3/testdata"
)

//go:embed dsa/*
var cont string

func main() {
	fmt.Println("test data")
	cg.TestData()
	fmt.Println("license")
	cg.License()
	fmt.Println("license dir")
	cg.LicenseDir()

	fmt.Println()
	fmt.Println(cont)

	fmt.Println()
	fs, err := test.File.ReadDir(".")
	fmt.Println("error", err)
	for _, f := range fs {
		// io.ReadAll(f)
		fmt.Println(f.Name())
		fmt.Println(f.IsDir())
		d, err := test.File.ReadFile(f.Name())
		fmt.Println(string(d), err)
		fmt.Println()
	}
	// fmt.Println(fs[0], fs[1], fs[2], err)
}
