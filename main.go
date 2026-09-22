package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	if botToken == "" || channelID == "" {
		fmt.Println("Error: Please set SLACK_BOT_TOKEN and CHANNEL_ID environment variables.")
		return
	}

	api := slack.New(botToken)
	channelArr := []string{channelID}
	fileArr := []string{"git-cheat-sheet.pdf"}

	for i := 0; i < len(fileArr); i++ {
		fileInfo, err := os.Stat(fileArr[i])
		if err != nil {
			fmt.Printf("File error: %s\n", err.Error())
			return
		}

		params := slack.UploadFileParameters{
			Channels: channelArr,
			File:     fileArr[i],
			Filename: fileInfo.Name(),
			FileSize: int(fileInfo.Size()),
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error : %s\n", err.Error())
			return
		}

		fmt.Printf("ID : %s, Title : %s\n", file.ID, file.Title)
	}
	
}	