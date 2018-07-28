package deform

import (
	"fmt"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/dp"
)

func DebugPrintNodes(ns []*node.Node) {
	for _, n := range ns {
		fmt.Println(n.Geom.WKT())
	}
}

func ctxGeom(wkt string) *ctx.ContextGeometry {
	return ctx.New(geom.ReadGeometry(wkt), 0, -1)
}

func linearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createNodes(indxs [][]int, coords []geom.Point) []node.Node {
	poly := pln.New(coords)
	hulls := make([]node.Node, 0)
	for _, o := range indxs {
		r := rng.Range(o[0], o[1])
		hulls = append(hulls, node.CreateNode(poly.SubCoordinates(r), r, dp.NodeGeometry))
	}
	return hulls
}



//hull geom
func hullGeom(coords []geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0]
	}
	return g
}
