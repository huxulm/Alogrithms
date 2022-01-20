package graph

import (
	"reflect"
	"testing"
)

func TestDeepthFirstSearch(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		vertices int
		paths    []struct {
			to   int
			path []int
		}
	}{
		{
			name: "tiny graph",
			edges: [][]int{
				{0, 2},
				{0, 1},
				{0, 5},
				{1, 0},
				{1, 2},
				{2, 0},
				{2, 1},
				{2, 3},
				{2, 4},
				{3, 5},
				{3, 4},
				{3, 2},
				{4, 3},
				{4, 2},
				{5, 3},
				{5, 0},
			},
			vertices: 6,
			paths: []struct {
				to   int
				path []int
			}{
				{
					to:   0,
					path: []int{0},
				},
				{
					to:   1,
					path: []int{0, 2, 1},
				},
				{
					to:   2,
					path: []int{0, 2},
				},
				{
					to:   3,
					path: []int{0, 2, 3},
				},
				{
					to:   4,
					path: []int{0, 2, 3, 4},
				},
				{
					to:   5,
					path: []int{0, 2, 3, 5},
				},
			},
		},
	}

	for _, test := range tests {
		// undirect graph
		graph := New(test.vertices)
		for _, edge := range test.edges {
			graph.AddEdge(edge[0], edge[1])
		}
		dfs := NewDeepthFirstSearch(graph, 0)
		// dfs.DFS(graph, 0)
		if count := dfs.Count(); count != graph.V() {
			t.Errorf("Expect %d but got %d", graph.V(), count)
		}
		for _, p := range test.paths {
			if path := dfs.PathTo(p.to); !reflect.DeepEqual(path, p.path) {
				t.Errorf("Path from %d to %d should be: %v got: %v", dfs.start, p.to, p.path, path)
			}
		}
	}
}
