package gogm

import (
	"constraints"
	"fmt"
	"math"
)

type number interface {
	constraints.Integer | constraints.Float
}

// Vec2 is a vector with 2 components, of type T.
type Vec2[T number] [2]T

// Vec2CopyVec2 copies the content of src to dst.
func Vec2CopyVec2[T1, T2 number](dst *Vec2[T1], src *Vec2[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Vec2CopyVec3 copies the content of src to dst.
func Vec2CopyVec3[T1, T2 number](dst *Vec2[T1], src *Vec3[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Vec2CopyVec4 copies the content of src to dst.
func Vec2CopyVec4[T1, T2 number](dst *Vec2[T1], src *Vec4[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// String returns a string representation of the vector.
func (v1 *Vec2[T]) String() string {
	return fmt.Sprintf("{%v, %v}", v1[0], v1[1])
}

// Len returns the length of the vector.
func (v1 *Vec2[T]) Len() float64 {
	return math.Hypot(float64(v1[0]), float64(v1[1]))
}

// Normalize normalizes v2, and stores the result in v1.
func (v1 *Vec2[T]) Normalize(v2 *Vec2[T]) {
	l := T(v2.Len())
	v1[0] = v2[0] / l
	v1[1] = v2[1] / l
}

// Inverse sets v1 to the inverse of v2.
// v1 = -v2
func (v1 *Vec2[T]) Inverse(v2 *Vec2[T]) {
	v1[0] = -v2[0]
	v1[1] = -v2[1]
}

// Add adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec2[T]) Add(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
}

// Sub subtracts v2 from v3 component-wise, and stores the result in v1.
// v1 = v2 - v3
func (v1 *Vec2[T]) Sub(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
}

// Mul multiplies v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 * v3
func (v1 *Vec2[T]) Mul(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] * v3[0]
	v1[1] = v2[1] * v3[1]
}

// Div divides v2 by v3 component-wise, and stores the result in v1.
// v1 = v2 / v3
func (v1 *Vec2[T]) Div(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] / v3[0]
	v1[1] = v2[1] / v3[1]
}

// AddS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 + s
func (v1 *Vec2[T]) AddS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] + s
	v1[1] = v2[1] + s
}

// SubS subtracts each component of v2 by s, and stores the result in v1.
// v1 = v2 - s
func (v1 *Vec2[T]) SubS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] - s
	v1[1] = v2[1] - s
}

// MulS multiplies each component of v2 with s, and stores the result in v1.
// v1 = v2 * s
func (v1 *Vec2[T]) MulS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] * s
	v1[1] = v2[1] * s
}

// DivS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 / s
func (v1 *Vec2[T]) DivS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] / s
	v1[1] = v2[1] / s
}

// Cross takes the cross product of v1 and v2, and returns the result.
func (v1 *Vec2[T]) Cross(v2 *Vec2[T]) T {
	return v1[0]*v2[1] - v1[1]*v2[0]
}

// Dot takes the dot product of v1 and v2, and returns the result.
func (v1 *Vec2[T]) Dot(v2 *Vec2[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1]
}
