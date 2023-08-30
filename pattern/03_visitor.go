package main

import (
	"fmt"
	"math"
)

/*

Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
не изменяя классы объектов, над которыми эти операции могут выполняться.

Плюсы:
- Упрощает добавление операций, работающих со сложными структурами объектов.
- Объединяет родственные операции в одном классе.
- Посетитель может накапливать состояние при обходе структуры элементов.

Минусы:
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.

*/

func main() {
	shapes := []Acceptor{
		&Circle{r: 3.0},
		&Rectangle{w: 10, h: 30},
	}

	areaCalculator := &AreaCalculator{}

	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}
}

type AreaCalculator struct{}

func (calculator *AreaCalculator) VisitCircle(circle *Circle) {
	area := math.Pi * math.Pow(circle.r, 2)
	fmt.Printf("area shape \"%s\" = %.2f\n", circle.Type(), area)
}

func (calculator *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.w * rectangle.h
	fmt.Printf("area shape \"%s\" = %.2f\n", rectangle.Type(), area)
}

type Acceptor interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

type Circle struct {
	r float64
}

func (circle *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(circle)
}

func (circle *Circle) Type() string {
	return "circle"
}

type Rectangle struct {
	w, h float64
}

func (rectangle *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(rectangle)
}

func (rectangle *Rectangle) Type() string {
	return "rectangle"
}
