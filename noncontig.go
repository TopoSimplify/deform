package deform

import (
	"simplex/lnr"
	"simplex/node"
)

//select non-contiguous candidates
func nonContiguousCandidates(self lnr.Linear, a, b *node.Node) []*node.Node {
	var aseg = a.Segment()
	var bseg = b.Segment()

	var aln = a.Polyline
	var bln = b.Polyline

	var asegGeom = aseg.Segment
	var bsegGeom = bseg.Segment

	var alnGeom = aln.Geometry
	var blnGeom = bln.Geometry

	var asegIntersBseg = asegGeom.Intersects(bsegGeom)
	var asegIntersBln = asegGeom.Intersects(blnGeom)
	var bsegIntersAln = bsegGeom.Intersects(alnGeom)
	var alnIntersBln = alnGeom.Intersects(blnGeom)
	var selection = []*node.Node{}

	if asegIntersBseg && asegIntersBln && (!alnIntersBln) {
		addToSelection(&selection, a)
	} else if asegIntersBseg && bsegIntersAln && (!alnIntersBln) {
		addToSelection(&selection, b)
	} else if alnIntersBln {
		// find out whether is a shared vertex or overlap
		// is aseg inter bset  --- dist --- aln inter bln > relax dist
		var ptLns = alnGeom.Intersection(blnGeom)
		var atSeg = aseg.Intersection(bsegGeom)

		// if segs are disjoint but lines intersect, deform a&b
		if len(atSeg) == 0 && len(ptLns) > 0 {
			addToSelection(&selection, a, b)
			return selection
		}

		for _, ptln := range ptLns {
			for _, ptseg := range atSeg {
				delta := ptln.Distance(ptseg)
				if delta > self.Options().RelaxDist {
					addToSelection(&selection, a, b)
					return selection
				}
			}
		}
	}
	return selection
}
