package deform

import (
	"simplex/knn"
	"simplex/lnr"
	"simplex/node"
	"github.com/intdxdt/rtree"
)

//find context deformation list
func Select(self lnr.Linear, hulldb *rtree.RTree, hull *node.Node) []*node.Node {
	var seldict = make(map[[2]int]*node.Node, 0)
	var ctx_hulls = knn.FindNodeNeighbours(hulldb, hull, knn.EpsilonDist)

	// for each item in the context list
	for _, cn := range ctx_hulls {
		// find which item to deform against current hull
		h := castAsNode(cn)
		inters, contig, n := node.IsContiguous(hull, h)

		if !inters {
			continue
		}

		sels := make([]*node.Node, 0)
		if contig && n > 1 { //contiguity with overlap greater than a vertex
			sels = _contiguous_candidates(self, hull, h)
		} else if !contig {
			sels = _non_contiguous_candidates(self, hull, h)
		}

		// add candidate deformation hulls to selection list
		for _, s := range sels {
			seldict[s.Range.AsArray()] = s
		}
	}

	items := make([]*node.Node, 0)
	for _, v := range seldict {
		items = append(items, v)
	}
	return items
}

//add to hull selection based on range size
// add if range size is greater than 1 : not a segment
func _add_to_selection(selection *[]*node.Node, hulls ...*node.Node) {
	for _, h := range hulls {
		//add to selection for deformation - if polygon
		if h.Range.Size() > 1 {
			*selection = append(*selection, h)
		}
	}
}
