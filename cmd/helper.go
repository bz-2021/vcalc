package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/vansante/go-ffprobe.v2"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Error 打印错误信息
func Error(cmd *cobra.Command, args []string, err error) {
	_, err = fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v", cmd.Name(), args, err)
	if err != nil {
		return
	}
	os.Exit(1)
}

// 递归计算总时长
func getTotalDuration(cmd *cobra.Command, args []string, filesPath string, entry *[]os.DirEntry) time.Duration {
	//var duration int
	for _, v := range *entry {
		vPath := path.Join(filesPath, v.Name())
		if isVideo(vPath) {
			timeDuration, err := getVideoDuration(vPath)
			if err != nil {
				Error(cmd, args, err)
			}
			fmt.Println(v.Name(), "duration:", timeDuration)
		}
	}
	return 0
}

// 打印格式化后的时长信息
func printReformat(duration time.Duration) {
	return
}

// 通过 ffprobe 获取单个视频的时长
func getVideoDuration(videoPath string) (time.Duration, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, videoPath)
	if err != nil {
		return 0, err
	}
	return data.Format.Duration(), nil
}

// 判断是否是视频类型（对大多数常见视频类型有效）
func isVideo(filename string) bool {
	mimeType := mime.TypeByExtension(filepath.Ext(filename))
	if strings.HasPrefix(mimeType, "video/") {
		return true
	}
	return false
}
