package main

import (
	"testing"

	g "github.com/maragudk/gomponents"
)

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
