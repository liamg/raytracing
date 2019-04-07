package raytracer

type Light struct {
	direction Vector
	position  Vector
	intensity Colour
	lightType LightType
}

type LightType uint8

const (
	Ambient LightType = iota
	Point
	Directional
)

func NewAmbientLight(intensity Colour) Light {
	return Light{
		lightType: Ambient,
		intensity: intensity,
	}
}

func NewPointLight(intensity Colour, point Vector) Light {
	return Light{
		lightType: Point,
		intensity: intensity,
		position:  point,
	}
}

func NewDirectionalLight(intensity Colour, direction Vector) Light {
	return Light{
		lightType: Directional,
		intensity: intensity,
		direction: direction,
	}
}
