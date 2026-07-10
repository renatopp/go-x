package dsx

import "math"

type Vec1 struct {
	X float64
}

func NewVec1(x float64) Vec1 {
	return Vec1{X: x}
}

func (v Vec1) Add(other Vec1) Vec1     { return Vec1{X: v.X + other.X} }
func (v Vec1) Sub(other Vec1) Vec1     { return Vec1{X: v.X - other.X} }
func (v Vec1) Mul(scalar float64) Vec1 { return Vec1{X: v.X * scalar} }
func (v Vec1) Div(scalar float64) Vec1 { return Vec1{X: v.X / scalar} }
func (v Vec1) Neg() Vec1               { return Vec1{X: -v.X} }
func (v Vec1) Abs() Vec1               { return Vec1{X: math.Abs(v.X)} }
func (v Vec1) Dot(other Vec1) float64  { return v.X * other.X }
func (v Vec1) Length() float64         { return math.Abs(v.X) }
func (v Vec1) Normalize() Vec1 {
	length := v.Length()
	if length == 0 {
		return Vec1{X: 0}
	}
	return Vec1{X: v.X / length}
}

type Vec2 struct {
	X, Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Add(other Vec2) Vec2     { return Vec2{X: v.X + other.X, Y: v.Y + other.Y} }
func (v Vec2) Sub(other Vec2) Vec2     { return Vec2{X: v.X - other.X, Y: v.Y - other.Y} }
func (v Vec2) Mul(scalar float64) Vec2 { return Vec2{X: v.X * scalar, Y: v.Y * scalar} }
func (v Vec2) Div(scalar float64) Vec2 { return Vec2{X: v.X / scalar, Y: v.Y / scalar} }
func (v Vec2) Neg() Vec2               { return Vec2{X: -v.X, Y: -v.Y} }
func (v Vec2) Abs() Vec2               { return Vec2{X: math.Abs(v.X), Y: math.Abs(v.Y)} }
func (v Vec2) Dot(other Vec2) float64  { return v.X*other.X + v.Y*other.Y }
func (v Vec2) Length() float64         { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length == 0 {
		return Vec2{X: 0, Y: 0}
	}
	return Vec2{X: v.X / length, Y: v.Y / length}
}

type Vec3 struct {
	X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Add(other Vec3) Vec3 { return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z} }
func (v Vec3) Sub(other Vec3) Vec3 { return Vec3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z} }
func (v Vec3) Mul(scalar float64) Vec3 {
	return Vec3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}
func (v Vec3) Div(scalar float64) Vec3 {
	return Vec3{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}
func (v Vec3) Neg() Vec3              { return Vec3{X: -v.X, Y: -v.Y, Z: -v.Z} }
func (v Vec3) Abs() Vec3              { return Vec3{X: math.Abs(v.X), Y: math.Abs(v.Y), Z: math.Abs(v.Z)} }
func (v Vec3) Dot(other Vec3) float64 { return v.X*other.X + v.Y*other.Y + v.Z*other.Z }
func (v Vec3) Length() float64        { return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z) }
func (v Vec3) Normalize() Vec3 {
	length := v.Length()
	if length == 0 {
		return Vec3{X: 0, Y: 0, Z: 0}
	}
	return Vec3{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}
