package deform

import (
	"simplex/lnr"
	"simplex/node"
	"simplex/knn"
	"github.com/intdxdt/rtree"
)

//find context_geom deformable hulls
func  SelectFeatureClass(self lnr.Linear, hulldb *rtree.RTree, hull *node.Node) []*node.Node {
	var n int
	var inters, contig bool
	var seldict   = make(map[[2]int]*node.Node, 0)
	var ctx_hulls = knn.FindNodeNeighbours(hulldb, hull, knn.EpsilonDist)

	// for each item in the context_geom list
	for _, cn := range ctx_hulls {
		n = 0
		h := castAsNode(cn)

		var same_feature = isSame(hull.Instance,h.Instance)
		// find which item to deform against current hull
		if same_feature { // check for contiguity
			inters, contig, n = node.IsContiguous(hull, h)
		} else {
			// contiguity is by default false for different features
			contig = false
			ga, gb := hull.Geom, h.Geom

			inters = ga.Intersects(gb)
			if inters {
				interpts := ga.Intersection(gb)
				inters = len(interpts) > 0
				n = len(interpts)
			}
		}

		if !inters { // disjoint : nothing to do, continue
			continue
		}

		var sels = make([]*node.Node, 0)
		if contig && n > 1 { // contiguity with overlap greater than a vertex
			sels = _contiguous_candidates(self, hull, h)
		} else if !contig {
			sels = _non_contiguous_candidates(self, hull, h)
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
