package gogm

import "fmt"

// Vec4 is a vector with 4 components, of type T.
type Vec4[T number] [4]T

// Vec4CopyVec4 copies the content of src to dst.
func Vec4CopyVec4[T1, T2 number](dst *Vec4[T1], src *Vec4[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
	dst[2] = T1(src[2])
	dst[3] = T1(src[3])
}

// Vec4CopyVec2 copies the content of src to dst.
func Vec4CopyVec2[T1, T2 number](dst *Vec4[T1], src *Vec2[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Vec4CopyVec3 copies the content of src to dst.
func Vec4CopyVec3[T1, T2 number](dst *Vec4[T1], src *Vec3[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
	dst[2] = T1(src[2])
}

// String returns a string representation of the vector.
func (v1 *Vec4[T]) String() string {
	return fmt.Sprintf("{%v, %v, %v, %v}", v1[0], v1[1], v1[2], v1[3])
}

// Inverse sets v1 to the inverse of v2.
// v1 = -v2
func (v1 *Vec4[T]) Inverse(v2 *Vec4[T]) {
	v1[0] = -v2[0]
	v1[1] = -v2[1]
	v1[2] = -v2[2]
	v1[3] = -v2[3]
}

// Add adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec4[T]) Add(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
	v1[2] = v2[2] + v3[2]
	v1[3] = v2[3] + v3[3]
}

// Sub subtracts v2 from v3 component-wise, and stores the result in v1.
// v1 = v2 - v3
func (v1 *Vec4[T]) Sub(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
	v1[2] = v2[2] - v3[2]
	v1[3] = v2[3] - v3[3]
}

// Mul multiplies v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 * v3
func (v1 *Vec4[T]) Mul(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[0] * v3[0]
	v1[1] = v2[1] * v3[1]
	v1[2] = v2[2] * v3[2]
	v1[3] = v2[3] * v3[3]
}

// Div divides v2 by v3 component-wise, and stores the result in v1.
// v1 = v2 / v3
func (v1 *Vec4[T]) Div(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[0] / v3[0]
	v1[1] = v2[1] / v3[1]
	v1[2] = v2[2] / v3[2]
	v1[3] = v2[3] / v3[3]
}

// AddS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 + s
func (v1 *Vec4[T]) AddS(v2 *Vec4[T], s T) {
	v1[0] = v2[0] + s
	v1[1] = v2[1] + s
	v1[2] = v2[2] + s
	v1[3] = v2[3] + s
}

// SubS subtracts each component of v2 by s, and stores the result in v1.
// v1 = v2 - s
func (v1 *Vec4[T]) SubS(v2 *Vec4[T], s T) {
	v1[0] = v2[0] - s
	v1[1] = v2[1] - s
	v1[2] = v2[2] - s
	v1[3] = v2[3] - s
}

// MulS multiplies each component of v2 with s, and stores the result in v1.
// v1 = v2 * s
func (v1 *Vec4[T]) MulS(v2 *Vec4[T], s T) {
	v1[0] = v2[0] * s
	v1[1] = v2[1] * s
	v1[2] = v2[2] * s
	v1[3] = v2[3] * s
}

// DivS divides each component of v2 by s, and stores the result in v1.
// v1 = v2 / s
func (v1 *Vec4[T]) DivS(v2 *Vec4[T], s T) {
	v1[0] = v2[0] / s
	v1[1] = v2[1] / s
	v1[2] = v2[2] / s
	v1[3] = v2[3] / s
}

// AddHomog adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec4[T]) AddHomog(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
	v1[2] = v2[2] + v3[2]
}

// CrossHomog takes the 3D cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec4[T]) CrossHomog(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0], v1[1], v1[2] = v2[1]*v3[2]-v3[1]*v2[2], v2[2]*v3[0]-v3[2]*v2[0], v2[0]*v3[1]-v3[0]*v2[1]
}

// CrossHomogFast takes the 3D cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec4[T]) CrossHomogFast(v2 *Vec4[T], v3 *Vec4[T]) {
	v1[0] = v2[1]*v3[2] - v3[1]*v2[2]
	v1[1] = v2[2]*v3[0] - v3[2]*v2[0]
	v1[2] = v2[0]*v3[1] - v3[0]*v2[1]
}

// Vec4Dot takes the dot product of v1 and v2, and returns the result.
func Vec4Dot[T number](v1 *Vec4[T], v2 *Vec4[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

// Vec4DotHomog takes the dot product of v1 and v2, and returns the result.
func Vec4DotHomog[T number](v1 *Vec4[T], v2 *Vec4[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
