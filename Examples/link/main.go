package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Link struct {
	Text string
	Href string
}

func main() {
	flagHtmlFIleName := flag.String("html", "ex.html", "The path to Html file to parse")
	flag.Parse()
	f, err := os.Open(*flagHtmlFIleName)
	if err != nil {
		log.Fatalf("Failed to open file %q: %v", *flagHtmlFIleName, err)
	}
	defer f.Close()
	root, err := html.Parse(f)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}
	as := make(chan *html.Node)
	go findAnchors(root, as)
	for n := range as {
		fmt.Println(Link{
			Text: extractText(n),
			Href: extractHref(n),
		})
	}
}

func findAnchors(n *html.Node, as chan *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		as <- n
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchors(c, as)
	}
	if n.Parent == nil {
		close(as)
	}
}
func extractHref(a *html.Node) string {
	for _, attr := range a.Attr {
		if attr.Key != "href" {
			continue
		}
		return attr.Val
	}
	return ""
}

func extractText(a *html.Node) string {
	var text string
	for c := a.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
			continue
		}
		text += c.Data
	}
	return strings.TrimSpace(text)
}
