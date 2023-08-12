package math

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func NewVec2(x float64, y float64) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func NewVec2Default() Vec2 {
	return NewVec2(0, 0)
}

func (v Vec2) Add(o Vec2) Vec2 {
	return NewVec2(v.X+o.X, v.Y+o.Y)
}

func (v *Vec2) MutAdd(o Vec2) {
	v.X += o.X
	v.Y += o.Y
}

func (v Vec2) Sub(o Vec2) Vec2 {
	return NewVec2(v.X-o.X, v.Y-o.Y)
}

func (v *Vec2) MutSub(o Vec2) {
	v.X -= o.X
	v.Y -= o.Y
}

func (v *Vec2) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

func (v Vec2) Dot(o Vec2) float64 {
	return (v.X * o.X) + (v.Y * o.Y)
}

func (v *Vec2) Norm() {
  mag := v.Mag();

  v.X = v.X/mag;
  v.Y = v.Y/mag;
}

func (v Vec2) ToNorm() Vec2 {
  mag := v.Mag();

  return NewVec2(v.X/mag, v.Y/mag);
}

func (v Vec2) Mag() float64 {
	return math.Sqrt((v.X*v.X) + (v.Y * v.Y));
}
