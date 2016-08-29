package main

import "testing"

type Testable interface {
	PerformTest() bool
}

type scaleTestData struct {
	initialValue Vertex
	scaledValue  Vertex
}

type lengthSquaredTestData struct {
	initialValue Vertex
	sqLength     float64
}

type normalizedTestData struct {
	initialValue    Vertex
	normalizedValue Vertex
}

func (test *scaleTestData) PerformTest() bool {
	var scaledTestValue = test.initialValue.Scale(2)
	return areVerticesEqual(*scaledTestValue, test.initialValue)
}

var scaleTestValues = []scaleTestData{
	{Vertex{0, 0, 0}, Vertex{0, 0, 0}},
	{Vertex{1, 0, 0}, Vertex{1, 0, 0}},
	{Vertex{2, 0, 0}, Vertex{4, 0, 0}},
	{Vertex{3, 0, 0}, Vertex{6, 0, 0}},
	{Vertex{3, 3, 0}, Vertex{6, 6, 0}},
}

var lengthSqTestValues = []lengthSquaredTestData{
	{Vertex{0, 0, 0}, 0},
	{Vertex{1, 0, 0}, 1},
	{Vertex{2, 0, 0}, 4},
	{Vertex{3, 0, 0}, 9},
	{Vertex{3, 3, 0}, 18},
}

var normalizeTestValues = []normalizedTestData{
	{Vertex{1, 0, 0}, Vertex{1, 0, 0}},
	{Vertex{2, 0, 0}, Vertex{1, 0, 0}},
	{Vertex{0, 2, 0}, Vertex{0, 1, 0}},
	{Vertex{0, 0, 4}, Vertex{0, 0, 1}},
}

func areVerticesEqual(v1 Vertex, v2 Vertex) bool {
	return v1.X == v2.X && v2.Y == v2.Y && v1.Z == v2.Z
}

func GenericTest(t *testing.T, testValues []Testable) {
	for _, testValue := range testValues {
		var success = testValue.PerformTest()
		if !success {
			t.Error("Error for", testValue)
		}
	}
}

func TestScale(t *testing.T) {
	//GenericTest(t, scaleTestValues)
	for _, testValue := range scaleTestValues {
		var scaledTestValue = testValue.initialValue.Scale(2)
		if !areVerticesEqual(*scaledTestValue, testValue.initialValue) {
			t.Error(
				"For", testValue.initialValue,
				"expected", testValue.scaledValue,
				"got", scaledTestValue,
			)
		}
	}
}

func TestLengthSq(t *testing.T) {
	for _, testValue := range lengthSqTestValues {
		var lengthSqTestValue = testValue.initialValue.LengthSq()
		if lengthSqTestValue != testValue.sqLength {
			t.Error(
				"For", testValue.initialValue,
				"expected", testValue.sqLength,
				"got", lengthSqTestValue,
			)
		}
	}
}

func TestNormalize(t *testing.T) {
	for _, testValue := range normalizeTestValues {
		var normalized = testValue.initialValue.Normalize()
		if !areVerticesEqual(*normalized, testValue.normalizedValue) {
			t.Error(
				"For", testValue.initialValue,
				"expected", testValue.normalizedValue,
				"got", normalized,
			)
		}
	}
}
