package ffmpeg

import (
	"bytes"
	"log"
	"testing"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func TestFfmpegVersion(t *testing.T) {
	// 获取ffmpeg版本信息
	cmd := ffmpeg_go.Input("/dev/null").Output("/dev/null", ffmpeg_go.KwArgs{"version": ""})
	buf := bytes.NewBuffer(nil)
	err := cmd.WithOutput(buf).WithErrorOutput(buf).Run()
	if err != nil {
		// 即使有错误也输出获取到的版本信息
		log.Printf("FFmpeg版本信息(执行出错):\n%s", buf.String())
		log.Fatalf("获取ffmpeg版本失败: %v", err)
	}
	log.Printf("FFmpeg版本信息:\n%s", buf.String())
}

func TestFfmpegProbe(t *testing.T) {
	videoPath := ""
	// Get video properties using ffprobe
	probe, err := ffmpeg_go.Probe(videoPath)
	if err != nil {
		log.Fatalf("failed to probe video: %v", err)
	}
	log.Println(probe)
}
