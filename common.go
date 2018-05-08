package deform

import (
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/node"
)

func isSame(a, b lnr.Linegen) bool {
	return a.Id() == b.Id()
}

func castAsNode(o interface{}) *node.Node {
	return o.(*node.Node)
}
