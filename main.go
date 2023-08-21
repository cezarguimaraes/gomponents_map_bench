package main

import (
	"os"

	g "github.com/maragudk/gomponents"
)

type BigStruct struct {
	vals [1000]int
}

type BigStruct10x struct {
	vals [10000]int
}

func MapRefPreAlloc[T any](ts []T, cb func(*T) g.Node) []g.Node {
	nodes := make([]g.Node, 0, len(ts))
	for i := range ts {
		nodes = append(nodes, cb(&ts[i]))
	}
	return nodes
}

func MapRef[T any](ts []T, cb func(*T) g.Node) []g.Node {
	var nodes []g.Node
	for i := range ts {
		nodes = append(nodes, cb(&ts[i]))
	}
	return nodes
}

func MapPreAlloc[T any](ts []T, cb func(T) g.Node) []g.Node {
	nodes := make([]g.Node, 0, len(ts))
	for _, t := range ts {
		nodes = append(nodes, cb(t))
	}
	return nodes
}

func render[T any](b T) g.Node {
	return g.El("a")
}

func main() {
	var foo BigStruct
	render(foo).Render(os.Stdout)
}
