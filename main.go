package main

import (
	"os"
	"sync"

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

// type appoximations (~) are deliberately ignored
// so Map can be implemented without reflection
type NodeMapper[T any] interface {
	func(T) g.Node | func(*T) g.Node
}

func Map3[Y NodeMapper[T], T any](ts []T, cb Y) []g.Node {
	nodes := make([]g.Node, 0, len(ts))
	if f, ok := any(cb).(func(T) g.Node); ok {
		for _, t := range ts {
			nodes = append(nodes, f(t))
		}
	} else if f, ok := any(cb).(func(*T) g.Node); ok {
		for i := range ts {
			nodes = append(nodes, f(&ts[i]))
		}
	}
	return nodes
}

func Map2[Y NodeMapper[T], T any](ts []T, cb Y) []g.Node {
	nodes := make([]g.Node, 0, len(ts))
	switch f := interface{}(cb).(type) {
	case func(T) g.Node:
		for _, t := range ts {
			nodes = append(nodes, f(t))
		}
		return nodes
	case func(*T) g.Node:
		for i := range ts {
			nodes = append(nodes, f(&ts[i]))
		}
		return nodes
	}
	panic("unreachable")
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

type Foo struct {
	val string
}

func (f *Foo) GetVal() string {
	return f.val
}

type FooInterface interface {
	GetVal() string
}

// type alias work do work with the proposed interface
// type definitions (i.e removing the equals sign below)
// would require type aproximation and, consequently, reflection
// to discover the underlying type
type FooMapper = func(x *Foo) g.Node

func main() {
	myFoos := []Foo{
		{val: "foo"},
		{val: "bar"},
	}

	// works with func(T) -> Node
	Map2(myFoos, func(x Foo) g.Node {
		return g.Textf(x.val)
	})[0].Render(os.Stdout)

	// works with func(*T) -> Node
	Map2(myFoos, func(x *Foo) g.Node {
		return g.Textf(x.val)
	})[1].Render(os.Stdout)

	var myFn FooMapper
	myFn = func(x *Foo) g.Node {
		return g.El("bla")
	}
	Map2(myFoos, FooMapper(myFn))

	fooIs := []FooInterface{
		&myFoos[0],
		&myFoos[1],
	}
	// works with interfaces
	Map2(fooIs, func(x FooInterface) g.Node {
		return g.Textf(x.GetVal())
	})[0].Render(os.Stdout)

	// the following constucts do not work due to incorrect types
	/*
		Map2(myFoos, func(x *Foo) int {
			return int(0)
		})

		Map2(myFoos, func(x int) g.Node {
			return g.El("foo")
		})
	*/
}
