package deform

import (
	"simplex/lnr"
	"simplex/node"
)

//select contiguous candidates
func contiguousCandidates(self lnr.Linear, a, b *node.Node) []*node.Node {
	var selection = make([]*node.Node, 0)
	// compute sidedness relation between contiguous hulls to avoid hull flip
	hulls := node.NewNodes().Extend(a, b).Sort()
	// future should not affect the past
	ha, hb := hulls.Get(0), hulls.Get(1)

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
