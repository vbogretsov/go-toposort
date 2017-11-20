package toposort

import (
	"fmt"
)

type Graph struct {
	graph map[interface{}]map[interface{}]bool
}

func NewGraph() *Graph {
	return &Graph{map[interface{}]map[interface{}]bool{}}
}

func (g *Graph) AddEdge(from interface{}, to interface{}) {
	g.graph[from][to] = true
}

func (g *Graph) AddNode(v interface{}) {
	_, ok := g.graph[v]
	if !ok {
		g.graph[v] = map[interface{}]bool{}
	}
}

func (g *Graph) Resolve(vertex interface{}) ([]interface{}, error) {
	if _, ok := g.graph[vertex]; !ok {
		return nil, fmt.Errorf("not found %v", vertex)
	}

	order := []interface{}{}
	visit := map[interface{}]bool{}

	var dfs func(interface{}) []interface{}

	dfs = func(v interface{}) []interface{} {
		visited, ok := visit[v]
		if !ok {
			visit[v] = false

			for w, _ := range g.graph[v] {
				cycle := dfs(w)
				if cycle != nil {
					cycle = append([]interface{}{v}, cycle...)
					return cycle
				}
			}

			visit[v] = true
			order = append(order, v)
		} else if !visited {
			return []interface{}{v}
		}

		return nil
	}

	cycle := dfs(vertex)
	if cycle != nil {
		return nil, fmt.Errorf("cycle detected: %v", cycle)
	}

	return order, nil
}
