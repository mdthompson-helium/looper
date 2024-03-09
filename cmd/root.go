/*
Copyright © 2024 matt@heliumdev.com
*/
package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "looper",
	Short: "An annoying prank to play on your friends/coworkers",
	Long: `looper is a simple command line tool that plays an annoying sound at a regular interval.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		loop()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func loop() {
	for {
		cmd := exec.Command("say", "-v", "grandpa", "-r", "10", "'chicken'")
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}