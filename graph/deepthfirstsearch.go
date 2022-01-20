package graph

type DeepthFirstSearch struct {
	marked []bool
	count  int

	start  int // 起点
	edgeTo []int
}

func NewDeepthFirstSearch(g *Graph, s int) *DeepthFirstSearch {
	dfs := &DeepthFirstSearch{}
	dfs.marked = make([]bool, g.V())
	dfs.edgeTo = make([]int, g.V())
	dfs.start = s
	dfs.DFS(g, s)
	return dfs
}

func (dfs *DeepthFirstSearch) DFS(g *Graph, v int) {
	dfs.marked[v] = true
	dfs.count++
	for w := range g.edges[v] {
		if !dfs.marked[w] {
			dfs.edgeTo[w] = v
			dfs.DFS(g, w)
		}
	}
}

func (dfs *DeepthFirstSearch) Marked(w int) bool {
	return dfs.marked[w]
}

func (dfs *DeepthFirstSearch) Count() int {
	return dfs.count
}

func (dfs *DeepthFirstSearch) HasPathTo(v int) bool {
	return dfs.marked[v]
}

func (dfs *DeepthFirstSearch) PathTo(v int) []int {
	if !dfs.HasPathTo(v) {
		return nil
	}
	path := make([]int, 0)
	for x := v; x != dfs.start; x = dfs.edgeTo[x] {
		path = append([]int{x}, path...)
	}
	path = append([]int{dfs.start}, path...)
	return path
}
