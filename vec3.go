package gogm

// Vec3 is a vector with 3 components, of type T.
type Vec3[T number] [3]T

// V3CopyV3 copies the content of src to dst.
func V3CopyV3[T1, T2 number](dst *Vec3[T1], src *Vec3[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
	dst[2] = T1(src[2])
}

// V3CopyV2 copies the content of src to dst.
func V3CopyV2[T1, T2 number](dst *Vec3[T1], src *Vec2[T2]) {
	dst[0] = T1(src[0])
	dst[1] = T1(src[1])
}

// Invert sets v1 to the inverse of v2.
// v1 = -v2
func (v1 *Vec3[T]) Inverse(v2 *Vec3[T]) {
	v1[0] = -v2[0]
	v1[1] = -v2[1]
	v1[2] = -v2[2]
}

// V3AddV3 adds v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 + v3
func (v1 *Vec3[T]) V3AddV3(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
	v1[2] = v2[2] + v3[2]
}

// V3SubV3 subtracts v2 from v3 component-wise, and stores the result in v1.
// v1 = v2 - v3
func (v1 *Vec3[T]) V3SubV3(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
	v1[2] = v2[2] - v3[2]
}

// V3MulV3 multiplies v2 with v3 component-wise, and stores the result in v1.
// v1 = v2 * v3
func (v1 *Vec3[T]) V3MulV3(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] * v3[0]
	v1[1] = v2[1] * v3[1]
	v1[2] = v2[2] * v3[2]
}

// V3DivV3 divides v2 by v3 component-wise, and stores the result in v1.
// v1 = v2 / v3
func (v1 *Vec3[T]) V3DivV3(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[0] / v3[0]
	v1[1] = v2[1] / v3[1]
	v1[2] = v2[2] / v3[2]
}

// V3AddS adds each component of v2 with s, and stores the result in v1.
// v1 = v2 + s
func (v1 *Vec3[T]) V3AddS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] + s
	v1[1] = v2[1] + s
	v1[2] = v2[2] + s
}

// V3SubS subtracts each component of v2 by s, and stores the result in v1.
// v1 = v2 - s
func (v1 *Vec3[T]) V3SubS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] - s
	v1[1] = v2[1] - s
	v1[2] = v2[2] - s
}

// V3MulS multiplies each component of v2 with s, and stores the result in v1.
// v1 = v2 * s
func (v1 *Vec3[T]) V3MulS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] * s
	v1[1] = v2[1] * s
	v1[2] = v2[2] * s
}

// V3DivS divides each component of v2 by s, and stores the result in v1.
// v1 = v2 / s
func (v1 *Vec3[T]) V3DivS(v2 *Vec3[T], s T) {
	v1[0] = v2[0] / s
	v1[1] = v2[1] / s
	v1[2] = v2[2] / s
}

// V3CrossV3 takes the cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec3[T]) V3CrossV3(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0], v1[1], v1[2] = v2[1]*v3[2]-v3[1]*v2[2], v2[2]*v3[0]-v3[2]*v2[0], v2[0]*v3[1]-v3[0]*v2[1]
}

// V3CrossV3Fast takes the cross product of v1 and v2, and stores the result in v1.
// v1 = v2 x v3
func (v1 *Vec3[T]) V3CrossV3Fast(v2 *Vec3[T], v3 *Vec3[T]) {
	v1[0] = v2[1]*v3[2] - v3[1]*v2[2]
	v1[1] = v2[2]*v3[0] - v3[2]*v2[0]
	v1[2] = v2[0]*v3[1] - v3[0]*v2[1]
}

// V3DotV3 takes the dot product of v1 and v2, and returns the result.
func V3DotV3[T number](v1 *Vec3[T], v2 *Vec3[T]) T {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
