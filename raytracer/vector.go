package raytracer

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) Minus(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector) MultiplyN(n float64) Vector {
	return v.Multiply(Vector{n, n, n})
}

func (v Vector) DivideN(n float64) Vector {
	return v.Divide(Vector{n, n, n})
}

func (v Vector) Multiply(v2 Vector) Vector {
	return Vector{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}

func (v Vector) Divide(v2 Vector) Vector {
	return Vector{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
		Z: v.Z / v2.Z,
	}
}

func (v Vector) DotProduct(v2 Vector) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y) + (v.Z * v2.Z)
}

func (v Vector) CrossProduct(v2 Vector) Vector {
	return Vector{
		X: (v.Y * v2.Z) - (v.Z * v2.Y),
		Y: (v.Z * v2.X) - (v.X * v2.Z),
		Z: (v.X * v2.Y) - (v.Y * v2.X),
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v Vector) Reverse() Vector {
	return v.MultiplyN(-1)
}

func (v Vector) Normalise() Vector {
	return v.DivideN(v.Length())
}
