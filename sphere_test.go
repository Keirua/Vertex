package main

import "testing"

type raySphereIntersectionTestData struct {
	ray             Ray
	sphere          Sphere
	hasIntersection bool
}

var raySphereIntersectionTestValues = []raySphereIntersectionTestData{
	{Ray{Vertex{0, 0, 0}, Vertex{1, 0, 0}}, Sphere{Vertex{0, 0, 0}, 4}, true},
	{Ray{Vertex{0, 0, 0}, Vertex{1, 0, 0}}, Sphere{Vertex{100, 0, 0}, 1}, true},
	{Ray{Vertex{0, 0, 0}, Vertex{1, 0, 0}}, Sphere{Vertex{0, 0, 1000}, 4}, false},
}

func TestIntersect(t *testing.T) {
	for _, testValue := range raySphereIntersectionTestValues {
		var intersectResult = testValue.sphere.Intersect(testValue.ray)
		if intersectResult != testValue.hasIntersection {
			t.Error(
				"For", testValue,
				"expected", testValue.hasIntersection,
				"got", intersectResult,
			)
		}
	}
}
