package gogm

import "fmt"

// Vec3 is a vector with 3 components, of type T.
type Vec3[T number] [3]T

// Vec3CopyVec3 copies the content of src to dst.
func Vec3CopyVec3[T1, T2 number](dst *Vec3[T1], src *Vec3[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
	dst[2] = T1(src[2])
}

// Vec3CopyVec2 copies the content of src to dst.
func Vec3CopyVec2[T1, T2 number](dst *Vec3[T1], src *Vec2[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Vec3CopyVec4 copies the content of src to dst.
func Vec3CopyVec4[T1, T2 number](dst *Vec3[T1], src *Vec4[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
	dst[2] = T1(src[2])
}

// String returns a string representation of the vector.
func (v1 *Vec3[T]) String() string {
	return fmt.Sprintf("{%v, %v, %v}", v1[0], v1[1], v1[2])
}

// Inverse sets v1 to the inverse of v2.
// v1 = -v2
func (v1 *Vec3[T]) Inverse(v2 *Vec3[T]) {
	v1[0] = -v2[0]
	v1[1] = -v2[1]
	v1[2] = -v2[2]
}

// Add adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec3[T]) Add(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
	v1[2] = v2[2] + v3[2]
}

// Sub subtracts v2 from v3 component-wise, and stores the result in v1.
// v1 = v2 - v3
func (v1 *Vec3[T]) Sub(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
	v1[2] = v2[2] - v3[2]
}

// Mul multiplies v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 * v3
func (v1 *Vec3[T]) Mul(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] * v3[0]
	v1[1] = v2[1] * v3[1]
	v1[2] = v2[2] * v3[2]
}

// Div divides v2 by v3 component-wise, and stores the result in v1.
// v1 = v2 / v3
func (v1 *Vec3[T]) Div(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] / v3[0]
	v1[1] = v2[1] / v3[1]
	v1[2] = v2[2] / v3[2]
}

// AddS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 + s
func (v1 *Vec3[T]) AddS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] + s
	v1[1] = v2[1] + s
	v1[2] = v2[2] + s
}

// SubS subtracts each component of v2 by s, and stores the result in v1.
// v1 = v2 - s
func (v1 *Vec3[T]) SubS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] - s
	v1[1] = v2[1] - s
	v1[2] = v2[2] - s
}

// MulS multiplies each component of v2 with s, and stores the result in v1.
// v1 = v2 * s
func (v1 *Vec3[T]) MulS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] * s
	v1[1] = v2[1] * s
	v1[2] = v2[2] * s
}

// DivS divides each component of v2 by s, and stores the result in v1.
// v1 = v2 / s
func (v1 *Vec3[T]) DivS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] / s
	v1[1] = v2[1] / s
	v1[2] = v2[2] / s
}

// Cross takes the cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec3[T]) Cross(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0], v1[1], v1[2] = v2[1]*v3[2]-v3[1]*v2[2], v2[2]*v3[0]-v3[2]*v2[0], v2[0]*v3[1]-v3[0]*v2[1]
}

// CrossFast takes the cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec3[T]) CrossFast(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[1]*v3[2] - v3[1]*v2[2]
	v1[1] = v2[2]*v3[0] - v3[2]*v2[0]
	v1[2] = v2[0]*v3[1] - v3[0]*v2[1]
}

// Vec3Dot takes the dot product of v1 and v2, and returns the result.
func Vec3Dot[T number](v1 *Vec3[T], v2 *Vec3[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
