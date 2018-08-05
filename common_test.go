package deform

import (
	"fmt"
				"github.com/TopoSimplify/node"
			)

func DebugPrintNodes(ns []*node.Node) {
	for _, n := range ns {
		fmt.Println(n.Geom.WKT())
	}
}

