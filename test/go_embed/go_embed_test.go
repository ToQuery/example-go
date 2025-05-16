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
var fs embed.FS

func TestGEReadDir(t *testing.T) {
	dirEntries, _ := fs.ReadDir("assets")
	for _, de := range dirEntries {
		fmt.Printf("assets/%s IsDir=%t \n", de.Name(), de.IsDir())
	}
	publicDirEntries, _ := fs.ReadDir("assets/public")
	for _, de := range publicDirEntries {
		fmt.Printf("assets/public/%s IsDir=%t \n", de.Name(), de.IsDir())
	}
}
