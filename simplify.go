package gotgswp

import (
	"regexp"
	"strings"
)

func Simplify(html string) string {
	html = regexp.MustCompile(`(?i)<a href="[^"]*">([^<]*)</a>`).ReplaceAllString(html, "$1")
	html = regexp.MustCompile(`<i class="emoji"[^>]*>.*?</i>`).ReplaceAllString(html, "🐱")
	html = regexp.MustCompile(`(?s)<span[^>]*>(.*?)</span>`).ReplaceAllString(html, "$1")
	html = regexp.MustCompile(`(?s)<a\s+[^>]*>(.*?)</a>`).ReplaceAllString(html, "$1")
	html = regexp.MustCompile(`(?i)<emoji[^>]*>`).ReplaceAllString(html, "🐱")
	html = regexp.MustCompile(`(?i)<br\s*/?>`).ReplaceAllString(html, "\n")
	html = regexp.MustCompile(`(?i)</?b>`).ReplaceAllString(html, "")
	html = strings.ReplaceAll(html, "&quot;", "\"")
	return strings.ReplaceAll(html, "&#39;", "'")
}
