package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	M [16]float64
}

func NewIdentityMatrix() *Matrix {
	var m Matrix
	m.M[0]  = 1;
	m.M[5]  = 1;
	m.M[10] = 1;
	m.M[15] = 1;
	return &m
}


func NewScaleMatrix(v Vertex) *Matrix {
	var m Matrix
	m.M[0]  = v.X;
	m.M[5]  = v.Y;
	m.M[10] = v.Z;
	m.M[15] = 1;
	return &m
}

func NewTranslationMatrix(v Vertex) *Matrix {
	var m *Matrix = NewIdentityMatrix()
	m.M[3]  = v.X;
	m.M[7]  = v.Y;
	m.M[11] = v.Z;
	return m
}

func NewRotationAroundXMatrix(angle float64) *Matrix {
	var m *Matrix = NewIdentityMatrix()

	m.M[5] = math.Cos(angle)
	m.M[6] = -math.Sin(angle)
	m.M[9] = math.Sin(angle)
	m.M[10] = math.Cos(angle)

	return m
}

func NewRotationAroundYMatrix(angle float64) *Matrix {
	var m *Matrix = NewIdentityMatrix()

	m.M[0] = math.Cos(angle)
	m.M[2] = math.Sin(angle)
	m.M[8] = -math.Sin(angle)
	m.M[10] = math.Cos(angle)

	return m
}

func NewRotationAroundZMatrix(angle float64) *Matrix {
	var m *Matrix = NewIdentityMatrix()

	m.M[0] = math.Cos(angle)
	m.M[1] = -math.Sin(angle)
	m.M[4] = math.Sin(angle)
	m.M[5] = math.Cos(angle)

	return m
}

func (m *Matrix) MulV(v Vertex) Vertex{
	var out Vertex;
	out.X = v.X*m.M[0] + v.Y*m.M[1] + v.Z*m.M[2] + m.M[3]
	out.Y = v.X*m.M[4] + v.Y*m.M[5] + v.Z*m.M[6] + m.M[7]
	out.Z = v.X*m.M[8] + v.Y*m.M[9] + v.Z*m.M[10] + m.M[11]

	return out
}


func (m *Matrix) Inverse() Matrix{
	var b Matrix;

	var s0 = m.M[0] * m.M[5] - m.M[4] * m.M[1];
	var s1 = m.M[0] * m.M[6] - m.M[4] * m.M[2];
	var s2 = m.M[0] * m.M[7] - m.M[4] * m.M[3];
	var s3 = m.M[1] * m.M[6] - m.M[5] * m.M[2];
	var s4 = m.M[1] * m.M[7] - m.M[5] * m.M[3];
	var s5 = m.M[2] * m.M[7] - m.M[6] * m.M[3];

	var c5 = m.M[10] * m.M[15] - m.M[14] * m.M[11];
	var c4 = m.M[9] * m.M[15] - m.M[13] * m.M[11];
	var c3 = m.M[9] * m.M[14] - m.M[13] * m.M[10];
	var c2 = m.M[8] * m.M[15] - m.M[12] * m.M[11];
	var c1 = m.M[8] * m.M[14] - m.M[12] * m.M[10];
	var c0 = m.M[8] * m.M[13] - m.M[12] * m.M[9];

	// Should check for 0 determinant
	var invdet = 1.0 / (s0 * c5 - s1 * c4 + s2 * c3 + s3 * c2 - s4 * c1 + s5 * c0);

	b.M[0] = ( m.M[5] * c5 - m.M[6] * c4 + m.M[7] * c3) * invdet;
	b.M[1] = (-m.M[1] * c5 + m.M[2] * c4 - m.M[3] * c3) * invdet;
	b.M[2] = ( m.M[13] * s5 - m.M[14] * s4 + m.M[15] * s3) * invdet;
	b.M[3] = (-m.M[9] * s5 + m.M[10] * s4 - m.M[11] * s3) * invdet;

	b.M[4] = (-m.M[4] * c5 + m.M[6] * c2 - m.M[7] * c1) * invdet;
	b.M[5] = ( m.M[0] * c5 - m.M[2] * c2 + m.M[3] * c1) * invdet;
	b.M[6] = (-m.M[12] * s5 + m.M[14] * s2 - m.M[15] * s1) * invdet;
	b.M[7] = ( m.M[8] * s5 - m.M[10] * s2 + m.M[11] * s1) * invdet;

	b.M[8] = ( m.M[4] * c4 - m.M[5] * c2 + m.M[7] * c0) * invdet;
	b.M[9] = (-m.M[0] * c4 + m.M[1] * c2 - m.M[3] * c0) * invdet;
	b.M[10] = ( m.M[12] * s4 - m.M[13] * s2 + m.M[15] * s0) * invdet;
	b.M[11] = (-m.M[8] * s4 + m.M[9] * s2 - m.M[11] * s0) * invdet;

	b.M[12] = (-m.M[4] * c3 + m.M[5] * c1 - m.M[6] * c0) * invdet;
	b.M[13] = ( m.M[0] * c3 - m.M[1] * c1 + m.M[2] * c0) * invdet;
	b.M[14] = (-m.M[12] * s3 + m.M[13] * s1 - m.M[14] * s0) * invdet;
	b.M[15] = ( m.M[8] * s3 - m.M[9] * s1 + m.M[10] * s0) * invdet;

	return b;
}

func (mat_a *Matrix) MulM(mat_b *Matrix) Matrix {
	var mat_r Matrix;

    for i := 0; i < 16; i+=4 {
        for j := 0; j < 4; j++ {
            mat_r.M[i + j] = (mat_b.M[i] * mat_a.M[j]) + (mat_b.M[i + 1] * mat_a.M[j +  4]) + (mat_b.M[i + 2] * mat_a.M[j +  8]) + (mat_b.M[i + 3] * mat_a.M[j + 12]);
       }
    }

    return mat_r
}

func (m *Matrix) Print() {
	for i := 0; i < 4; i++ {
		fmt.Println( m.M[i*4:(i+1)*4])
	}
}

/*
var v = Vertex{1,2,3}
fmt.Println("v", v)

var m *Matrix = NewIdentityMatrix()
m.Print()
fmt.Println("Identity", m.MulV(v))

var s *Matrix = NewScaleMatrix(Vertex{1,2,3})
s.Print()
fmt.Println("Identity", s.MulV(v))

var t *Matrix = NewTranslationMatrix(Vertex{1,2,3})
t.Print()
fmt.Println("Identity", t.MulV(v))

fmt.Println("Matrix inversions")
var inverseTranslation = t.Inverse()
inverseTranslation.Print()

var inverseScale = s.Inverse()
inverseScale.Print()

fmt.Println("Should be identity")
var i = inverseTranslation.MulM(t)
i.Print()

var is = inverseScale.MulM(s)
is.Print()
*/