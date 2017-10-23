package deform

import (
	"simplex/lnr"
	"simplex/node"
)

func isSame(a, b lnr.Linegen) bool {
	return a.Id() == b.Id()
}

func castAsNode(o interface{}) *node.Node {
	return o.(*node.Node)
}

//add to hull selection based on range size
// add if range size is greater than 1 : not a segment
func addToSelection(selection *[]*node.Node, hulls ...*node.Node) {
	for _, h := range hulls {
		//add to selection for deformation - if polygon
		if h.Range.Size() > 1 {
			*selection = append(*selection, h)
		}
	}
}
