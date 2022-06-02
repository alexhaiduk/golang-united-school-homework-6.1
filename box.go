package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		//shapes:         make([]Shape, 0),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity != 0 && len(b.shapes) != b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return fmt.Errorf("out of the shapesCapacity")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < len(b.shapes) && i >= 0 {
		return b.shapes[i], nil
	} else {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	extracted := make([]Shape, 0)
	if i >= 0 && i < len(b.shapes) {
		extracted = append(extracted, b.shapes[i])
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return extracted[0], nil
	}
	return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	removed := make([]Shape, 0)
	temp := make([]Shape, 0)
	if i >= 0 && i < len(b.shapes) {
		removed = append(removed, b.shapes[i])
		temp = append(temp, b.shapes[i+1:]...)
		b.shapes = append(b.shapes[:i], shape)
		b.shapes = append(b.shapes[:i+1], temp...)
		return removed[0], nil
	}
	return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcArea()
	}
	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var result error
	if len(b.shapes) > 0 {
		c := 0
		for i := 0; i < len(b.shapes); i++ {
			_, ok := b.shapes[i].(*Circle)
			if ok {
				_, err := b.ExtractByIndex(i)
				if err != nil {
					result = err
				}
				c++
				i--
			}
		}
		if c == 0 {
			result = fmt.Errorf("whether circles are not exist in the list")
		}
	}
	return result
}
