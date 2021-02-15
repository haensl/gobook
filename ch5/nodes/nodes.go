package nodes

import "golang.org/x/net/html"

// ForEachNode executes functions pre and post for each
// node in n. pre is executed before traversal of
// children, post afterwards.
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
