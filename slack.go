package main

import (
	"github.com/nlopes/slack"
)

type Slack struct {
	Username  string
	IconEmoji string
	Token     string
	ChannelID string
	Api       slack.Client
}

func newSlack(username string, iconEmoji string, channelID string, token string) *Slack {
	return &Slack{
		Username:  username,
		IconEmoji: iconEmoji,
		ChannelID: channelID,
		Api:       *slack.New(token),
	}
}

func (c *Slack) Notify(url string, hash string) error {
	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			{
				Title: "URL",
				Value: url,
			},
			{
				Title: "Hash",
				Value: hash,
			},
		},
		Color: "#748931",
	}

	params := slack.NewPostMessageParameters()
	params.IconEmoji = c.IconEmoji
	params.Username = c.Username

	opt := []slack.MsgOption{
		slack.MsgOptionText("Site Updated", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionPostMessageParameters(params),
	}

	_, _, err := c.Api.PostMessage(c.ChannelID, opt...)
	return err
}
