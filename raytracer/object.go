package raytracer

type Object interface {
	FindIntersections(origin Vector, destination Vector) []float64
	Colour() Colour
	Specularity() float64  // any number
	Reflectivity() float64 // 0 -> 1
	NormalAtPoint(intersectionPoint Vector) Vector
}
