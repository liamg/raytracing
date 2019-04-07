package raytracer

type Colour struct {
	R float64
	G float64
	B float64
}

func (c Colour) Multiply(c2 Colour) Colour {
	return Colour{
		R: cap(c.R * c2.R),
		G: cap(c.G * c2.G),
		B: cap(c.B * c2.B),
	}
}

func (c Colour) MultiplyN(n float64) Colour {
	return c.Multiply(Colour{n, n, n})
}

func (c Colour) Add(c2 Colour) Colour {
	return Colour{
		R: cap(c.R + c2.R),
		G: cap(c.G + c2.G),
		B: cap(c.B + c2.B),
	}
}

func cap(n float64) float64 {
	if n > 1 {
		return 1
	}
	if n < 0 {
		n = 0
	}
	return n
}
