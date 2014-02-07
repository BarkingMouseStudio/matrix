// Package matrix implements a simple dense matrix.
package matrix

import (
	"errors"
	"fmt"
	"math"
)

// Contains the Matrix definition and common matrix operations.

var DimensionsError = errors.New("Incorrect dimensions")
var IndexError = errors.New("Index out of range")

type Matrix struct {
	elements   []float64
	rows, cols int
}

func (m *Matrix) Arrays() [][]float64 {
	a := make([][]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		a[i] = m.elements[i*m.cols : i*m.cols+m.cols]
	}
	return a
}

func (m *Matrix) Array() []float64 {
	return m.elements
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Size() int {
	return m.rows * m.cols
}

// Implement Stringer interface
func (m *Matrix) String() string {
	if m.rows == 0 || m.cols == 0 {
		return "[]"
	}
	out := "["
	for i := 0; i < m.rows; i++ {
		if i > 0 {
			out += " "
		}
		for j := 0; j < m.cols; j++ {
			out += fmt.Sprintf("%f", m.elements[i*m.cols+j]) + ","
		}
		if i < m.rows-1 {
			out = out[:len(out)-1]
		}
		out += "\n"
	}
	out = out[:len(out)-2]
	out += "]"
	return out
}

func (m *Matrix) Get(row, col int) (float64, error) {
	if row >= m.rows || col >= m.cols {
		return 0, IndexError
	}
	return m.elements[row*m.cols+col], nil
}

func (m *Matrix) Set(row, col int, v float64) error {
	if row >= m.rows || col >= m.cols {
		return IndexError
	}
	m.elements[row*m.cols+col] = v
	return nil
}

func (m *Matrix) SetCol(col int, v float64) error {
	if col >= m.cols {
		return IndexError
	}
	for i := 0; i < m.rows; i++ {
		m.elements[i*m.cols+col] = v
	}
	return nil
}

func (m *Matrix) SetRow(row int, v float64) error {
	if row >= m.rows {
		return IndexError
	}
	for j := 0; j < m.cols; j++ {
		m.elements[row*m.cols+j] = v
	}
	return nil
}

func (m *Matrix) SetSlice(row, col int, n *Matrix) error {
	if n.rows+row > m.rows || n.cols+col > m.cols {
		return IndexError
	}
	for i := 0; i < n.rows; i++ {
		for j := 0; j < n.cols; j++ {
			m.elements[(i+row)*m.cols+(j+col)] = n.elements[i*n.cols+j]
		}
	}
	return nil
}

func (m *Matrix) Copy() *Matrix {
	elements := make([]float64, m.rows*m.cols)
	copy(elements, m.elements)
	return &Matrix{elements, m.rows, m.cols}
}

func (m *Matrix) Sum() float64 {
	var s float64
	for _, x := range m.elements {
		s += x
	}
	return s
}

func (m *Matrix) Add(n *Matrix) error {
	if m.rows != n.rows || m.cols != n.cols {
		return DimensionsError
	}
	for i, x := range n.elements {
		m.elements[i] += x
	}
	return nil
}

func (m *Matrix) Sub(n *Matrix) error {
	if m.cols != n.cols || m.rows != n.rows {
		return DimensionsError
	}
	for i, x := range n.elements {
		m.elements[i] -= x
	}
	return nil
}

func (m *Matrix) Slice(row, col, rows, cols int) (*Matrix, error) {
	if row+rows > m.rows || col+cols > m.cols {
		return new(Matrix), IndexError
	}
	elements := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = m.elements[(i+row)*m.cols+(j+col)]
		}
	}
	return &Matrix{elements, rows, cols}, nil
}

func (m *Matrix) MultiplyScalar(x float64) {
	for i := range m.elements {
		m.elements[i] *= x
	}
}

func (m *Matrix) DivideScalar(x float64) {
	for i := range m.elements {
		m.elements[i] /= x
	}
}

func (m *Matrix) Equals(n *Matrix) bool {
	if m.rows != n.rows || m.cols != n.cols {
		return false
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if m.elements[i*m.cols+j] != n.elements[i*n.cols+j] {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Pow(p float64) {
	for i, x := range m.elements {
		m.elements[i] = math.Pow(x, p)
	}
}
