package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"howett.net/plist" // go get howett.net/plist
)

type AppInfo struct {
	BindleID string
	Name     string
	Version  string
	BundleID string
	Path     string
}

func main() {
	appDirs := []string{"/Applications", filepath.Join(os.Getenv("HOME"), "Applications")}

	var apps []AppInfo

	for _, dir := range appDirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if strings.HasSuffix(info.Name(), ".app") && info.IsDir() {
				app := parseAppBundle(path)
				if app != nil {
					apps = append(apps, *app)
				}
				// 不递归进入 .app 目录
				return filepath.SkipDir
			}
			return nil
		})
	}

	// 输出
	for _, a := range apps {
		fmt.Printf("%s (%s)[%s] - %s\n", a.Name, a.Version, a.BindleID, a.Path)
	}
}

// 解析 Info.plist
func parseAppBundle(appPath string) *AppInfo {
	plistPath := filepath.Join(appPath, "Contents", "Info.plist")
	data, err := os.ReadFile(plistPath)
	if err != nil {
		return nil
	}

	var info map[string]interface{}
	_, err = plist.Unmarshal(data, &info)
	if err != nil {
		return nil
	}

	bindleID, _ := info["CFBundleIdentifier"].(string)
	name, _ := info["CFBundleName"].(string)
	version, _ := info["CFBundleShortVersionString"].(string)
	bundleID, _ := info["CFBundleIdentifier"].(string)

	return &AppInfo{
		BindleID: bindleID,
		Name:     name,
		Version:  version,
		BundleID: bundleID,
		Path:     appPath,
	}
}
