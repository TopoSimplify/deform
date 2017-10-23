package deform

import (
	"simplex/lnr"
	"simplex/node"
)

//select non-contiguous candidates
func  _non_contiguous_candidates(self lnr.Linear, a, b *node.Node) []*node.Node {
	var aseg = a.Segment()
	var bseg = b.Segment()

	var aln = a.SubPolyline()
	var bln = b.SubPolyline()

	var aseg_geom = aseg.Segment
	var bseg_geom = bseg.Segment

	var aln_geom = aln.Geometry
	var bln_geom = bln.Geometry

	var aseg_inters_bseg = aseg_geom.Intersects(bseg_geom)
	var aseg_inters_bln  = aseg_geom.Intersects(bln_geom)
	var bseg_inters_aln  = bseg_geom.Intersects(aln_geom)
	var aln_inters_bln   = aln_geom.Intersects(bln_geom)

	selection := []*node.Node{}
	if aseg_inters_bseg && aseg_inters_bln && (!aln_inters_bln) {
		_add_to_selection(&selection, a)
	} else if aseg_inters_bseg && bseg_inters_aln && (!aln_inters_bln) {
		_add_to_selection(&selection, b)
	} else if aln_inters_bln {
		// find out whether is a shared vertex or overlap
		// is aseg inter bset  --- dist --- aln inter bln > relax dist
		pt_lns := aln_geom.Intersection(bln_geom)
		at_seg := aseg.Intersection(bseg_geom)

		// if segs are disjoint but lines intersect, deform a&b
		if len(at_seg) == 0 && len(pt_lns) > 0 {
			_add_to_selection(&selection, a, b)
			return selection
		}

		for _, ptln := range pt_lns {
			for _, ptseg := range at_seg {
				delta := ptln.Distance(ptseg)
				if delta > self.Options().RelaxDist {
					_add_to_selection(&selection, a, b)
					return selection
				}
			}
		}
	}
	return selection
}
