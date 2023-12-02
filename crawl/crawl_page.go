package crawl

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Extract(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response returned with error %v, url - %s \n", res.Status, url)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error with url: %s parsing html: %v \n", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, errUrl := res.Request.URL.Parse(a.Val)
				if errUrl != nil {
					fmt.Printf("Failed link: %v \n", errUrl)
					continue
				}
				fmt.Printf("Link parsed %s \n", link)
				links = append(links, link.String())
			}
		}
	}
	forEachHtmlNode(doc, visitNode, nil)
	return links, nil
}

func forEachHtmlNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachHtmlNode(c, pre, post)
	}
	if post != nil {
		post(node)
	}
}
