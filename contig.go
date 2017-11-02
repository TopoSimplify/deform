package deform

import (
	"simplex/node"
	"sort"
)

//select contiguous candidates
func contiguousCandidates(a, b *node.Node) []*node.Node {
	var selection = make([]*node.Node, 0)
	// compute sidedness relation between contiguous hulls to avoid hull flip
	var hulls = []*node.Node{a, b}
	sort.Sort(node.Nodes(hulls))
	// future should not affect the past
	ha, hb := hulls[0], hulls[1]

	// all hulls that are simple should be collapsible
	// if not collapsible -- add to selection for deformation
	// to reach collapsibility

	//& the present should not affect the future
	bln := ha.Collapsible(hb)
	if !bln {
		selection = append(selection, ha)
	}

	// future should not affect the present
	bln = hb.Collapsible(ha)
	if !bln {
		selection = append(selection, hb)
	}
	return selection
}
