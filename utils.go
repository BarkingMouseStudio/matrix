package matrix

import (
	"math/rand"
)

// Contains functions which return new matrices (without modification to the
// original) or operate on multiple matrices.

// Constructs a new matrix given its elements and dimensions
func New(elements []float64, rows, cols int) (*Matrix, error) {
	if len(elements) != rows*cols {
		return &Matrix{}, DimensionsError
	}
	return &Matrix{elements, rows, cols}, nil
}

// Constructs a new matrix and initialize it with normalized random values
func NewRandNorm(stdDev, mean float64, rows, cols int) *Matrix {
	elements := make([]float64, rows*cols)
	for i := range elements {
		elements[i] = (rand.NormFloat64() * stdDev) + mean
	}
	return &Matrix{elements, rows, cols}
}

func NewRand(rows, cols int) *Matrix {
	elements := make([]float64, cols*rows)
	for i := range elements {
		elements[i] = rand.Float64()
	}
	return &Matrix{elements, rows, cols}
}

func NewOnes(rows, cols int) *Matrix {
	elements := make([]float64, rows*cols)
	for i := range elements {
		elements[i] = 1
	}
	return &Matrix{elements, rows, cols}
}

func NewZeros(rows, cols int) *Matrix {
	return &Matrix{make([]float64, rows*cols), rows, cols}
}

func Reshape(m *Matrix, rows, cols int) (*Matrix, error) {
	size := rows * cols
	if size != m.Size() {
		return &Matrix{}, DimensionsError
	}
	elements := make([]float64, rows*cols)
	copy(elements, m.elements)
	return &Matrix{elements, rows, cols}, nil
}

type F func(float64) float64

func Map(m *Matrix, fn F) *Matrix {
	n := NewZeros(m.rows, m.cols)
	for i, x := range m.elements {
		n.elements[i] = fn(x)
	}
	return n
}

func Transpose(m *Matrix) *Matrix {
	n := NewZeros(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			n.elements[j*n.cols+i] = m.elements[i*m.cols+j]
		}
	}
	return n
}

func Dot(a *Matrix, b *Matrix) (*Matrix, error) {
	if a.cols != b.rows {
		return &Matrix{}, DimensionsError
	}
	c := NewZeros(a.rows, b.cols)
	for i := 0; i < a.rows; i++ {
		for j := 0; j < b.cols; j++ {
			sum := float64(0)
			for k := 0; k < a.cols; k++ {
				sum += a.elements[i*a.cols+k] * b.elements[k*b.cols+j]
			}
			c.elements[i*c.cols+j] = sum
		}
	}
	return c, nil
}

func Add(m *Matrix, n *Matrix) (*Matrix, error) {
	k := m.Copy()
	err := k.Add(n)
	if err != nil {
		return &Matrix{}, err
	}
	return k, nil
}

func Sub(m *Matrix, n *Matrix) (*Matrix, error) {
	k := m.Copy()
	err := k.Sub(n)
	if err != nil {
		return &Matrix{}, err
	}
	return k, nil
}

func MultiplyScalar(m *Matrix, x float64) *Matrix {
	n := m.Copy()
	n.MultiplyScalar(x)
	return n
}

func DivideScalar(m *Matrix, x float64) *Matrix {
	n := m.Copy()
	n.DivideScalar(x)
	return n
}

func Greater(m *Matrix, n *Matrix) *Matrix {
	var rows, cols int
	if m.rows < n.rows {
		rows = m.rows
	} else {
		rows = n.rows
	}
	if m.cols < n.cols {
		cols = m.cols
	} else {
		cols = n.cols
	}
	p := NewZeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			k := i*cols + j
			mk := m.elements[k]
			nk := n.elements[k]
			if mk > nk {
				p.elements[k] = 1
			} else {
				p.elements[k] = 0
			}
		}
	}
	return p
}

func Augment(m *Matrix, n *Matrix) (*Matrix, error) {
	if m.rows != n.rows {
		return nil, DimensionsError
	}
	rows := m.rows
	cols := m.cols + n.cols
	elements := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j < m.cols {
				elements[i*cols+j] = m.elements[i*m.cols+j]
			} else {
				elements[i*cols+j] = n.elements[i*n.cols+(j-m.cols)]
			}
		}
	}
	return &Matrix{elements, rows, cols}, nil
}

func Stack(m *Matrix, n *Matrix) (*Matrix, error) {
	if m.cols != n.cols {
		return nil, DimensionsError
	}
	rows := m.rows + n.rows
	cols := m.cols
	elements := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i < m.rows {
				elements[i*cols+j] = m.elements[i*m.cols+j]
			} else {
				elements[i*cols+j] = n.elements[(i-m.rows)*n.cols+j]
			}
		}
	}
	return &Matrix{elements, rows, cols}, nil
}
