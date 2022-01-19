package main

import (
	"fmt"
	"math"
)

func main() {
	K := 4
	graph := Graph_newGraph(K)
	graph.Graph_generateEdges()
	graph.Graph_printGraphGLPK("test.jl")
	dGraph := graph.Graph_copyForDynic()

	alg := newAlg(graph)
	abba := Dynic_newDynic(dGraph)
	fmt.Println(alg.Alg_alg(0, pow(2, K)-1))
	fmt.Println(abba.Dynic_maxFlow(0, pow(2, K)-1))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
