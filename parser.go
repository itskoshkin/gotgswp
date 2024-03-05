package gotgswp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func ParseChannelInfo(url string) (Channel, error) {
	body, err := Get(url)
	if err != nil {
		return Channel{}, err
	}
	channel, err := ParseChannel(body)
	if err != nil {
		return Channel{}, err
	}
	return channel, nil
}

func FindLastPost(html string) (int, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return -1, fmt.Errorf("error loading HTML: %v", err)
	}
	var lastPostID int
	var lastPostDiv *goquery.Selection

	doc.Find(".tgme_widget_message").Each(func(i int, s *goquery.Selection) {
		dataPost, exists := s.Attr("data-post")
		if exists {
			parts := strings.Split(dataPost, "/")
			if len(parts) > 1 {
				postID, err := strconv.Atoi(parts[1])
				if err == nil {
					if postID > lastPostID {
						lastPostID = postID
						lastPostDiv = s
					}
				}
			}
		}
	})

	if lastPostDiv != nil {
		return lastPostID, nil
	} else {
		return -1, fmt.Errorf("last post not found")
	}
}

func ParseSpecificMessage(url string, post int) (string, error) {
	body, err := Get(url)
	if err != nil {
		return "", err
	}
	message, err := ParseMessage(body, post)
	if err != nil {
		return "", err
	}
	return message, nil
}
