package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

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

	body, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	hash := md5sum(body)

	job := func() {
		body, err := fetch(url)
		if err != nil {
			log.Fatal(err)
		}

		current := md5sum(body)
		if current != hash {
			hash = current
			err := slack.Notify(url, hash)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	_, err = scheduler.Every(10).Seconds().Run(job)
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()

}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func md5sum(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
