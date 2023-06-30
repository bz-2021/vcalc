/*
Copyright © 2023 bz2021 <bz2021@mail.sdu.edu.cn>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// durationCmd represents the duration command
var durationCmd = &cobra.Command{
	Use:     "total",
	Aliases: []string{"t"}, // 命令行的简写
	Short:   "Calculate total duration.",
	Long: `Calculate the total duration 
of all videos in the specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("total called")
	},
}

func init() {
	rootCmd.AddCommand(durationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// durationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// durationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
