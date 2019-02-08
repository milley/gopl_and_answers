package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and return the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					images++
				}
			}
		}
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWords, childImages := countWordsAndImages(c)
		words += childWords
		images += childImages
	}
	return words, images
}

func main() {
	url := "https://golang.org"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	fmt.Println("Parsing", url)
	words, images, _ := CountWordsAndImages(url)
	fmt.Println("Words:", words, "images:", images)
}
