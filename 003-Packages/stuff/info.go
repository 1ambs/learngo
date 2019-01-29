package stuff

import (
	"fmt"
	"runtime"
)

func GoInfo() {
	fmt.Printf("runtime: %s\narchitecture: %s", runtime.GOOS, runtime.GOARCH)
}
