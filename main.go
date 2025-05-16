package main

import (
	"embed"
	"fmt"
	"runtime"
)

//go:embed assets
var fs embed.FS

var (
	Version   = "dev" // 默认值，会被构建时覆盖
	BuildTime = "unknown"
)

func main() {
	osType := runtime.GOOS
	osArch := runtime.GOARCH

	fmt.Printf("example-go %s (%s_%s) built at %s\n", Version, osType, osArch, BuildTime)
}
