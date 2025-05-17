package main

import (
	"fmt"
	"runtime"
)

var (
	Version   = "dev" // 默认值，会被构建时覆盖
	BuildId   = "0x0001"
	BuildTime = "unknown"
)

func main() {
	osType := runtime.GOOS
	osArch := runtime.GOARCH

	fmt.Printf("example-wails version %s (BuildId: %s) %s_%s built at %s\n", Version, BuildId, osType, osArch, BuildTime)
}
