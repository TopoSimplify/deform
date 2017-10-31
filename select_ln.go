package deform

import (
    "simplex/db"
    "simplex/knn"
    "simplex/opts"
    "simplex/node"
)

//find context deformation list
func Select(options *opts.Opts, hullDB *db.DB, hull *node.Node) []*node.Node {
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

        var selections = make([]*node.Node, 0)
        if contig && n > 1 { //contiguity with overlap greater than a vertex
            selections = contiguousCandidates(hull, h)
        } else if !contig {
            selections = nonContiguousCandidates(options, hull, h)
        }

        // add candidate deformation hulls to selection list
        for _, s := range selections {
            dict[s.Range.AsArray()] = s
        }
    }

    var items = make([]*node.Node, 0)
    for _, v := range dict {
        items = append(items, v)
    }
    return items
}
