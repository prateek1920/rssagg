package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	log.Println("Going to fetch posts for feed : ", url)
	var rssFeed RSSFeed

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	err = xml.Unmarshal(respBytes, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
