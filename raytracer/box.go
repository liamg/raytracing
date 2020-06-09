package raytracer

type Box struct {
	triangles    [12]*Triangle
	colour       Colour
	specularity  float64
	reflectivity float64
}

func NewBox(pos Vector, size Vector, colour Colour, specularity float64, reflectivity float64) *Box {

	min := pos
	max := pos.Add(size)

	triangles := [12]*Triangle{
		// front 2
		NewTriangle(
			Vector{
				X: min.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: max.X,
				Y: max.Y,
				Z: min.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		//back 2
		NewTriangle(
			Vector{
				X: min.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: max.X,
				Y: max.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		// left face
		NewTriangle(
			Vector{
				X: min.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: min.X,
				Y: max.Y,
				Z: max.Z,
			},
			Vector{
				X: min.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),

		// right face
		NewTriangle(
			Vector{
				X: max.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: max.X,
				Y: max.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: max.Y,
				Z: min.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		// top face
		NewTriangle(
			Vector{
				X: min.X,
				Y: max.Y,
				Z: min.Z,
			},
			Vector{
				X: max.X,
				Y: max.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: max.X,
				Y: max.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: max.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: max.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		// bottom face
		NewTriangle(
			Vector{
				X: min.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: min.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
		NewTriangle(
			Vector{
				X: max.X,
				Y: min.Y,
				Z: max.Z,
			},
			Vector{
				X: max.X,
				Y: min.Y,
				Z: min.Z,
			},
			Vector{
				X: min.X,
				Y: min.Y,
				Z: max.Z,
			},
			colour,
			specularity,
			reflectivity,
		),
	}

	return &Box{
		triangles:    triangles,
		colour:       colour,
		specularity:  specularity,
		reflectivity: reflectivity,
	}
}

func (box *Box) Colour() Colour {
	return box.colour
}

func (box *Box) Specularity() float64 {
	return box.specularity
}

func (box *Box) Reflectivity() float64 {
	return box.reflectivity
}

func (box *Box) FindIntersections(origin Vector, direction Vector) []float64 {

	var tValues []float64

	for _, triangle := range box.triangles {
		ts := triangle.FindIntersections(origin, direction)
		if len(ts) > 0 {
			tValues = append(tValues, ts...)
		}
	}

	return tValues
}

func (box *Box) NormalAtPoint(origin, intersectionPoint Vector) Vector {

	var best *Triangle
	closest := 0.0

	direction := intersectionPoint.Minus(origin)

	//figure out which triangle we intersected
	for i := range box.triangles {
		intersections := box.triangles[i].FindIntersections(origin, direction)
		for _, intersect := range intersections {
			if best == nil || intersect < closest {
				closest = intersect
				best = box.triangles[i]
			}
		}
	}

	if best == nil {
		return Vector{}
	}

	return best.NormalAtPoint(origin, intersectionPoint)

}
