package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	res, httpErr := httpClient.Get(url)
	if httpErr != nil {
		return RSSFeed{}, httpErr
	}
	defer res.Body.Close()

	data, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return RSSFeed{}, readErr
	}
	rssFeed := RSSFeed{}
	transformErr := xml.Unmarshal(data, &rssFeed)
	if transformErr != nil {
		return RSSFeed{}, transformErr
	}
	return rssFeed, nil
}
