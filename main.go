package main

import (
	"fmt" // formatting package

	"os" // Accessing stdinput

	"github.com/joho/godotenv"  // for accessing .env variables
	"github.com/slack-go/slack" //github.com slack library
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(os.Getenv("SLACK_CHANNEL"), slack.MsgOptionText("Efi rules", false))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Message has been sent %s %s\n", channelID, timestamp)
}
