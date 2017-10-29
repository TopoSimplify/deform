package deform

import (
	"simplex/knn"
	"simplex/opts"
	"simplex/node"
	"github.com/intdxdt/rtree"
)

//find context deformation list
func Select(options *opts.Opts, hulldb *rtree.RTree, hull *node.Node) []*node.Node {
	var seldict = make(map[[2]int]*node.Node, 0)
	var ctxHulls = knn.FindNodeNeighbours(hulldb, hull, knn.EpsilonDist)

	// for each item in the context list
	for _, cn := range ctxHulls {
		// find which item to deform against current hull
		h := castAsNode(cn)
		inters, contig, n := node.IsContiguous(hull, h)

		if !inters {
			continue
		}

		sels := make([]*node.Node, 0)
		if contig && n > 1 { //contiguity with overlap greater than a vertex
			sels = contiguousCandidates(hull, h)
		} else if !contig {
			sels = nonContiguousCandidates(options, hull, h)
		}

		// add candidate deformation hulls to selection list
		for _, s := range sels {
			seldict[s.Range.AsArray()] = s
		}
	}

	var items = make([]*node.Node, 0)
	for _, v := range seldict {
		items = append(items, v)
	}
	return items
}
