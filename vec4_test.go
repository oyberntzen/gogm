package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestVec4CopyVec4(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[uint]{1, 2, 3, 4}
	expected := Vec4[int]{1, 2, 3, 4}

	Vec4CopyVec4(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4CopyVec2(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec2[uint]{1, 2}
	expected := Vec4[int]{1, 2, 0, 0}

	Vec4CopyVec2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4CopyVec3(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec3[uint]{1, 2, 3}
	expected := Vec4[int]{1, 2, 3, 0}

	Vec4CopyVec3(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4String(t *testing.T) {
	v1 := Vec4[int]{1, 2, 3, 4}
	expected := "{1, 2, 3, 4}"

	result := v1.String()
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec4Len(t *testing.T) {
	v1 := Vec4[int]{1, 1, 1, 1}
	expected := float64(2)

	result := v1.Len()
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec4Normalize(t *testing.T) {
	v1 := Vec4[float64]{}
	v2 := Vec4[float64]{0, 3, 0, 0}
	expected := Vec4[float64]{0, 1, 0, 0}

	v1.Normalize(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Inverse(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	expected := Vec4[int]{-1, -2, -3, -4}

	v1.Inverse(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Add(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	v3 := Vec4[int]{5, 6, 7, 8}
	expected := Vec4[int]{6, 8, 10, 12}

	v1.Add(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Sub(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	v3 := Vec4[int]{5, 6, 7, 8}
	expected := Vec4[int]{-4, -4, -4, -4}

	v1.Sub(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Mul(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	v3 := Vec4[int]{5, 6, 7, 8}
	expected := Vec4[int]{5, 12, 21, 32}

	v1.Mul(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Div(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{4, 5, 6, 12}
	v3 := Vec4[int]{2, 1, 3, 4}
	expected := Vec4[int]{2, 5, 2, 3}

	v1.Div(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4AddS(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	s := int(5)
	expected := Vec4[int]{6, 7, 8, 9}

	v1.AddS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4SubS(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	s := int(5)
	expected := Vec4[int]{-4, -3, -2, -1}

	v1.SubS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4MulS(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{1, 2, 3, 4}
	s := int(5)
	expected := Vec4[int]{5, 10, 15, 20}

	v1.MulS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4DivS(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{2, 4, 8, 18}
	s := int(2)
	expected := Vec4[int]{1, 2, 4, 9}

	v1.DivS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4CrossHomog(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{2, 1, -1, 0}
	v3 := Vec4[int]{-3, 4, 1, 0}
	expected := Vec4[int]{5, 1, 11, 0}

	v1.CrossHomog(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4CrossHomogFast(t *testing.T) {
	v1 := Vec4[int]{}
	v2 := Vec4[int]{2, 1, -1, 0}
	v3 := Vec4[int]{-3, 4, 1, 0}
	expected := Vec4[int]{5, 1, 11, 0}

	v1.CrossHomogFast(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec4Dot(t *testing.T) {
	v1 := Vec4[int]{1, 2, 3, 4}
	v2 := Vec4[int]{5, 6, 7, 8}
	expected := int(70)

	result := v1.Dot(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec4DotHomog(t *testing.T) {
	v1 := Vec4[int]{1, 2, 3, 1}
	v2 := Vec4[int]{4, 5, 6, 1}
	expected := int(32)

	result := v1.DotHomog(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}
