package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

//go:embed assets/alarm.wav
var alarmSound embed.FS

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
		fmt.Printf("Alarm set for %s (tomorrow)\n", targetTime.Format(layout))
	} else {
		fmt.Printf("Alarm set for %s (today)\n", targetTime.Format(layout))
	}

	diff := targetTime.Sub(now)
	fmt.Println("Waiting for alarm... Press Ctrl + C to cancel.")

	time.Sleep(diff)

	done := make(chan bool)
	go playAlarmSound(done)

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
			done <- true
			break
		}
	}
}

func playAlarmSound(done chan bool) {
	alarmFile, err := alarmSound.Open("assets/alarm.wav")
	if err != nil {
    log.Fatal(err)
	}
	defer alarmFile.Close()

	streamer, format, err := wav.Decode(alarmFile)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := beep.Loop(-1, streamer)

	go func() {
		select {
		case <-done:
		case <-time.After(15 * time.Minute):
			fmt.Println("15 minutes have passed, stopping the alarm automatically")
      os.Exit(0)
		}
		speaker.Clear() // Clear the speaker to stop playing
	}()

	speaker.Play(loop)
}
