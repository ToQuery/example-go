package go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var s2 string

func TestGEString(t *testing.T) {
	fmt.Println(s)
	fmt.Println(s2)
}

//go:embed hello.txt
var b []byte

func TestGEByte(t *testing.T) {
	fmt.Println(b)
}

//go:embed assets
var gea embed.FS

//go:embed all:assets
var geaa embed.FS

//go:embed assets/*
var gea_ embed.FS // 只包含 assets/ 下的文件，不递归

//go:embed all:assets/*
var geaa_ embed.FS // 包含整个 assets/ 目录及子目录、文件

func TestGEReadDir(t *testing.T) {
	fmt.Printf("\n\n[go:embed assets]\n")
	PrintDirAndSubDir(gea)

	fmt.Printf("\n\n[go:embed all:assets]\n")
	PrintDirAndSubDir(geaa)

	fmt.Printf("\n\n[go:embed assets/*]\n")
	PrintDirAndSubDir(gea_)

	fmt.Printf("\n\n[go:embed all:assets/*]\n")
	PrintDirAndSubDir(geaa_)
}

func PrintDirAndSubDir(assets embed.FS) {
	dirEntries, _ := assets.ReadDir("assets")
	for _, de := range dirEntries {
		fmt.Printf("assets/%s IsDir=%t \n", de.Name(), de.IsDir())
	}
	publicDirEntries, _ := assets.ReadDir("assets/public")
	for _, de := range publicDirEntries {
		fmt.Printf("assets/public/%s IsDir=%t \n", de.Name(), de.IsDir())
	}
}
