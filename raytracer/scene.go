package raytracer

import "math"

type Scene struct {
	canvasSize       Vector
	viewportSize     Vector
	viewportDistance float64 //distance of viewport from the camera
	objects          []Object
	backgroundColour Colour
	lights           []Light
}

func NewScene(canvasSize Vector) *Scene {
	return &Scene{
		canvasSize:       canvasSize,
		viewportSize:     Vector{X: 1, Y: 1},
		viewportDistance: 1,
		objects:          []Object{},
		lights:           []Light{},
	}
}

func (scene *Scene) AddObject(object Object) {
	scene.objects = append(scene.objects, object)
}

func (scene *Scene) AddLight(light Light) {
	scene.lights = append(scene.lights, light)
}

func (scene *Scene) CanvasToViewport(canvasPoint Vector) Vector {
	return Vector{
		X: canvasPoint.X * (scene.viewportSize.X / scene.canvasSize.X),
		Y: canvasPoint.Y * (scene.viewportSize.Y / scene.canvasSize.Y),
		Z: scene.viewportDistance,
	}
}

func (scene *Scene) findClosestIntersection(origin Vector, destination Vector, tMin float64, tMax float64) (Object, float64) {
	var closestT float64
	var closestObject Object
	for i := range scene.objects {
		object := scene.objects[i]
		ts := object.FindIntersections(origin, destination)
		for _, t := range ts {
			if t >= tMin && (t <= tMax || tMax < tMin) && (t < closestT || closestObject == nil) {
				closestT = t
				closestObject = object
			}
		}
	}
	return closestObject, closestT
}

// TraceRay traces a ray from the observer origin toward the destination
func (scene *Scene) TraceRay(origin Vector, destination Vector, tMin float64, tMax float64, depth int) Colour {

	closestObject, closestT := scene.findClosestIntersection(origin, destination, tMin, tMax)
	if closestObject == nil {
		return scene.backgroundColour
	}
	intersectionPoint := origin.Add(destination.Multiply(Vector{closestT, closestT, closestT}))
	normal := closestObject.NormalAtPoint(intersectionPoint)

	// ensure vector has length 1
	l := normal.Length()
	normal = normal.Divide(Vector{l, l, l})

	localColour := closestObject.Colour().Multiply(scene.computeLighting(intersectionPoint, normal, destination.Reverse(), closestObject.Specularity()))

	// If we hit the recursion limit or the object is not reflective, we're done
	reflectivity := closestObject.Reflectivity()
	if depth <= 0 || reflectivity <= 0 {
		return localColour
	}

	// Compute the reflected color
	reflection := reflectRay(destination.Reverse(), normal)
	reflectedColour := scene.TraceRay(intersectionPoint, reflection, 0.001, -1, depth-1)

	return localColour.MultiplyN(1 - reflectivity).Add(reflectedColour.MultiplyN(reflectivity))

}

func reflectRay(ray Vector, normal Vector) Vector {
	return normal.MultiplyN(2).MultiplyN(ray.DotProduct(normal)).Minus(ray)
}

func (scene *Scene) computeLighting(point Vector, normal Vector, toCamera Vector, specularity float64) Colour {
	intensity := Colour{}
	tMax := 0.0
	for _, light := range scene.lights {
		if light.lightType == Ambient {
			intensity = intensity.Add(light.intensity)
		} else {
			var lightVector Vector
			if light.lightType == Point {
				lightVector = light.position.Minus(point)
				tMax = 1
			} else {
				lightVector = light.direction
				tMax = -1
			}

			// Shadow check - is anything blocking this light?
			shadowObject, _ := scene.findClosestIntersection(point, lightVector, 0.000001, tMax)
			if shadowObject != nil {
				continue
			}

			//# Diffuse
			nl := normal.DotProduct(lightVector)
			if nl > 0 {
				chanIntensity := nl / (normal.Length() * lightVector.Length())
				intensity = intensity.Add(light.intensity.Multiply(Colour{chanIntensity, chanIntensity, chanIntensity}))
			}

			// Specular
			if specularity > 0 {
				reflection := reflectRay(lightVector, normal)
				rv := reflection.DotProduct(toCamera)
				if rv > 0 {
					intensity = intensity.Add(light.intensity.MultiplyN(math.Pow(rv/(reflection.Length()*toCamera.Length()), specularity)))
				}
			}
		}
	}
	return intensity
}
