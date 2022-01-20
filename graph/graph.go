package graph

type Graph struct {
	vertices int                 // Vertice counts
	edges    map[int]map[int]int // Stores weight of an edge
	Directed bool
}

func New(v int) *Graph {
	return &Graph{
		vertices: v,
	}
}

func (g *Graph) AddVertice(v int) {
	if g.edges == nil {
		g.edges = make(map[int]map[int]int)
	}
	// Check vertice present or not
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[int]int)
	}
}

// AddEdge will add a new edge between the provided vertices in the graph
func (g *Graph) AddEdge(v1, v2 int) {
	// Add an edge with 0 weight
	g.AddWeightEdge(v1, v2, 0)
}

func (g *Graph) AddWeightEdge(v1, v2, weight int) {
	// Add vertice v1
	g.AddVertice(v1)
	g.AddVertice(v2)
	// And finally add the edges
	// one->two and two->one for undirected graph
	// one->two for directed graphs
	g.edges[v1][v2] = weight
	if !g.Directed {
		g.edges[v2][v1] = weight
	}
}
