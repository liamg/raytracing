package raytracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntersection(t *testing.T) {

	triangle := NewTriangle(
		Vector{X: -1, Y: 1, Z: 10},
		Vector{X: -1, Y: -1, Z: 10},
		Vector{X: 1, Y: -1, Z: 10},
		Colour{},
		0,
		0,
	)

	intersections := triangle.FindIntersections(
		Vector{},
		Vector{Z: 1},
	)

	require.Len(t, intersections, 1)

	assert.Equal(t, float64(10), intersections[0])

}

func TestIntersectionFromOffset(t *testing.T) {

	triangle := NewTriangle(
		Vector{X: -1, Y: 1, Z: 10},
		Vector{X: -1, Y: -1, Z: 10},
		Vector{X: 1, Y: -1, Z: 10},
		Colour{},
		0,
		0,
	)

	origin := Vector{X: 0, Y: 0, Z: 3}
	direction := Vector{Z: 1}

	intersections := triangle.FindIntersections(
		origin,
		direction,
	)

	require.Len(t, intersections, 1)

	assert.Equal(t, float64(7), intersections[0])
	assert.Equal(t, Vector{X: 0, Y: 0, Z: 10}, origin.Add(direction.MultiplyN(intersections[0])))

}
