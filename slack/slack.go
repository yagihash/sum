package slack

import (
	"fmt"

	s "github.com/nlopes/slack"
	"github.com/yagihashoo/sum/client"
)

type Slack struct {
	Username  string
	IconEmoji string
	Token     string
	ChannelID string
	Api       s.Client
}

func NewSlack(username string, iconEmoji string, channelID string, token string) *Slack {
	return &Slack{
		Username:  username,
		IconEmoji: iconEmoji,
		ChannelID: channelID,
		Api:       *s.New(token),
	}
}

func (c *Slack) NotifyUpdate(clients ...client.Client) error {
	params := s.NewPostMessageParameters()
	params.IconEmoji = c.IconEmoji
	params.Username = c.Username

	baseOpt := []s.MsgOption{
		s.MsgOptionPostMessageParameters(params),
	}

	for _, cl := range clients {
		opt := append(baseOpt, s.MsgOptionAttachments(s.Attachment{
			Fields: []s.AttachmentField{
				{
					Title: "URL",
					Value: cl.URL,
					Short: true,
				},
				{
					Title: "Hash",
					Value: cl.Md5sum,
					Short: true,
				},
			},
			Color: "#DF1111",
			Title: fmt.Sprintf("Detected update on %s", cl.URL),
		}))
		if _, _, err := c.Api.PostMessage(c.ChannelID, opt...); err != nil {
			return err
		}
	}
	return nil
}

func (c *Slack) NotifyStart(clients ...client.Client) error {
	params := s.NewPostMessageParameters()
	params.IconEmoji = c.IconEmoji
	params.Username = c.Username

	baseOpt := []s.MsgOption{
		s.MsgOptionPostMessageParameters(params),
	}

	for _, cl := range clients {
		opt := append(baseOpt, s.MsgOptionAttachments(s.Attachment{
			Fields: []s.AttachmentField{
				{
					Title: "URL",
					Value: cl.URL,
					Short: true,
				},
				{
					Title: "Hash",
					Value: cl.Md5sum,
					Short: true,
				},
			},
			Color: "#748931",
			Title: fmt.Sprintf("Started watching on %s", cl.URL),
		}))
		if _, _, err := c.Api.PostMessage(c.ChannelID, opt...); err != nil {
			return err
		}
	}
	return nil
}
