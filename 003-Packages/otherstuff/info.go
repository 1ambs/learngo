package stuff

import (
	"fmt"
	"runtime"
)

func Dave() {
	fmt.Printf("runtime: %s\narchitecture: %s", runtime.GOOS, runtime.GOARCH)
}
