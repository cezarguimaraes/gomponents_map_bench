package main

import (
	"fmt"
	"os"
	"sync"
	"unsafe"

	g "github.com/maragudk/gomponents"
)

// performance issue
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

// noCopy issue
type NoCopyElement struct {
	mu sync.Mutex
}

func main() {
	var foo BigStruct
	fmt.Println(unsafe.Sizeof(foo))
	render(foo).Render(os.Stdout)

	g.Map(make([]NoCopyElement, 10), func(_ NoCopyElement) g.Node {
		return g.El("p")
	})

	MapRef(make([]NoCopyElement, 10), func(_ *NoCopyElement) g.Node {
		return g.El("p")
	})
}
