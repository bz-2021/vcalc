/*
Copyright © 2023 bz2021 <bz2021@mail.sdu.edu.cn>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
)

var Level int8

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Quickly visualizing the directory structure.",
	Long:  `Quickly visualizing the directory structure and the video files it contains.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tree called")
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)

	treeCmd.Flags().Int8VarP(&Level, "level", "l", math.MaxInt8, "the depth of the files tree")
}
