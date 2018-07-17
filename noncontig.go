package deform

import (
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/node"
)

//select non-contiguous candidates
func nonContiguousCandidates(options *opts.Opts, a, b *node.Node) (*node.Node, *node.Node) {
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
	var sa, sb *node.Node

	if asegIntersBseg && asegIntersBln && (!alnIntersBln) {
		sa = a
	} else if asegIntersBseg && bsegIntersAln && (!alnIntersBln) {
		sb = b
	} else if alnIntersBln {
		// find out whether is a shared vertex or overlap
		// is aseg inter bset  --- dist --- aln inter bln > relax dist
		var ptLns = alnGeom.Intersection(blnGeom)
		var atSeg = aseg.Intersection(bsegGeom)

		// if segs are disjoint but lines intersect, deform a&b
		if len(atSeg) == 0 && len(ptLns) > 0 {
			sa, sb = a, b
		} else {
			outer: for i := range ptLns {
				for j := range atSeg {
					delta := ptLns[i].Distance(atSeg[j])
					if delta > options.RelaxDist {
						sa, sb = a, b
						break outer
					}
				}
			}
		}
	}

	return sa, sb
}
