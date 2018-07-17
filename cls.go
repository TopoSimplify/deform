package deform

import (
	"github.com/TopoSimplify/knn"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/rtree"
)

//find context_geom deformable hulls
func SelectFeatureClass(options *opts.Opts, hullDB *rtree.RTree, hull *node.Node) []*node.Node {
	var n int
	var inters, contig bool
	var dict = make(map[[2]int]*node.Node, 0)
	var ctxHulls = knn.FindNodeNeighbours(hullDB, hull, knn.EpsilonDist)

	// for each item in the context_geom list
	for _, cn := range ctxHulls {
		n = 0
		var h = cn.Object.(*node.Node)
		var sameFeature = isSame(hull.Instance, h.Instance)
		// find which item to deform against current hull
		if sameFeature { // check for contiguity
			inters, contig, n = node.IsContiguous(hull, h)
		} else {
			// contiguity is by default false for different features
			contig = false

			inters = hull.Geometry.Intersects(h.Geometry)
			if inters {
				var pts = hull.Geometry.Intersection(h.Geometry)
				inters = len(pts) > 0
				n = len(pts)
			}
		}

		if !inters { // disjoint : nothing to do, continue
			continue
		}

		var sa, sb *node.Node
		if contig && n > 1 { // contiguity with overlap greater than a vertex
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

	var items = make([]*node.Node, 0, len(dict))
	for _, v := range dict {
		items = append(items, v)
	}
	return items
}
