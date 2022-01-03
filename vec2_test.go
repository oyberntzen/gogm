package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestVec2CopyVec2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[uint]{1, 2}
	expected := Vec2[int]{1, 2}

	Vec2CopyVec2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2CopyVec3(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec3[uint]{1, 2, 3}
	expected := Vec2[int]{1, 2}

	Vec2CopyVec3(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2CopyVec4(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec4[uint]{1, 2, 3, 4}
	expected := Vec2[int]{1, 2}

	Vec2CopyVec4(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2String(t *testing.T) {
	v1 := Vec2[int]{1, 2}
	expected := "{1, 2}"

	result := v1.String()
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec2Inverse(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	expected := Vec2[int]{-1, -2}

	v1.Inverse(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Add(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{4, 6}

	v1.Add(&v2, &v3)

	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Sub(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{-2, -2}

	v1.Sub(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Mul(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{3, 8}

	v1.Mul(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Div(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{3, 4}
	v3 := Vec2[int]{1, 2}
	expected := Vec2[int]{3, 2}

	v1.Div(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2AddS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{4, 5}

	v1.AddS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2SubS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{-2, -1}

	v1.SubS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2MulS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{3, 6}

	v1.MulS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2DivS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{3, 6}
	s := int(3)
	expected := Vec2[int]{1, 2}

	v1.DivS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Cross(t *testing.T) {
	v1 := Vec2[int]{-3, 1}
	v2 := Vec2[int]{2, 1}
	expected := int(-5)

	result := Vec2Cross(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestVec2Dot(t *testing.T) {
	v1 := Vec2[int]{1, 2}
	v2 := Vec2[int]{3, 4}
	expected := int(11)

	result := Vec2Dot(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}
