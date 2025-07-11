// v is a 2d vector package
package v

import (
	"fmt"
	"math"
)

var (
	// One Vec{1, 1} vector is a vector with all components set to 1.
	One = Vec{1, 1}
	// Left unit vector. Vec{-1, 0} Represents the direction of left.
	Left = Vec{-1, 0}
	// Right unit vector. Vec{1, 0} Represents the direction of right.
	Right = Vec{1, 0}
	// Up unit vector. Vec{0, -1} Y is down in 2D, so this vector points -Y.
	Up = Vec{0, -1}
	// Down unit vector. Vec{0, 1} Y is down in 2D, so this vector points +Y.
	Down = Vec{0, 1}
)

type Vec struct {
	X, Y float64
}

// Add returns this + a
func (v Vec) Add(a Vec) Vec {
	return Vec{v.X + a.X, v.Y + a.Y}
}

// Sub returns this - a
func (v Vec) Sub(a Vec) Vec {
	return Vec{v.X - a.X, v.Y - a.Y}
}

// Div divides this vector by a.
func (v Vec) Div(a Vec) Vec {
	return Vec{v.X / a.X, v.Y / a.Y}
}

// DivS divides this vector by scalar value s.
func (v Vec) DivS(s float64) Vec {
	return Vec{v.X / s, v.Y / s}
}

// Mul returns this * a
func (v Vec) Mul(a Vec) Vec {
	return Vec{v.X * a.X, v.Y * a.Y}
}

// Scale scales vector
func (v Vec) Scale(s float64) Vec {
	return Vec{v.X * s, v.Y * s}
}

// Unit returns a normalized copy of this vector (unit vector).
func (v Vec) Unit() Vec {
	// return v.Mult(1.0 / (v.Length() + math.SmallestNonzeroFloat64))
	return v.Scale(1.0 / (v.Mag() + 1e-50))
}

// Abs returns the absolute value of vector.
func (v Vec) Abs() Vec {
	return Vec{math.Abs(v.X), math.Abs(v.Y)}
}

// AbsX returns the absolute X value of vector.
func (v Vec) AbsX() float64 {
	return math.Abs(v.X)
}

// AbsY returns the absolute Y value of vector.
func (v Vec) AbsY() float64 {
	return math.Abs(v.Y)
}

// Neg negates a vector.
func (v Vec) Neg() Vec {
	return Vec{-v.X, -v.Y}
}

// NegY negates X.
func (v Vec) NegX() Vec {
	return Vec{-v.X, v.Y}
}

// NegY negates Y.
func (v Vec) NegY() Vec {
	return Vec{v.X, -v.Y}
}

// Dot returns dot product
func (v Vec) Dot(other Vec) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross calculates the 2D vector cross product analog.
// The cross product of 2D vectors results in a 3D vector with only a z component.
// This function returns the magnitude of the z value.
func (v Vec) Cross(other Vec) float64 {
	return v.X*other.Y - v.Y*other.X
}

// Returns the vector projection onto other.
func (v Vec) Project(other Vec) Vec {
	return other.Scale(v.Dot(other) / other.Dot(other))
}

// Angle returns the angular direction v is pointing in (in radians).
func (v Vec) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Rotate a vector by an angle in radians
func (v Vec) Rotate(angle float64) Vec {
	return Vec{
		X: v.X*math.Cos(angle) - v.Y*math.Sin(angle),
		Y: v.X*math.Sin(angle) + v.Y*math.Cos(angle),
	}
}

// Mag returns the magnitude (length) of the vector.
func (v Vec) Mag() float64 {
	return math.Hypot(v.X, v.Y)
}

// MagSq returns the magnitude (length) of the vector, squared.
//
// This method is often used to improve performance since, unlike Mag(),
// it does not require a Sqrt() operation.
func (v Vec) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Slerp performs spherical linear interpolation between two vectors with given weight value in [0,1] range, returning interpolated vector
func (v Vec) Slerp(to Vec, weight float64) Vec {
	startLengthSq := v.MagSq()
	endLengthSq := to.MagSq()
	if startLengthSq == 0.0 || endLengthSq == 0.0 {
		return v.Lerp(to, weight)
	}
	startLength := math.Sqrt(startLengthSq)
	resultLength := (1-weight)*startLength + weight*math.Sqrt(endLengthSq)
	angle := v.AngleTo(to)
	return v.Rotate(angle * weight).Scale(resultLength / startLength)
}

// AngleTo returns the angle to the given vector, in radians.
func (v Vec) AngleTo(other Vec) float64 {
	return math.Atan2(v.Cross(other), v.Dot(other))
}

// Limits a vector's magnitude to a maximum value.
func (v Vec) Limit(max float64) Vec {
	if v.Mag() > max {
		return v.Unit().Scale(max)
	}
	return v
}

// Lerp linearly interpolates between this and other vector.
func (v Vec) Lerp(other Vec, t float64) Vec {
	return v.Scale(1.0 - t).Add(other.Scale(t))
}

// IsZero returns true if vector is zero vector
func (v Vec) IsZero() bool {
	return v == Vec{}
}

// Dist returns distance between v and other.
func (v Vec) Dist(other Vec) float64 {
	return math.Hypot(v.X-other.X, v.Y-other.Y)
}

// DistSq returns the squared distance between this and other.
//
// Faster than v.Dist() when you only need to compare distances.
func (v Vec) DistSq(other Vec) float64 {
	return v.Sub(other).MagSq()
}

// Round returns the nearest integer Vector, rounding half away from zero.
func (v Vec) Round() Vec {
	return Vec{math.Round(v.X), math.Round(v.Y)}
}

// Floor returns vector with all components rounded down (towards negative infinity).
func (v Vec) Floor() Vec {
	return Vec{math.Floor(v.X), math.Floor(v.Y)}
}

// Ceil returns vector with all components rounded up (towards positive infinity).
func (v Vec) Ceil() Vec {
	return Vec{math.Ceil(v.X), math.Ceil(v.Y)}
}

// FromAngle makes a new 2D unit vector from an angle
func FromAngle(angle float64) Vec {
	return Vec{math.Cos(angle), math.Sin(angle)}
}

// EqualsP returns they are practically equal with each other within a delta tolerance.
func (v Vec) EqualsPr(other Vec, allowedDelta float64) bool {
	return (math.Abs(v.X-other.X) <= allowedDelta) &&
		(math.Abs(v.Y-other.Y) <= allowedDelta)
}

// Equals checks if two vectors are equal. (Be careful when comparing floating point numbers!)
func (v Vec) Equals(other Vec) bool {
	return v.X == other.X && v.Y == other.Y
}

// Reflect returns the reflection of the vector v over the given normal.
// normal should be a normalized (unit) vector.
func (v Vec) Reflect(normal Vec) Vec {
    return v.Sub(normal.Scale(2 * v.Dot(normal)))
}

// String returns string representation of this vector.
func (v Vec) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", v.X, v.Y)
}
