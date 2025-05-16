package runtime

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestRuntimeInfo(t *testing.T) {

	// 获取当前操作系统类型
	osType := runtime.GOOS
	osArch := runtime.GOARCH

	log.Printf("OSType: %s", osType)
	log.Printf("OSArch: %s", osArch)

	envs := os.Environ()
	for i, env := range envs {
		log.Printf("[%d] %s", i, env)
	}

}

func TestExecBinary(t *testing.T) {
	envs := os.Environ()
	for _, env := range envs {
		envKV := strings.Split(env, "=")
		envKey := envKV[0]
		envVal := envKV[1]
		if envKey == "PATH" {
			err := os.Setenv("PATH", envVal+":/Users/toquery/Projects/Example/example-go/assets/binary/example_0.0.1/darwin-arm64")
			if err != nil {
				return
			}
		}
	}

	args := []string{"-version"}
	cmd := exec.CommandContext(t.Context(), "example-go", args...)

	buf := bytes.NewBuffer(nil)
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = stdErrBuf

	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Errorf("[%s] %w", string(stdErrBuf.Bytes()), err)
		fmt.Println(errMsg)
	}
	msg := string(buf.Bytes())
	fmt.Println(msg)
}
