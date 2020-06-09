package raytracer

import "math"

type Matrix3x3 [9]float64

func RotateXY(v Vector, xAngle, yAngle float64) Vector {

	xMatrix := CreateRotationMatrixAboutX(xAngle)
	yMatrix := CreateRotationMatrixAboutY(yAngle)

	v = Vector{
		X: (v.X * xMatrix[0]) + (v.Y * xMatrix[1]) + (v.Z * xMatrix[2]),
		Y: (v.X * xMatrix[3]) + (v.Y * xMatrix[4]) + (v.Z * xMatrix[5]),
		Z: (v.X * xMatrix[6]) + (v.Y * xMatrix[7]) + (v.Z * xMatrix[8]),
	}

	return Vector{
		X: (v.X * yMatrix[0]) + (v.Y * yMatrix[1]) + (v.Z * yMatrix[2]),
		Y: (v.X * yMatrix[3]) + (v.Y * yMatrix[4]) + (v.Z * yMatrix[5]),
		Z: (v.X * yMatrix[6]) + (v.Y * yMatrix[7]) + (v.Z * yMatrix[8]),
	}
}

func CreateRotationMatrixAboutX(angle float64) Matrix3x3 {
	return Matrix3x3(
		[9]float64{
			1, 0, 0,
			0, math.Cos(angle), -math.Sin(angle),
			0, math.Sin(angle), math.Cos(angle),
		},
	)
}

func CreateRotationMatrixAboutY(angle float64) Matrix3x3 {
	return Matrix3x3(
		[9]float64{
			math.Cos(angle), 0, math.Sin(angle),
			0, 1, 0,
			-math.Sin(angle), 0, math.Cos(angle),
		},
	)
}
