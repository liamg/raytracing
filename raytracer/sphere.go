package raytracer

import (
	"math"
)

type Sphere struct {
	position     Vector
	radius       float64
	colour       Colour
	specularity  float64
	reflectivity float64
}

func NewSphere(position Vector, radius float64, colour Colour, specularity float64, reflectivity float64) *Sphere {
	return &Sphere{
		position:     position,
		radius:       radius,
		colour:       colour,
		specularity:  specularity,
		reflectivity: reflectivity,
	}
}

func (sphere *Sphere) Colour() Colour {
	return sphere.colour
}

func (sphere *Sphere) Specularity() float64 {
	return sphere.specularity
}

func (sphere *Sphere) Reflectivity() float64 {
	return sphere.reflectivity
}

func (sphere *Sphere) FindIntersections(origin Vector, direction Vector) []float64 {

	oc := origin.Minus(sphere.position)

	k1 := direction.DotProduct(direction)
	k2 := 2 * oc.DotProduct(direction)
	k3 := oc.DotProduct(oc) - (sphere.radius * sphere.radius)

	discriminant := (k2 * k2) - (4 * k1 * k3)
	if discriminant < 0 {
		return []float64{}
	}

	t1 := (-k2 + math.Sqrt(discriminant)) / (2 * k1)
	t2 := (-k2 - math.Sqrt(discriminant)) / (2 * k1)

	return []float64{t1, t2}
}

func (sphere *Sphere) NormalAtPoint(origin, intersectionPoint Vector) Vector {
	return intersectionPoint.Minus(sphere.position)
}
