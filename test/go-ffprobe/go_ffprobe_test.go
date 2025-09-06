package go_ffprobe

import (
	"context"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"testing"
	"time"
)

func TestAcf(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, "/Users/toquery/Projects/Example/example-go/test/test-video.mp4")
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	log.Printf("%+v", data)
}
