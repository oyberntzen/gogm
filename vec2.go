package gogm

import "constraints"

type number interface {
	constraints.Integer | constraints.Float
}

// Vec2 is a vector with 2 components, of type T.
type Vec2[T number] [2]T

// V2CopyV2 copies the content of src to dst.
func V2CopyV2[T1, T2 number](dst *Vec2[T1], src *Vec2[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// V2CopyV3 copies the content of src to dst.
func V2CopyV3[T1, T2 number](dst *Vec2[T1], src *Vec3[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Invert sets v1 to the inverse of v2.
// v1 = -v2
func (v1 *Vec2[T]) Inverse(v2 *Vec2[T]) {
	v1[0] = -v2[0]
	v1[1] = -v2[1]
}

// V2AddV2 adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec2[T]) V2AddV2(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
}

// V2SubV2 subtracts v2 from v3 component-wise, and stores the result in v1.
// v1 = v2 - v3
func (v1 *Vec2[T]) V2SubV2(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
}

// V2MulV2 multiplies v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 * v3
func (v1 *Vec2[T]) V2MulV2(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] * v3[0]
	v1[1] = v2[1] * v3[1]
}

// V2DivV2 divides v2 by v3 component-wise, and stores the result in v1.
// v1 = v2 / v3
func (v1 *Vec2[T]) V2DivV2(v2 *Vec2[T], v3 *Vec2[T]) {
	v1[0] = v2[0] / v3[0]
	v1[1] = v2[1] / v3[1]
}

// V2AddS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 + s
func (v1 *Vec2[T]) V2AddS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] + s
	v1[1] = v2[1] + s
}

// V2SubS subtracts each component of v2 by s, and stores the result in v1.
// v1 = v2 - s
func (v1 *Vec2[T]) V2SubS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] - s
	v1[1] = v2[1] - s
}

// V2MulS multiplies each component of v2 with s, and stores the result in v1.
// v1 = v2 * s
func (v1 *Vec2[T]) V2MulS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] * s
	v1[1] = v2[1] * s
}

// V2DivS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 / s
func (v1 *Vec2[T]) V2DivS(v2 *Vec2[T], s T) {
	v1[0] = v2[0] / s
	v1[1] = v2[1] / s
}

// V2CrossV2 takes the cross product of v1 and v2, and returns the result.
func V2CrossV2[T number](v1 *Vec2[T], v2 *Vec2[T]) T {
	return v1[0]*v2[1] - v1[1]*v2[0]
}

// V2DotV2 takes the dot product of v1 and v2, and returns the result.
func V2DotV2[T number](v1 *Vec2[T], v2 *Vec2[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1]
}
