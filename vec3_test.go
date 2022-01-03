package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestVec3CopyVec3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[uint]{1, 2, 3}
	expected := Vec3[int]{1, 2, 3}

	Vec3CopyVec3(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3CopyVec2(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec2[uint]{1, 2}
	expected := Vec3[int]{1, 2, 0}

	Vec3CopyVec2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3CopyVec4(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec4[uint]{1, 2, 3, 4}
	expected := Vec3[int]{1, 2, 3}

	Vec3CopyVec4(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3String(t *testing.T) {
	v1 := Vec3[int]{1, 2, 3}
	expected := "{1, 2, 3}"

	result := v1.String()
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec3Len(t *testing.T) {
	v1 := Vec3[int]{2, 3, 6}
	expected := float64(7)

	result := v1.Len()
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec3Normalize(t *testing.T) {
	v1 := Vec3[float64]{}
	v2 := Vec3[float64]{0, 3, 0}
	expected := Vec3[float64]{0, 1, 0}

	v1.Normalize(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Inverse(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	expected := Vec3[int]{-1, -2, -3}

	v1.Inverse(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Add(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{5, 7, 9}

	v1.Add(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Sub(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{-3, -3, -3}

	v1.Sub(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Mul(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{4, 10, 18}

	v1.Mul(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Div(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{4, 5, 6}
	v3 := Vec3[int]{2, 1, 3}
	expected := Vec3[int]{2, 5, 2}

	v1.Div(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3AddS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{5, 6, 7}

	v1.AddS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3SubS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{-3, -2, -1}

	v1.SubS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3MulS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{4, 8, 12}

	v1.MulS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3DivS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 4, 8}
	s := int(2)
	expected := Vec3[int]{1, 2, 4}

	v1.DivS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Cross(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 1, -1}
	v3 := Vec3[int]{-3, 4, 1}
	expected := Vec3[int]{5, 1, 11}

	v1.Cross(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3CrossFast(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 1, -1}
	v3 := Vec3[int]{-3, 4, 1}
	expected := Vec3[int]{5, 1, 11}

	v1.CrossFast(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Dot(t *testing.T) {
	v1 := Vec3[int]{1, 2, 3}
	v2 := Vec3[int]{4, 5, 6}
	expected := int(32)

	result := v1.Dot(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}
