package deform

import (
	"github.com/intdxdt/rtree"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/knn"
)


//find context deformation list
func Select(options *opts.Opts, hullDB *rtree.RTree, hull *node.Node) []*node.Node {
	var dict = make(map[[2]int]*node.Node, 0)
	var ctxHulls = knn.FindNodeNeighbours(hullDB, hull, knn.EpsilonDist)

	// for each item in the context list
	for _, cn := range ctxHulls {
		// find which item to deform against current hull
		h := castAsNode(cn)
		inters, contig, n := node.IsContiguous(hull, h)

		if !inters {
			continue
		}

		var sa, sb *node.Node
		if contig && n > 1 { //contiguity with overlap greater than a vertex
			sa, sb = contiguousCandidates(hull, h)
		} else if !contig {
			sa, sb = nonContiguousCandidates(options, hull, h)
		}

		// add candidate deformation hulls to selection list
		if sa != nil {
			dict[sa.Range.AsArray()] = sa
		}
		if sb != nil {
			dict[sb.Range.AsArray()] = sb
		}
	}

	var items = make([]*node.Node, 0)
	for _, v := range dict {
		items = append(items, v)
	}
	return items
}
