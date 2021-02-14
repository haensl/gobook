// Package graph encasulates directed graph functionality.
package graph

// Graph is a directed graph.
type Graph map[string]map[string]bool

// NewGraph returns a new Graph.
func NewGraph() Graph {
	return make(Graph)
}

// AddEdge adds an edge to Graph g
func AddEdge(g Graph, from, to string) {
	edges := g[from]
	if edges == nil {
		edges = make(map[string]bool)
		g[from] = edges
	}
	edges[to] = true
}

// HasEdge determines if Graph g contains an edge from from to to.
func HasEdge(g Graph, from, to string) bool {
	return g[from][to]
}
