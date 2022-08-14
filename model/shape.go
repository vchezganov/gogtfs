//Done
package model

type Shape struct {
	ShapeID           string  `csv:"shape_id"`
	ShapePtLat        float64 `csv:"shape_pt_lat"`
	ShapePtLon        float64 `csv:"shape_pt_lon"`
	ShapePtSequence   uint    `csv:"shape_pt_sequence"`
	ShapeDistTraveled float64 `csv:"shape_dist_traveled"`
}

func (s Shape) TableFileName() string {
	return "shapes.txt"
}
