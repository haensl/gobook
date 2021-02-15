package main

import (
	"fmt"
	"haensl/gobook/ch5/nodes"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	ct := res.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		res.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	nodes.ForEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	title(os.Args[:1])
}
