package v

import (
	"math"
	"testing"
)

func TestVec_Add(t *testing.T) {
	v1 := Vec{1, 2}
	v2 := Vec{3, 4}
	expected := Vec{4, 6}
	result := v1.Add(v2)
	if !result.Equals(expected) {
		t.Errorf("Add failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Sub(t *testing.T) {
	v1 := Vec{5, 7}
	v2 := Vec{2, 3}
	expected := Vec{3, 4}
	result := v1.Sub(v2)
	if !result.Equals(expected) {
		t.Errorf("Sub failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Div(t *testing.T) {
	v1 := Vec{10, 20}
	v2 := Vec{2, 4}
	expected := Vec{5, 5}
	result := v1.Div(v2)
	if !result.Equals(expected) {
		t.Errorf("Div failed: expected %v, got %v", expected, result)
	}
}

func TestVec_DivS(t *testing.T) {
	v := Vec{10, 20}
	s := 2.0
	expected := Vec{5, 10}
	result := v.DivS(s)
	if !result.Equals(expected) {
		t.Errorf("DivS failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Mul(t *testing.T) {
	v1 := Vec{2, 3}
	v2 := Vec{4, 5}
	expected := Vec{8, 15}
	result := v1.Mul(v2)
	if !result.Equals(expected) {
		t.Errorf("Mul failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Scale(t *testing.T) {
	v := Vec{2, 3}
	s := 2.5
	expected := Vec{5, 7.5}
	result := v.Scale(s)
	if !result.Equals(expected) {
		t.Errorf("Scale failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Unit(t *testing.T) {
	tests := []struct {
		name     string
		v        Vec
		expected Vec
	}{
		{"normal vector", Vec{3, 4}, Vec{0.6, 0.8}},
		{"zero vector", Vec{0, 0}, Vec{0, 0}},
		{"already unit", Vec{1, 0}, Vec{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.v.Unit()
			if !result.EqualsPr(tt.expected, 1e-9) {
				t.Errorf("Unit failed: expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVec_Abs(t *testing.T) {
	v := Vec{-2.5, 3.7}
	expected := Vec{2.5, 3.7}
	result := v.Abs()
	if !result.Equals(expected) {
		t.Errorf("Abs failed: expected %v, got %v", expected, result)
	}
}

func TestVec_AbsX(t *testing.T) {
	v := Vec{-2.5, 3.7}
	expected := 2.5
	result := v.AbsX()
	if result != expected {
		t.Errorf("AbsX failed: expected %v, got %v", expected, result)
	}
}

func TestVec_AbsY(t *testing.T) {
	v := Vec{-2.5, 3.7}
	expected := 3.7
	result := v.AbsY()
	if result != expected {
		t.Errorf("AbsY failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Neg(t *testing.T) {
	v := Vec{2, -3}
	expected := Vec{-2, 3}
	result := v.Neg()
	if !result.Equals(expected) {
		t.Errorf("Neg failed: expected %v, got %v", expected, result)
	}
}

func TestVec_NegX(t *testing.T) {
	v := Vec{2, -3}
	expected := Vec{-2, -3}
	result := v.NegX()
	if !result.Equals(expected) {
		t.Errorf("NegX failed: expected %v, got %v", expected, result)
	}
}

func TestVec_NegY(t *testing.T) {
	v := Vec{2, -3}
	expected := Vec{2, 3}
	result := v.NegY()
	if !result.Equals(expected) {
		t.Errorf("NegY failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Dot(t *testing.T) {
	v1 := Vec{1, 2}
	v2 := Vec{3, 4}
	expected := 11.0
	result := v1.Dot(v2)
	if result != expected {
		t.Errorf("Dot failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Cross(t *testing.T) {
	v1 := Vec{1, 2}
	v2 := Vec{3, 4}
	expected := -2.0
	result := v1.Cross(v2)
	if result != expected {
		t.Errorf("Cross failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Project(t *testing.T) {
	v := Vec{3, 4}
	other := Vec{1, 0}
	expected := Vec{3, 0}
	result := v.Project(other)
	if !result.Equals(expected) {
		t.Errorf("Project failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Angle(t *testing.T) {
	v := Vec{1, 0}
	expected := 0.0
	result := v.Angle()
	if result != expected {
		t.Errorf("Angle failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Rotate(t *testing.T) {
	v := Vec{1, 0}
	angle := math.Pi / 2
	expected := Vec{0, 1}
	result := v.Rotate(angle)
	if !result.EqualsPr(expected, 1e-9) {
		t.Errorf("Rotate failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Mag(t *testing.T) {
	v := Vec{3, 4}
	expected := 5.0
	result := v.Mag()
	if result != expected {
		t.Errorf("Mag failed: expected %v, got %v", expected, result)
	}
}

func TestVec_SetMag(t *testing.T) {
	v := Vec{3, 4}
	m := 10.0
	expected := Vec{6, 8}
	result := v.SetMag(m)
	if !result.Equals(expected) {
		t.Errorf("SetMag failed: expected %v, got %v", expected, result)
	}
}

func TestVec_SetMag_ZeroVector(t *testing.T) {
	v := Vec{0, 0}
	result := v.SetMag(5.0)
	if !result.Equals(v) {
		t.Errorf("SetMag with zero vector failed: expected %v, got %v", v, result)
	}
}

func TestVec_MagSq(t *testing.T) {
	v := Vec{3, 4}
	expected := 25.0
	result := v.MagSq()
	if result != expected {
		t.Errorf("MagSq failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Slerp(t *testing.T) {
	tests := []struct {
		name     string
		v        Vec
		to       Vec
		weight   float64
		expected Vec
	}{
		{
			name:     "normal slerp",
			v:        Vec{1, 0},
			to:       Vec{0, 1},
			weight:   0.5,
			expected: Vec{0.7071067811865475, 0.7071067811865475},
		},
		{
			name:     "start vector is zero - falls back to lerp",
			v:        Vec{0, 0},
			to:       Vec{10, 10},
			weight:   0.5,
			expected: Vec{5, 5},
		},
		{
			name:     "end vector is zero - falls back to lerp",
			v:        Vec{10, 10},
			to:       Vec{0, 0},
			weight:   0.5,
			expected: Vec{5, 5},
		},
		{
			name:     "both vectors are zero - falls back to lerp",
			v:        Vec{0, 0},
			to:       Vec{0, 0},
			weight:   0.5,
			expected: Vec{0, 0},
		},
		{
			name:     "weight at 0",
			v:        Vec{1, 0},
			to:       Vec{0, 1},
			weight:   0,
			expected: Vec{1, 0},
		},
		{
			name:     "weight at 1",
			v:        Vec{1, 0},
			to:       Vec{0, 1},
			weight:   1,
			expected: Vec{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.v.Slerp(tt.to, tt.weight)
			if !result.EqualsPr(tt.expected, 1e-9) {
				t.Errorf("Slerp failed: expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVec_AngleTo(t *testing.T) {
	v1 := Vec{1, 0}
	v2 := Vec{0, 1}
	expected := math.Pi / 2
	result := v1.AngleTo(v2)
	if result != expected {
		t.Errorf("AngleTo failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Limit(t *testing.T) {
	v := Vec{3, 4}
	max := 3.0
	expected := Vec{1.8, 2.4}
	result := v.Limit(max)
	if !result.EqualsPr(expected, 1e-9) {
		t.Errorf("Limit failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Limit_NoChange(t *testing.T) {
	v := Vec{1, 1}
	max := 10.0
	result := v.Limit(max)
	if !result.Equals(v) {
		t.Errorf("Limit with max > magnitude failed: expected %v, got %v", v, result)
	}
}

func TestVec_Lerp(t *testing.T) {
	v1 := Vec{0, 0}
	v2 := Vec{10, 10}
	tVal := 0.5
	expected := Vec{5, 5}
	result := v1.Lerp(v2, tVal)
	if !result.Equals(expected) {
		t.Errorf("Lerp failed: expected %v, got %v", expected, result)
	}
}

func TestVec_IsZero(t *testing.T) {
	tests := []struct {
		name     string
		v        Vec
		expected bool
	}{
		{"zero vector", Vec{0, 0}, true},
		{"non-zero vector", Vec{1, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.v.IsZero()
			if result != tt.expected {
				t.Errorf("IsZero failed: expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVec_Dist(t *testing.T) {
	v1 := Vec{0, 0}
	v2 := Vec{3, 4}
	expected := 5.0
	result := v1.Dist(v2)
	if result != expected {
		t.Errorf("Dist failed: expected %v, got %v", expected, result)
	}
}

func TestVec_DistSq(t *testing.T) {
	v1 := Vec{0, 0}
	v2 := Vec{3, 4}
	expected := 25.0
	result := v1.DistSq(v2)
	if result != expected {
		t.Errorf("DistSq failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Round(t *testing.T) {
	v := Vec{2.3, 2.7}
	expected := Vec{2, 3}
	result := v.Round()
	if !result.Equals(expected) {
		t.Errorf("Round failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Floor(t *testing.T) {
	v := Vec{2.3, -2.7}
	expected := Vec{2, -3}
	result := v.Floor()
	if !result.Equals(expected) {
		t.Errorf("Floor failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Ceil(t *testing.T) {
	v := Vec{2.3, -2.7}
	expected := Vec{3, -2}
	result := v.Ceil()
	if !result.Equals(expected) {
		t.Errorf("Ceil failed: expected %v, got %v", expected, result)
	}
}

func TestFromAngle(t *testing.T) {
	angle := math.Pi / 2
	expected := Vec{0, 1}
	result := FromAngle(angle)
	if !result.EqualsPr(expected, 1e-9) {
		t.Errorf("FromAngle failed: expected %v, got %v", expected, result)
	}
}

func TestVec_EqualsPr(t *testing.T) {
	v1 := Vec{1.0000001, 2}
	v2 := Vec{1.0000002, 2}
	delta := 1e-6
	expected := true
	result := v1.EqualsPr(v2, delta)
	if result != expected {
		t.Errorf("EqualsPr failed: expected %v, got %v", expected, result)
	}
}

func TestVec_Equals(t *testing.T) {
	v1 := Vec{1, 2}
	v2 := Vec{1, 2}
	v3 := Vec{3, 4}
	if !v1.Equals(v2) {
		t.Errorf("Equals failed: expected true, got false")
	}
	if v1.Equals(v3) {
		t.Errorf("Equals failed: expected false, got true")
	}
}

func TestVec_Reflect(t *testing.T) {
	v := Vec{1, -1}
	normal := Vec{0, 1}
	expected := Vec{1, 1}
	result := v.Reflect(normal)
	if !result.Equals(expected) {
		t.Errorf("Reflect failed: expected %v, got %v", expected, result)
	}
}

func TestVec_String(t *testing.T) {
	v := Vec{1.5, 2.3}
	expected := "(1.5, 2.3)"
	result := v.String()
	if result != expected {
		t.Errorf("String failed: expected %v, got %v", expected, result)
	}
}

func TestConstants(t *testing.T) {
	if !One.Equals(Vec{1, 1}) {
		t.Errorf("One constant is incorrect")
	}
	if !Left.Equals(Vec{-1, 0}) {
		t.Errorf("Left constant is incorrect")
	}
	if !Right.Equals(Vec{1, 0}) {
		t.Errorf("Right constant is incorrect")
	}
	if !Up.Equals(Vec{0, -1}) {
		t.Errorf("Up constant is incorrect")
	}
	if !Down.Equals(Vec{0, 1}) {
		t.Errorf("Down constant is incorrect")
	}
}
