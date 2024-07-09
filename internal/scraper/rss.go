package scraper

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title string    `xml:"title"`
		Link  string    `xml:"link"`
		Desc  string    `xml:"description"`
		Lang  string    `xml:"language"`
		Items []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// FetchFeed fetches an RSS feed from the given URL and returns the parsed feed.
func FetchFeed(url string) (RSSFeed, error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	var feed RSSFeed
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	if err != nil {
		return feed, fmt.Errorf("error parsing feed: %v", err)
	}
	return feed, nil
}
