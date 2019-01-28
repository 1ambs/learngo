package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
}
