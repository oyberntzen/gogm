package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestV3CopyV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[uint]{1, 2, 3}
	expected := Vec3[int]{1, 2, 3}

	V3CopyV3(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec3Inverse(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	expected := Vec3[int]{-1, -2, -3}

	v1.Inverse(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3AddV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{5, 7, 9}

	v1.V3AddV3(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3SubV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{-3, -3, -3}

	v1.V3SubV3(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3MulV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	v3 := Vec3[int]{4, 5, 6}
	expected := Vec3[int]{4, 10, 18}

	v1.V3MulV3(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3DivV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{4, 5, 6}
	v3 := Vec3[int]{2, 1, 3}
	expected := Vec3[int]{2, 5, 2}

	v1.V3DivV3(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3AddS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{5, 6, 7}

	v1.V3AddS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3SubS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{-3, -2, -1}

	v1.V3SubS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3MulS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{1, 2, 3}
	s := int(4)
	expected := Vec3[int]{4, 8, 12}

	v1.V3MulS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3DivS(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 4, 8}
	s := int(2)
	expected := Vec3[int]{1, 2, 4}

	v1.V3DivS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3CrossV3(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 1, -1}
	v3 := Vec3[int]{-3, 4, 1}
	expected := Vec3[int]{5, 1, 11}

	v1.V3CrossV3(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3CrossV3Fast(t *testing.T) {
	v1 := Vec3[int]{}
	v2 := Vec3[int]{2, 1, -1}
	v3 := Vec3[int]{-3, 4, 1}
	expected := Vec3[int]{5, 1, 11}

	v1.V3CrossV3Fast(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV3DotV3(t *testing.T) {
	v1 := Vec3[int]{1, 2, 3}
	v2 := Vec3[int]{4, 5, 6}
	expected := int(32)

	result := V3DotV3(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}
