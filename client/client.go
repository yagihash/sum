package client

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL    string
	Md5sum string
}

func NewClient(url string) (*Client, error) {
	c := &Client{
		URL: url,
	}
	_, _, err := c.Fetch()
	return c, err
}

func (c *Client) Fetch() (string, bool, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false, err
	}

	tmpHash := c.Md5sum
	c.Md5sum = c.md5sum(string(body))
	isUpdated := tmpHash != c.Md5sum

	return string(body), isUpdated, err
}

func (c *Client) md5sum(s string) string {
	sum := md5.Sum([]byte(s))
	c.Md5sum = hex.EncodeToString(sum[:])
	return c.Md5sum
}
