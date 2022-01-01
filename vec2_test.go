package gogm

import (
	"testing"

	"github.com/smyrman/subx"
)

func TestV2CopyV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[uint]{1, 2}
	V2CopyV2(&v1, &v2)
	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{1, 2})))
}

func TestVec2Inverse(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v1.Inverse(&v2)

	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{-1, -2})))
}

func TestV2AddV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	v1.V2AddV2(&v2, &v3)

	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{4, 6})))
}

func TestV2SubV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	v1.V2SubV2(&v2, &v3)

	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{-2, -2})))
}

func TestV2MulV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{1, 2}
	v3 := Vec2[int]{3, 4}
	v1.V2MulV2(&v2, &v3)

	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{3, 8})))
}

func TestV2DivV2(t *testing.T) {
	v1 := Vec2[int]{}
	v2 := Vec2[int]{3, 4}
	v3 := Vec2[int]{1, 2}
	v1.V2DivV2(&v2, &v3)

	t.Run("Expect correct result", subx.Test(func() Vec2[int] { return v1 }, subx.CompareEqual(Vec2[int]{3, 2})))
}
