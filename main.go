package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var (
	SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
	CHANNEL_ID  = os.Getenv("CHANNEL_ID")
)

func main() {
	api := slack.New(SLACK_TOKEN)
	attachment := slack.Attachment{
		Title: "Title",
		Text:  "",
		Color: "good",

		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Item1",
				Value: "```" + "this is sample value" + "```",
				Short: false,
			}, slack.AttachmentField{
				Title: "Item2",
				Value: "```" + "this is sample value" + "```",
				Short: false,
			}, slack.AttachmentField{
				Title: "Item3",
				Value: "```" + "this is sample value" + "```",
				Short: false,
			},
		},
	}

	channelID, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("Title", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
