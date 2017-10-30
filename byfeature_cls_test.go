package deform

import (
	"time"
	"testing"
	"simplex/opts"
	"simplex/dp"
	"simplex/offset"
	"github.com/franela/goblin"
)

func TestSelectFeatureClass(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("constrain by select fc", func() {
		g.It("should test fc selection", func() {
			g.Timeout(1 * time.Hour)

			var options = &opts.Opts{MinDist: 10}
			var coords = linearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			var hulls = createNodes([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, len(coords) - 1}}, coords)
			var inst = dp.New(coords, options, offset.MaxOffset)

			for _, h := range hulls {
				h.Instance = inst
			}

			var db = hullsDB(hulls)

			coords = linearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst
				db.Insert(h)
			}
			var q1 = hulls[0]
			coords = linearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst
				db.Insert(h)
			}
			var q2 = hulls[0]

			var selections = SelectFeatureClass(options, db, q1)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, q2)
			g.Assert(len(selections)).Equal(0)

		})
		g.It("should test fc selection different features", func() {
			g.Timeout(1 * time.Hour)

			var options = &opts.Opts{MinDist: 10}
			var coords = linearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
			var hulls = createNodes([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, len(coords) - 1}}, coords)
			//DebugPrintNodes(hulls)
			var q0 = hulls[2]

			var inst0 = dp.New(coords, options, offset.MaxOffset)

			for _, h := range hulls {
				h.Instance = inst0
			}

			var db = hullsDB(hulls)

			coords = linearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			var inst1 = dp.New(coords, options, offset.MaxOffset)

			for _, h := range hulls {
				h.Instance = inst1
				db.Insert(h)
			}

			var q1 = hulls[0]

			coords = linearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst1
				db.Insert(h)
			}
			var q2 = hulls[0]

			coords = linearCoords("LINESTRING ( 750.5719204078739 667.8504262852285, 731.1163192182406 669.4717263843646, 730.3819045734933 682.6968257108445, 734.5615819289048 700, 740.8441198130572 706.1536411273189, 756.0438082424582 709.5989038379831, 752.801208044186 700, 757.5947734471756 691.9692691038592 )")
			hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)

			for _, h := range hulls {
				h.Instance = inst1
				db.Insert(h)
			}
			var q3 = hulls[0]

			var selections = SelectFeatureClass(options, db, q0)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, q1)
			g.Assert(len(selections)).Equal(1)

			selections = SelectFeatureClass(options, db, q2)
			g.Assert(len(selections)).Equal(0)

			selections = SelectFeatureClass(options, db, q3)
			g.Assert(len(selections)).Equal(0)

			//fmt.Println(q0)
			//fmt.Println(q1)
			//fmt.Println(q2)

		})
		//
		//g.It("should test fc selection different features", func() {
		//	g.Timeout(1 * time.Hour)
		//
		//	var options = &opts.Opts{MinDist: 10}
		//	var coords = linearCoords("LINESTRING ( 780 600, 740 620, 720 660, 720 700, 760 740, 820 760, 860 740, 880 720, 900 700, 880 660, 840 680, 820 700, 800 720, 760 700, 780 660, 820 640, 840 620, 860 580, 880 620, 820 660 )")
		//	var hulls = createNodes([][]int{{0, 3}, {3, 8}, {8, 13}, {13, 17}, {17, len(coords) - 1}}, coords)
		//	//DebugPrintNodes(hulls)
		//	var q0 = hulls[2]
		//
		//	var inst0 = dp.New(coords, options, offset.MaxOffset)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst0
		//	}
		//
		//	var db = hullsDB(hulls)
		//
		//	coords = linearCoords("LINESTRING ( 760 660, 800 620, 800 600, 780 580, 720 580, 700 600 )")
		//	hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)
		//
		//	var inst1 = dp.New(coords, options, offset.MaxOffset)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst1
		//		db.Insert(h)
		//	}
		//
		//	var q1 = hulls[0]
		//	coords = linearCoords("LINESTRING ( 680 640, 660 660, 640 700, 660 740, 720 760, 740 780 )")
		//	hulls = createNodes([][]int{{0, len(coords) - 1}}, coords)
		//
		//	for _, h := range hulls {
		//		h.Instance = inst1
		//		db.Insert(h)
		//	}
		//	var q2 = hulls[0]
		//
		//	var selections = SelectFeatureClass(options, db, q0)
		//	g.Assert(len(selections)).Equal(1)
		//
		//	selections = SelectFeatureClass(options, db, q1)
		//	g.Assert(len(selections)).Equal(1)
		//
		//	selections = SelectFeatureClass(options, db, q2)
		//	g.Assert(len(selections)).Equal(0)
		//
		//	//fmt.Println(q0)
		//	//fmt.Println(q1)
		//	//fmt.Println(q2)
		//
		//})
	})
}