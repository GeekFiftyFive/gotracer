package vector

import "testing"

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
