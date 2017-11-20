package toposort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/vbogretsov/go-toposort"
)

type TestCase struct {
	graph  map[string][]string
	order  []interface{}
	vertex string
	err    error
}

func TestToposort(t *testing.T) {
	cases := []TestCase{
		{
			graph: map[string][]string{
				"f3": []string{"f2", "f1"},
				"f2": []string{"f1"},
				"f1": []string{"f0"},
				"f0": []string{},
			},
			vertex: "f3",
			order:  []interface{}{"f0", "f1", "f2", "f3"},
			err:    nil,
		},
		{
			graph: map[string][]string{
				"f3": []string{"f1"},
				"f1": []string{"f0"},
				"f0": []string{"f3"},
			},
			vertex: "f3",
			err:    fmt.Errorf("cycle detected: %v", []interface{}{"f3", "f1", "f0", "f3"}),
		},
		{
			graph: map[string][]string{
				"f3": []string{"f2", "f1"},
				"f2": []string{"f1"},
				"f1": []string{"f0"},
				"f0": []string{},
			},
			vertex: "f0",
			order:  []interface{}{"f0"},
			err:    nil,
		},
		{
			graph: map[string][]string{
				"f3": []string{"f2", "f1"},
				"f2": []string{"f1"},
				"f1": []string{"f0"},
				"f0": []string{},
			},
			vertex: "f4",
			order:  []interface{}{},
			err:    fmt.Errorf("not found %v", "f4"),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("toposort#%v", i), func(t *testing.T) {
			graph := toposort.NewGraph()

			for v, ns := range tc.graph {
				graph.AddNode(v)
				for _, n := range ns {
					graph.AddNode(n)
					graph.AddEdge(v, n)
				}
			}

			order, err := graph.Resolve(tc.vertex)
			if err != nil {
				if tc.err == nil {
					t.Errorf("unexpected error: %v", err)
					t.FailNow()
				}
				if err.Error() != tc.err.Error() {
					t.Errorf("expected %v but %v", tc.err, err)
					t.FailNow()
				}
			} else if !reflect.DeepEqual(tc.order, order) {
				t.Logf("order: %v", order)
				t.Error(pretty.Diff(tc.order, order))
			}
		})
	}
}
