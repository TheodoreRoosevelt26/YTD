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
		videoIDInput, _ := reader.ReadString('\n')
		videoID := strings.TrimSpace(videoIDInput)
		fmt.Println("Put in the title for your file")
		videoTitle, _ := reader.ReadString('\n')
		fileTitle := strings.TrimSpace(videoTitle)
		if videoID == "" {
			videoID = "dQw4w9WgXcQ"
		}
		if fileTitle == "" {
			fileTitle = "Rick Astley - Never Gonna Give You Up"
		}

		fmt.Println("Extracting video...")

		video, err := client.GetVideo(videoID)
		if err != nil {
			fmt.Println("ERROR GETTING VIDEO: Wrong URL / video ID.\nPaste either full link, e.g.: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', or the video ID, e.g.: 'dQw4w9WgXcQ'")
			time.Sleep(3 * time.Second)
			continue
		}

		formats := video.Formats.WithAudioChannels()
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			fmt.Printf("ERROR GETTING VIDEO STREAM: %s", err)
			continue
		}
		defer stream.Close()

		file, err := os.Create("video.mp4")
		if err != nil {
			fmt.Printf("ERROR CREATING VIDEO FILE: %s", err)
			continue
		}
		defer file.Close()

		bytesCopied, err := io.Copy(file, stream)
		if err != nil {
			fmt.Printf("ERROR CREATING VIDEO: %s", err)
			continue
		}
		if bytesCopied == 0 {
			fmt.Println("ERROR UNABLE TO RECEIVE VIDEO STEAM: Stream copying resulted in 0 bytes.")
			continue
		}

		fmt.Println("Video extracted successfully")

		fmt.Println("Extracting audio...")
		cmd := exec.Command("ffmpeg", "-i", "video.mp4", "-vn", "-ab", "128k", "-ar", "44100", "-y", fileTitle+".mp3")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error extracting audio:", err)
			continue
		}

		fmt.Println("Audio extracted successfully")

		// add: if statement to check removeVideo

		time.Sleep(1 * time.Second)
	}
}
