package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Round(2*3.14*c.Radius*100) / 100
}

func (c Circle) Area() float64 {
	return math.Round(3.14*c.Radius*c.Radius*100) / 100
}

type Triangle struct {
	Base      float64
	Height    float64
	OneSide   float64
	OtherSide float64
}

func (t Triangle) Perimeter() float64 {
	return math.Round((t.Base+t.OneSide+t.OtherSide)*100) / 100
}

func (t Triangle) Area() float64 {
	return math.Round(0.5*t.Base*t.Height*100) / 100
}

type Shape interface {
	Perimeter() float64
	Area() float64
}
