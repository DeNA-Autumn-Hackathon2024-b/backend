package infra

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

func (i *Infrastructure) ConvertVideoHLS(ctx context.Context, songID string, s3URL string) error {
	cmd := exec.Command("ffmpeg",
		"-i", s3URL,
		"-vn", "-ac", "2", "-acodec", "aac",
		"-f", "segment", "-segment_format", "mpegts", "-segment_time", "10",
		"-segment_list", "output/"+songID+".m3u8",
		"output/"+songID+"%05d.ts",
		"-y",
	)

	log.Println(cmd.Args)
	result, err := cmd.CombinedOutput()

	log.Println(string(result))
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	return nil
}
