package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestV2CopyV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[uint]{1, 2}
	expected := Vec2[int]{1, 2}

	V2CopyV2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestVec2Inverse(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	expected := Vec2[int]{-1, -2}

	v1.Inverse(&v2)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2AddV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{4, 6}

	v1.V2AddV2(&v2, &v3)

	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2SubV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{-2, -2}

	v1.V2SubV2(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2MulV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	expected := Vec2[int]{3, 8}

	v1.V2MulV2(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2DivV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{3, 4}
	v3 := Vec2[int]{1, 2}
	expected := Vec2[int]{3, 2}

	v1.V2DivV2(&v2, &v3)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2AddS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{4, 5}

	v1.V2AddS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2SubS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{-2, -1}

	v1.V2SubS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2MulS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	s := int(3)
	expected := Vec2[int]{3, 6}

	v1.V2MulS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2DivS(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{3, 6}
	s := int(3)
	expected := Vec2[int]{1, 2}

	v1.V2DivS(&v2, s)
	t.Run("Expect correct result", subx.Test(subx.Value(v1), subx.CompareEqual(expected)))
}

func TestV2CrossV2(t *testing.T) {
	v1 := Vec2[int]{-3, 1}
	v2 := Vec2[int]{2, 1}
	expected := int(-5)

	result := V2CrossV2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}

func TestV2DotV2(t *testing.T) {
	v1 := Vec2[int]{1, 2}
	v2 := Vec2[int]{3, 4}
	expected := int(11)

	result := V2DotV2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(subx.Value(result), subx.CompareEqual(expected)))
}
