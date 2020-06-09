package raytracer

import (
	"math"
)

type Scene struct {
	canvasSize       Vector
	viewportSize     Vector
	viewportDistance float64 //distance of viewport from the camera
	objects          []Object
	backgroundColour Colour
	lights           []Light
}

type Option func(scene *Scene)

func WithCanvasSize(size Vector) Option {
	return func(scene *Scene) {
		scene.canvasSize = size
	}
}

func WithBackgroundColour(colour Colour) Option {
	return func(scene *Scene) {
		scene.backgroundColour = colour
	}
}

func NewScene(options ...Option) *Scene {
	scene := &Scene{
		canvasSize:       Vector{X: 500, Y: 500},
		viewportSize:     Vector{X: 1, Y: 1, Z: 1},
		objects:          []Object{},
		lights:           []Light{},
		backgroundColour: Colour{0.5, 0.8, 0.88},
	}
	for _, option := range options {
		option(scene)
	}
	return scene
}

func (scene *Scene) ResetObjects() {
	scene.objects = []Object{}
}

func (scene *Scene) Resize(x, y int) {
	scene.canvasSize = Vector{
		X: float64(x),
		Y: float64(y),
	}
}

func (scene *Scene) AddObject(object Object) {
	scene.objects = append(scene.objects, object)
}

func (scene *Scene) AddLight(light Light) {
	scene.lights = append(scene.lights, light)
}

func (scene *Scene) ResetLighting() {
	scene.lights = []Light{}
}

func (scene *Scene) CanvasToViewport(canvasPoint Vector) Vector {
	return Vector{
		X: canvasPoint.X * (scene.viewportSize.X / scene.canvasSize.X),
		Y: canvasPoint.Y * (scene.viewportSize.Y / scene.canvasSize.Y),
		Z: scene.viewportSize.Z,
	}
}

func (scene *Scene) findClosestIntersection(origin Vector, direction Vector, tMin float64, tMax float64) (Object, float64) {
	var closestT float64
	var closestObject Object
	for i := range scene.objects {
		object := scene.objects[i]
		ts := object.FindIntersections(origin, direction)
		if ts != nil {
			for _, t := range ts {
				if t >= tMin && (t <= tMax || tMax < tMin) && (t < closestT || closestObject == nil) {
					closestT = t
					closestObject = object
				}
			}
		}
	}
	return closestObject, closestT
}

// TraceRay traces a ray from the observer origin in the given direction
func (scene *Scene) TraceRay(origin Vector, direction Vector, tMin float64, tMax float64, depth int) (Colour, float64) {

	//fmt.Printf("DIR %#v\n", direction)

	closestObject, closestT := scene.findClosestIntersection(origin, direction, tMin, tMax)
	if closestObject == nil {
		return scene.backgroundColour, tMax
	}
	intersectionPoint := origin.Add(direction.MultiplyN(closestT))
	normal := closestObject.NormalAtPoint(origin, intersectionPoint)

	// TODO is this required if objects work correctly?
	if normal.Length() == 0 {
		panic("missing normal")
	}

	normal = normal.Normalise()

	localColour := closestObject.Colour().Multiply(scene.computeLighting(intersectionPoint, normal, direction.Reverse(), closestObject.Specularity()))

	// If we hit the recursion limit or the object is not reflective, we're done
	reflectivity := closestObject.Reflectivity()
	if depth <= 0 || reflectivity <= 0 {
		return localColour, closestT
	}

	// Compute the reflected color
	reflection := reflectRay(direction.Reverse(), normal)

	reflectedColour, _ := scene.TraceRay(intersectionPoint, reflection, 0.001, -1, depth-1)

	return localColour.MultiplyN(1 - reflectivity).Add(reflectedColour.MultiplyN(reflectivity)), closestT

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
