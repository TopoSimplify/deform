package deform

import (
	"simplex/pln"
	"simplex/rng"
	"simplex/node"
	"github.com/intdxdt/geom"
)

//hull geom
func hullGeom(coordinates []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coordinates) > 2 {
		g = geom.NewPolygon(coordinates)
	} else if len(coordinates) == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		g = coordinates[0].Clone()
	}
	return g
}


func create_hulls(indxs [][]int, coords []*geom.Point) []*node.Node {
	poly := pln.New(coords)
	hulls := make([]*node.Node, 0)
	for _, o := range indxs {
		hulls = append(hulls, node.NewFromPolyline(poly, rng.NewRange(o[0], o[1]), hullGeom))
	}
	return hulls
}
