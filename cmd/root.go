/*
Copyright © 2023 bz2021 <bz2021@mail.sdu.edu.cn>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var filesPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vcalc",
	Short: "Calculate the total duration of video files",
	Long: `                                __           
                               /  |          
 __     __   _______   ______  $$ |  _______ 
/  \   /  | /       | /      \ $$ | /       |
$$  \ /$$/ /$$$$$$$/  $$$$$$  |$$ |/$$$$$$$/ 
 $$  /$$/  $$ |       /    $$ |$$ |$$ |      
  $$ $$/   $$ \_____ /$$$$$$$ |$$ |$$ \_____ 
   $$$/    $$       |$$    $$ |$$ |$$       |
    $/      $$$$$$$/  $$$$$$$/ $$/  $$$$$$$/ 

Vcalc is a simple CLI tool written by Go
allows you to easily calculate the total duration
of video files within a specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, ErrUnrecognizedSubcommand)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filesPath, "directory", "d", "./", "source directory to read from")
}
