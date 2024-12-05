package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
)

func main() {
	// removeVideo := true
	// location?
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("\n████████████████████████████████████████████████████████████████████████████████████████████████████████████\n█▄─█─▄█─▄▄─█▄─██─▄█─▄─▄─█▄─██─▄█▄─▄─▀█▄─▄▄─█▄─▄▄▀█─▄▄─█▄─█▀▀▀█─▄█▄─▀█▄─▄█▄─▄███─▄▄─██▀▄─██▄─▄▄▀█▄─▄▄─█▄─▄▄▀█\n██▄─▄██─██─██─██─████─████─██─███─▄─▀██─▄█▀██─██─█─██─██─█─█─█─███─█▄▀─███─██▀█─██─██─▀─███─██─██─▄█▀██─▄─▄█\n▀▀▄▄▄▀▀▄▄▄▄▀▀▄▄▄▄▀▀▀▄▄▄▀▀▀▄▄▄▄▀▀▄▄▄▄▀▀▄▄▄▄▄▀▄▄▄▄▀▀▄▄▄▄▀▀▄▄▄▀▄▄▄▀▀▄▄▄▀▀▄▄▀▄▄▄▄▄▀▄▄▄▄▀▄▄▀▄▄▀▄▄▄▄▀▀▄▄▄▄▄▀▄▄▀▄▄▀\n")

		client := youtube.Client{}
		fmt.Println("Please paste and enter the video ID")
		videoID, _ := reader.ReadString('\n')
		videoID = strings.TrimSpace(videoID)
		fmt.Println("Put in the title for your file")
		videoTitle, _ := reader.ReadString('\n')
		videoTitle = strings.TrimSpace(videoTitle)
		if videoID == "" {
			videoID = "dQw4w9WgXcQ"
		}
		if videoTitle == "" {
			videoTitle = "Rick Astley - Never Gonna Give You Up"
		}

		fmt.Println("Extracting video...")

		video, err := client.GetVideo(videoID)
		if err != nil {
			fmt.Println("Wrong URL / video ID.\nPaste either full link, e.g.: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', or the video ID, e.g.: 'dQw4w9WgXcQ'")
			time.Sleep(3 * time.Second)
			continue
		}

		formats := video.Formats.WithAudioChannels()
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			fmt.Print(err)
			continue
		}
		defer stream.Close()

		file, err := os.Create("video.mp4")
		if err != nil {
			fmt.Print(err)
			continue
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Println("Video extracted successfully")

		fmt.Println("Extracting audio...")
		cmd := exec.Command("ffmpeg", "-i", "video.mp4", "-vn", "-ab", "128k", "-ar", "44100", "-y", videoTitle+".mp3")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error extracting audio:", err)
			return
		}

		fmt.Println("Audio extracted successfully")

		// add: if statement to check removeVideo

		time.Sleep(1 * time.Second)

	}
}
