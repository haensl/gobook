package main

import (
	"fmt"
	"haensl/gobook/ch5/nodes"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	nodes.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetching %s: %v\n", url, err)
			continue
		}
		if res.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "response %s: %d\n", url, res.StatusCode)
			res.Body.Close()
			continue
		}
		html, err := html.Parse(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing %s: %v\n", url, err)
			res.Body.Close()
			continue
		}
		res.Body.Close()
		title, err := soleTitle(html)
		if err != nil {
			fmt.Fprintf(os.Stderr, "finding title %s: %v", url, err)
			continue
		}
		fmt.Printf("%s\t%s\n", url, title)
	}
}
