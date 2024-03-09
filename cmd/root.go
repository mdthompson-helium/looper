/*
Copyright © 2024 matt@heliumdev.com
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var random bool

var rootCmd = &cobra.Command{
	Use:   "looper",
	Short: "An annoying prank to play on your friends/coworkers",
	Long: `looper is a simple command line tool that plays an annoying sound at a regular interval.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		voice := cmd.Flag("voice").Value.String()
		rate := cmd.Flag("rate").Value.String()
		speech := cmd.Flag("speech").Value.String()
		loop(voice, rate, speech)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("voice", "v", "grandpa", "What voice do you want to use?")
	rootCmd.Flags().StringP("rate", "r", "100", "How fast should they speak?")
	rootCmd.Flags().StringP("speech", "s", "chicken", "What should they say?")
	rootCmd.Flags().BoolVar(&random, "random", false, "Should the speech be random?")
}


func loop(voice string, rate string, speech string) {
	for {
		cmd := exec.Command("say", "-v", voice, "-r", rate, fmt.Sprintf("'%s'", speech))
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		sleepTime := 1 * time.Second
		if random {
			sleepTime = time.Duration(5 + (rand.Intn(1200))) * time.Second
		}
		fmt.Println("Sleeping for", sleepTime)
		time.Sleep(sleepTime)
	}
}