package slack

import s "github.com/nlopes/slack"

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

func (c *Slack) NotifyUpdate(url string, hash string) error {
	attachment := s.Attachment{
		Fields: []s.AttachmentField{
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

	params := s.NewPostMessageParameters()
	params.IconEmoji = c.IconEmoji
	params.Username = c.Username

	opt := []s.MsgOption{
		s.MsgOptionText("Site Updated", false),
		s.MsgOptionAttachments(attachment),
		s.MsgOptionPostMessageParameters(params),
	}

	_, _, err := c.Api.PostMessage(c.ChannelID, opt...)
	return err
}

func (c *Slack) NotifyStart(url string, hash string) error {
	attachment := s.Attachment{
		Fields: []s.AttachmentField{
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

	params := s.NewPostMessageParameters()
	params.IconEmoji = c.IconEmoji
	params.Username = c.Username

	opt := []s.MsgOption{
		s.MsgOptionText("Site Monitoring Started", false),
		s.MsgOptionAttachments(attachment),
		s.MsgOptionPostMessageParameters(params),
	}

	_, _, err := c.Api.PostMessage(c.ChannelID, opt...)
	return err
}
