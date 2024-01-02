package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func main() {
	var rootCmd = &cobra.Command{Use: "goalarm"}
	
	var setCmd = &cobra.Command{
			Use:   "set [time]",
			Short: "Set an alarm",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
					setTime(args[0])
			},
	}


	rootCmd.AddCommand(setCmd)
	rootCmd.Execute()
}

func setTime(timeStr string) {
	layout := "15:04"
	now := time.Now()

	targetTime, err := time.ParseInLocation(layout, timeStr, now.Location())
	if err != nil {
		fmt.Println("Invalid time format. Please use HH:MM.")
		return
	}

	targetTime = time.Date(now.Year(), now.Month(), now.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, now.Location())

	if targetTime.Before(now) {
		targetTime = targetTime.Add(24 * time.Hour)
	}

	diff := targetTime.Sub(now)
	fmt.Println("Alarm set for", targetTime.Format(layout))

	time.Sleep(diff)

	fmt.Println("Alarm! The time is now", targetTime.Format(layout))
	fmt.Println("Press 'Enter' to stop the alarm.")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		if char == '\r' { // Enter pressed
			break
		}
	}
}
