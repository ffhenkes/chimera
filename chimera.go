package main

/*
#cgo CFLAGS: -I/usr/local/include/igraph
#cgo LDFLAGS: -L/usr/local/lib -ligraph
#include <igraph.h>
*/
import "C"
import "log"

func main() {

	var diameter C.igraph_integer_t
	var graph C.igraph_t

	var v1, v2, v3 C.igraph_integer_t = 1000, 0, 0
	var v4 C.igraph_vector_t

	C.igraph_rng_seed(C.igraph_rng_default(), 42)
	C.igraph_erdos_renyi_game(&graph, C.IGRAPH_ERDOS_RENYI_GNP, v1, 5.0/1000,
		C.IGRAPH_UNDIRECTED, C.IGRAPH_NO_LOOPS)
	C.igraph_diameter(&graph, &diameter, &v2, &v3, &v4, C.IGRAPH_UNDIRECTED, 1)

	log.Println("Diameter: ", diameter)
	C.igraph_destroy(&graph)
}
