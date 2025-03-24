package vector

import (
	"math"
	"testing"
)

func assertFloat(t *testing.T, name string, expected, actual float64) {
	if actual != expected {
		t.Errorf("Incorrect value for %s, expected %f got %f", name, expected, actual)
	}
}

func assertVector(t *testing.T, expected, actual Vector3) {
	if actual.X() != expected.X() || actual.Y() != expected.Y() || actual.Z() != expected.Z() {
		t.Errorf("Unexpected vector value. Expected %+v but got %+v", expected, actual)
	}
}

func TestNewVector3(t *testing.T) {
	vec := NewVector3(1, 2, 3)

	assertFloat(t, "X", 1.0, vec.X())
	assertFloat(t, "Y", 2.0, vec.Y())
	assertFloat(t, "Z", 3.0, vec.Z())
}

func TestAddVector(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   Vector3
	}{
		{"Should add positive whole vectors", vec3{1.0, 2.0, 3.0}, vec3{1.0, 1.0, 1.0}, vec3{2.0, 3.0, 4.0}},
		{"Should add positive decimal vectors", vec3{1.5, 2.5, 3.5}, vec3{1.5, 1.5, 1.5}, vec3{3.0, 4.0, 5.0}},
		{"Should add negative vectors", vec3{-1.5, -2.5, -3.5}, vec3{1.5, 1.5, 1.5}, vec3{0.0, -1.0, -2.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.AddVector(tt.input2)
			assertVector(t, tt.want, ans)
		})
	}
}

func TestSubtractVector(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   Vector3
	}{
		{"Should subtract positive whole vectors", vec3{1.0, 2.0, 3.0}, vec3{1.0, 1.0, 1.0}, vec3{0.0, 1.0, 2.0}},
		{"Should subtract positive decimal vectors", vec3{1.5, 2.5, 3.5}, vec3{1.5, 1.5, 1.5}, vec3{0.0, 1.0, 2.0}},
		{"Should subtract negative vectors", vec3{-1.5, -2.5, -3.5}, vec3{1.5, 1.5, 1.5}, vec3{-3.0, -4.0, -5.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.SubtractVector(tt.input2)
			assertVector(t, tt.want, ans)
		})
	}
}

func TestMultiplyVector(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   Vector3
	}{
		{"Should multiply positive whole vectors", vec3{1.0, 2.0, 3.0}, vec3{2.0, 2.0, 2.0}, vec3{2.0, 4.0, 6.0}},
		{"Should multiply positive decimal vectors", vec3{1.5, 2.5, 3.5}, vec3{1.5, 1.5, 1.5}, vec3{2.25, 3.75, 5.25}},
		{"Should multiply negative vectors", vec3{-1.5, -2.5, -3.5}, vec3{1.5, 1.5, 1.5}, vec3{-2.25, -3.75, -5.25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.MultiplyVector(tt.input2)
			assertVector(t, tt.want, ans)
		})
	}
}

func TestDivideVector(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   Vector3
	}{
		{"Should divide positive whole vectors", vec3{1.0, 2.0, 3.0}, vec3{2.0, 2.0, 2.0}, vec3{0.5, 1.0, 1.5}},
		{"Should divide positive decimal vectors", vec3{1.5, 2.5, 4.5}, vec3{1.5, 2.0, 1.5}, vec3{1.0, 1.25, 3.0}},
		{"Should divide negative vectors", vec3{-1.5, -2.5, -4.5}, vec3{1.5, 2.0, 1.5}, vec3{-1.0, -1.25, -3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.DivideVector(tt.input2)
			assertVector(t, tt.want, ans)
		})
	}
}

/*
TODO: Add coverage for Float arithmetic functions
*/

func TestLengthSquared(t *testing.T) {
	var tests = []struct {
		name  string
		input Vector3
		want  float64
	}{
		{"Should get length where all sides are 1", vec3{1.0, 1.0, 1.0}, 3.0},
		{"Should get length where all sides are mixed positive whole", vec3{1.0, 2.0, 3.0}, 14.0},
		{"Should get length where all sides are mixed positive decimal", vec3{1.5, 2.2, 3.66}, 20.4856},
		{"Should get length where all sides are mixed negative decimal", vec3{-1.5, -2.2, -3.66}, 20.4856},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.LengthSquared()
			assertFloat(t, "Length squared", tt.want, ans)
		})
	}
}

func TestLength(t *testing.T) {
	var tests = []struct {
		name  string
		input Vector3
		want  float64
	}{
		{"Should get length where all sides are 1", vec3{1.0, 1.0, 1.0}, math.Sqrt(3.0)},
		{"Should get length where all sides are mixed positive whole", vec3{1.0, 2.0, 3.0}, math.Sqrt(14.0)},
		{"Should get length where all sides are mixed positive decimal", vec3{1.5, 2.2, 3.66}, math.Sqrt(20.4856)},
		{"Should get length where all sides are mixed negative decimal", vec3{-1.5, -2.2, -3.66}, math.Sqrt(20.4856)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.Length()
			assertFloat(t, "Length", tt.want, ans)
		})
	}
}

func TestDot(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   float64
	}{
		{"Should dot vectors containing whole numbers", vec3{1.0, 2.0, 3.0}, vec3{4.0, -5.0, 6.0}, 12.0},
		// TODO: Add more scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.Dot(tt.input2)
			assertFloat(t, "Dot", tt.want, ans)
		})
	}
}

func TestCross(t *testing.T) {
	var tests = []struct {
		name   string
		input1 Vector3
		input2 Vector3
		want   Vector3
	}{
		{"Should cross vectors containing whole numbers", vec3{3.0, 4.0, 7.0}, vec3{4.0, 9.0, 2.0}, vec3{-55.0, 22.0, 11.0}},
		{"Should find parallelogram with area 0", vec3{3.0, -3.0, 1.0}, vec3{-12.0, 12.0, -4}, vec3{0.0, 0.0, 0.0}},
		// TODO: Add more scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input1.Cross(tt.input2)
			assertVector(t, tt.want, ans)
		})
	}
}

func TestUnit(t *testing.T) {
	var tests = []struct {
		name  string
		input Vector3
		want  Vector3
	}{
		{"Should find unit vector", vec3{3.0, 4.0, 0.0}, vec3{0.6, 0.8, 0.0}},
		// TODO: Add more scenarios
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.UnitVector()
			assertVector(t, tt.want, ans)
		})
	}
}
