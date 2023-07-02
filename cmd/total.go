/*
Copyright © 2023 bz2021 <bz2021@mail.sdu.edu.cn>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// totalCmd represents the total command
var totalCmd = &cobra.Command{
	Use:     "total",
	Aliases: []string{"t"},
	Short:   "Calculate and print the total duration.",
	Long: `Calculate and print the total duration 
of all videos in the specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 没有提供 -d，默认为当前路径
		if filesPath == "" {
			wd, err := os.Getwd()
			if err != nil {
				Error(cmd, args, err)
				return
			}
			filesPath = wd
		}
		// 获取 []DirEntry 数组
		dirArray, err := os.ReadDir(filesPath)
		if err != nil {
			Error(cmd, args, err)
			return
		}
		// 计算总时长
		duration := getTotalDuration(cmd, args, filesPath, &dirArray)
		// 打印总时长
		printReformat(&duration)
	},
}

func init() {
	rootCmd.AddCommand(totalCmd)
}
