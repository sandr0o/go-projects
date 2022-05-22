package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	fmt.Printf("%s is used\n", os.Getenv("SLACK_BOT_TOKEN"))

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test2_c.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
