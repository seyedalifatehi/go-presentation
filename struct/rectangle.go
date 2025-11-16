package main

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) perimeter() float64 {
    return 2.0 * (r.Height + r.Width)
}