package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		checkPerimeter(t, rectangle, 60.00)

	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.80)
	})
}

func TestGeometry(t *testing.T) {

	shapeTests := []struct {
		name          string
		shape         Shape
		wantArea      float64
		wantPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 20.0}, wantArea: 200.00, wantPerimeter: 60.00},
		{name: "Circle", shape: Circle{Radius: 10.0}, wantArea: 314.00, wantPerimeter: 62.80},
		{name: "Triangle", shape: Triangle{Base: 10.0, Height: 30.0, OneSide: 5.0, OtherSide: 7.0}, wantArea: 150.00, wantPerimeter: 22.00},
	}

	t.Run("Perimeter", func(t *testing.T) {
		for _, shapeTest := range shapeTests {
			t.Run(shapeTest.name, func(t *testing.T) {
				got := shapeTest.shape.Perimeter()
				if got != shapeTest.wantPerimeter {
					t.Errorf("got %.2f want %.2f", got, shapeTest.wantPerimeter)
				}
			})
		}
	})

	t.Run("Area", func(t *testing.T) {
		for _, shapeTest := range shapeTests {
			t.Run(shapeTest.name, func(t *testing.T) {
				got := shapeTest.shape.Area()
				if got != shapeTest.wantArea {
					t.Errorf("got %.2f want %.2f", got, shapeTest.wantArea)
				}
			})
		}
	})

}
