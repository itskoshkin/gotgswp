package gotgswp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ParseChannel(html string) (Channel, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return Channel{}, fmt.Errorf("error loading HTML: %v", err)
	}

	title, _ := doc.Find("meta[property='og:title']").Attr("content")
	description, _ := doc.Find("meta[property='og:description']").Attr("content")
	var count string
	doc.Find(".tgme_channel_info_counter").Each(func(i int, s *goquery.Selection) {
		if s.Find(".counter_type").Text() == "subscribers" {
			count = s.Find(".counter_value").Text()
		}
	})

	return Channel{
		Title:        title,
		Description:  description,
		MembersCount: count,
		Link:         "",
	}, nil
}

func ParseMessage(html string, post int) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", fmt.Errorf("error loading HTML: %v", err)
	}

	selector := fmt.Sprintf(".tgme_widget_message[data-post*='/%d']", post)
	messageHtml, err := doc.Find(selector).Find(".tgme_widget_message_text").Html()
	if err != nil {
		return "", fmt.Errorf("error finding message: %v", err)
	}
	return messageHtml, nil
}
