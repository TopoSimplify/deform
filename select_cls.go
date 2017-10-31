package deform

import (
    "simplex/knn"
    "simplex/node"
    "simplex/opts"
    "simplex/db"
)

//find context_geom deformable hulls
func SelectFeatureClass(options *opts.Opts, hullDB *db.DB, hull *node.Node) []*node.Node {
    var n int
    var inters, contig bool
    var dict = make(map[[2]int]*node.Node, 0)
    var ctxHulls = knn.FindNodeNeighbours(hullDB, hull, knn.EpsilonDist)

    // for each item in the context_geom list
    for _, cn := range ctxHulls {
        n = 0
        var h = castAsNode(cn)
        var sameFeature = isSame(hull.Instance, h.Instance)
        // find which item to deform against current hull
        if sameFeature { // check for contiguity
            inters, contig, n = node.IsContiguous(hull, h)
        } else {
            // contiguity is by default false for different features
            contig = false
            var ga, gb = hull.Geom, h.Geom

            inters = ga.Intersects(gb)
            if inters {
                var interpts = ga.Intersection(gb)
                inters = len(interpts) > 0
                n = len(interpts)
            }
        }

        if !inters { // disjoint : nothing to do, continue
            continue
        }

        var sels = make([]*node.Node, 0)
        if contig && n > 1 { // contiguity with overlap greater than a vertex
            sels = contiguousCandidates(hull, h)
        } else if !contig {
            sels = nonContiguousCandidates(options, hull, h)
        }

        // add candidate deformation hulls to selection list
        for _, s := range sels {
            dict[s.Range.AsArray()] = s
        }
    }

    var items = make([]*node.Node, 0)
    for _, v := range dict {
        items = append(items, v)
    }
    return items
}
