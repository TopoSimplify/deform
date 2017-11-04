package deform

import (
	"simplex/node"
)

//select contiguous candidates
func contiguousCandidates(a, b *node.Node) (*node.Node , *node.Node ){
	//var selection = make([]*node.Node, 0)
	// compute sidedness relation between contiguous hulls to avoid hull flip
	//var hulls = []*node.Node{a, b}
	//sort.Sort(node.Nodes(hulls))
	// future should not affect the past
	//ha, hb := hulls[0], hulls[1]

	// all hulls that are simple should be collapsible
	// if not collapsible -- add to selection for deformation
	// to reach collapsibility
	var  sa, sb *node.Node
	//& the present should not affect the future
	bln := a.Collapsible(b)
	if !bln {
		sa =  a
	}

	// future should not affect the present
	bln = b.Collapsible(a)
	if !bln {
		sb =  b
	}
	return sa, sb
}
