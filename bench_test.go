package main

import (
	"testing"

	g "github.com/maragudk/gomponents"
)

func Benchmark_Map3_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		// since render is itself a generic function
		// go can't infer how it should be called
		// so we tell the compiler we want a pointer
		Map3[func(*BigStruct10x) g.Node](bs, render)
	}
}

func Benchmark_Map2_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		// since render is itself a generic function
		// go can't infer how it should be called
		// so we tell the compiler we want a pointer
		Map2[func(*BigStruct10x) g.Node](bs, render)
	}
}

func Benchmark_MapRefPreAlloc_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		MapRefPreAlloc(bs, render)
	}
}

func Benchmark_MapRefPreAlloc_BigStruct(B *testing.B) {
	bs := make([]BigStruct, 1000)
	for i := 0; i < B.N; i++ {
		MapRefPreAlloc(bs, render)
	}
}

func Benchmark_MapRef_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		MapRef(bs, render)
	}
}

func Benchmark_MapRef_BigStruct(B *testing.B) {
	bs := make([]BigStruct, 1000)
	for i := 0; i < B.N; i++ {
		MapRef(bs, render)
	}
}

func Benchmark_gMapPreAlloc_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		MapPreAlloc(bs, render)
	}
}

func Benchmark_gMapPreAlloc_BigStruct(B *testing.B) {
	bs := make([]BigStruct, 1000)
	for i := 0; i < B.N; i++ {
		MapPreAlloc(bs, render)
	}
}

func Benchmark_gMap_BigStruct10x(B *testing.B) {
	bs := make([]BigStruct10x, 1000)
	for i := 0; i < B.N; i++ {
		g.Map(bs, render)
	}
}

func Benchmark_gMap_BigStruct(B *testing.B) {
	bs := make([]BigStruct, 1000)
	for i := 0; i < B.N; i++ {
		g.Map(bs, render)
	}
}
