/* Secondary message attachments
Ref: https://api.slack.com/reference/messaging/attachments
*/
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
		Pretext:    "This is slack post test by Go",
		Title:      "title",
		Color:      "#36a64f",
		AuthorName: "author_name",
		AuthorIcon: "https://placeimg.com/16/16/people",
		MarkdownIn: []string{"`textTomarkdown`"},
		Text:       "hello world `textTomarkdown`",
		ThumbURL:   "http://placekitten.com/g/200/200",
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",

		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Item1",
				Value: "this is value of item1",
				Short: false,
			}, slack.AttachmentField{
				Title: "Item2",
				Value: "this is value of item2",
				Short: true,
			}, slack.AttachmentField{
				Title: "Item3",
				Value: "```" + "this is value of item3" + "```",
				Short: false,
			},
		},
	}

	channelID, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("This is Title", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
