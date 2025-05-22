package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestHttpDiskFile(t *testing.T) {

	// 设置服务器监听的端口
	port := "8080"

	// 注册处理磁盘文件的路由
	http.HandleFunc("/disk_file/", handleDiskFile)

	// 启动HTTP服务器
	fmt.Printf("服务器已启动，监听端口: %s\n", port)
	fmt.Printf("访问文件示例: http://localhost:%s/disk_file/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// 处理磁盘文件请求的函数
func handleDiskFile(w http.ResponseWriter, r *http.Request) {
	// 从URL中提取文件路径
	filePath := strings.TrimPrefix(r.URL.Path, "/disk_file")

	// 安全检查：防止访问上级目录
	if strings.Contains(filePath, "../") || strings.Contains(filePath, "..\\") {
		http.Error(w, "禁止访问上级目录", http.StatusForbidden)
		return
	}

	// 检查文件是否存在
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "文件不存在", http.StatusNotFound)
		} else {
			http.Error(w, "无法访问文件", http.StatusInternalServerError)
		}
		return
	}

	// 如果是目录，返回错误
	if fileInfo.IsDir() {
		http.Error(w, "不能访问目录", http.StatusForbidden)
		return
	}

	// 提供文件下载
	http.ServeFile(w, r, filePath)
}
