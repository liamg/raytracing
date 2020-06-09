package raytracer

import (
	"math"
)

type Triangle struct {
	p1           Vector
	p2           Vector
	p3           Vector
	colour       Colour
	specularity  float64
	reflectivity float64
}

func NewTriangle(p1 Vector, p2 Vector, p3 Vector, colour Colour, specularity float64, reflectivity float64) *Triangle {
	return &Triangle{
		p1:           p1,
		p2:           p2,
		p3:           p3,
		colour:       colour,
		specularity:  specularity,
		reflectivity: reflectivity,
	}
}

func (triangle *Triangle) Colour() Colour {
	return triangle.colour
}

func (triangle *Triangle) Specularity() float64 {
	return triangle.specularity
}

func (triangle *Triangle) Reflectivity() float64 {
	return triangle.reflectivity
}

func (triangle *Triangle) FindIntersections(origin Vector, direction Vector) []float64 {

	// normal to triangle plane
	normal := (triangle.p2.Minus(triangle.p1)).CrossProduct(triangle.p3.Minus(triangle.p1)) // N

	// Step 1: finding P

	// check if ray and plane are parallel ?
	nDirection := normal.DotProduct(direction)
	if math.Abs(nDirection) < 0.00000001 { // almost 0
		return nil // they are parallel so they don't intersect !
	}

	// compute d parameter
	d := normal.DotProduct(triangle.p1)

	// compute t
	t := (d - normal.DotProduct(origin)) / nDirection
	// check if the triangle is in behind the ray
	if t < 0 {
		return nil // the triangle is behind
	}

	// compute the intersection point on the plane
	intersection := origin.Add(direction.MultiplyN(t))

	//fmt.Printf("\n\nIntersect = %#v || Origin = %#v || Direction = %#v || T = %f\n", intersection, origin, direction, t)

	// Step 2: inside-outside test

	// edge 0
	edge0 := triangle.p2.Minus(triangle.p1)
	vp0 := intersection.Minus(triangle.p1)
	c := edge0.CrossProduct(vp0)
	if normal.DotProduct(c) < 0 {
		return nil // intersection is on the right side
	}

	// edge 1
	edge1 := triangle.p3.Minus(triangle.p2)
	vp1 := intersection.Minus(triangle.p2)
	c = edge1.CrossProduct(vp1)
	if normal.DotProduct(c) < 0 {
		return nil // P is on the right side
	}

	// edge 2
	edge2 := triangle.p1.Minus(triangle.p3)
	vp2 := intersection.Minus(triangle.p3)
	c = edge2.CrossProduct(vp2)
	if normal.DotProduct(c) < 0 {
		return nil
	} // P is on the right side;

	//fmt.Printf("TRI %f nDir %f normal %#v dir %#v\n", t, nDirection, normal, direction)

	return []float64{t}
}

func (triangle *Triangle) NormalAtPoint(origin, _ Vector) Vector {

	normal := (triangle.p2.Minus(triangle.p1)).CrossProduct(triangle.p3.Minus(triangle.p1)) // N
	inverse := normal.Reverse()

	if triangle.p1.Add(normal).Minus(origin).Length() < triangle.p1.Add(inverse).Minus(origin).Length() {
		return normal
	}

	return inverse
}
