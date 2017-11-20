# go-toposort

Topological sorting implementation in Go

## Installation

```bash
$ go get github.com/vbogretsov/go-toposort
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/vbogretsov/go-toposort"
)

func main() {
    graph := toposort.NewGraph()
    graph.AddNode("f0")
    graph.AddNode("f1")
    graph.AddNode("f2")
    graph.AddNode("f3")
    graph.AddEdge("f1", "f0")
    graph.AddEdge("f3", "f2")
    graph.AddEdge("f2", "f1")
    graph.AddEdge("f3", "f1")

    order, err := graph.Resolve("f3")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(order)
}
```

Output:

```bash
$ ./main
[f0 f1 f2 f3]
```

## Licence

See the [LICENSE](https://github.com/vbogretsov/go-toposort/blob/master/LICENSE) file.
