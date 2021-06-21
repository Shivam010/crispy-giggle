package cg

import (
	"embed"
	"fmt"
)

//go:embed license testdata
var fs embed.FS

func TestData() {
	list, err := fs.ReadDir("testdata")
	if err != nil {
		fmt.Println("error testdata", err)
		return
	}
	fmt.Println("testdata dir")
	for _, f := range list {
		fmt.Println("name", f.Name())
		fmt.Println("dir", f.IsDir())
		fmt.Println()
	}
}

func LicenseDir() {
	list, err := fs.ReadDir("license")
	if err != nil {
		fmt.Println("error license dir", err)
		return
	}
	fmt.Println("license dir")
	for _, f := range list {
		fmt.Println("name", f.Name())
		fmt.Println("dir", f.IsDir())
		fmt.Println()
	}
}

func License() {
	ff, err := fs.ReadFile("license")
	if err != nil {
		fmt.Println("error license file", err)
		return
	}
	fmt.Println("license file", string(ff))
}
