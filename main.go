package main

import (
	"log"
	"os"
	"runtime"

	"github.com/yagihashoo/sum/client"

	"github.com/carlescere/scheduler"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("API token for Slack is required")
	}

	channelID := os.Getenv("SLACK_CHANNEL")
	if channelID == "" {
		log.Fatal("Channel ID for Slack is required")
	}

	username := os.Getenv("SLACK_NAME")
	if username == "" {
		username = "sum"
	}

	iconEmoji := os.Getenv("SLACK_EMOJI")
	if iconEmoji == "" {
		iconEmoji = ":robot_face:"
	}

	url := os.Getenv("URL")
	if url == "" {
		url = "http://www.baseballbear.com/live/"
	}

	slack := newSlack(username, iconEmoji, channelID, token)

	c, err := client.NewClient(url)
	if err != nil {
		log.Fatal(err)
	}

	job := func() {
		_, isUpdated, err := c.Fetch()
		if err != nil {
			log.Fatal(err)
		}

		if isUpdated {
			err := slack.Notify(c.URL, c.Md5sum)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	_, err = scheduler.Every(5).Minutes().Run(job)
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}
