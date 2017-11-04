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
