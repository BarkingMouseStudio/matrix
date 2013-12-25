package matrix

import (
	"fmt"
	"testing"
)

func TestMatrixNew(t *testing.T) {
	m, _ := New([]float64{
		1, 2, 3, 4, 5, 6, 7, 8,
	}, 4, 2)
	if m.Rows() != 4 {
		t.Fatal("Matrix returned incorrect cols", m.Rows())
	}
	if m.Cols() != 2 {
		t.Fatal("Matrix returned incorrect cols", m.Cols())
	}
	if m.Size() != 8 {
		t.Fatal("Matrix returned incorrect length")
	}
	_, err := New([]float64{1, 2, 3, 4, 5, 6, 7, 8}, 4, 4)
	if err == nil {
		t.Fatal("Expected matrix error")
	}
}

func TestMatrixNewMulti(t *testing.T) {
	m := NewMulti([][]float64{
		[]float64{1, 2},
		[]float64{3, 4},
		[]float64{5, 6},
		[]float64{7, 8},
	})
	if m.Rows() != 4 {
		t.Fatal("Matrix returned incorrect cols", m.Rows())
	}
	if m.Cols() != 2 {
		t.Fatal("Matrix returned incorrect cols", m.Cols())
	}
	if m.Size() != 8 {
		t.Fatal("Matrix returned incorrect length")
	}
	n, _ := New([]float64{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2)
	if !m.Equals(n) {
		t.Fatal("Matrices did not match")
	}
}

func TestMatrixNewRandNorm(t *testing.T) {
	m := NewRandNorm(0.1, 0, 3, 3)
	if m.Size() != 9 {
		t.Fatal("Matrix returned incorrect length")
	}
}

func TestMatrixNewZeros(t *testing.T) {
	m := NewZeros(3, 2)
	if m.Size() != 6 || m.cols != 2 || m.rows != 3 {
		t.Fatal("Matrix returned incorrect length")
	}
}

func TestMatrixArray(t *testing.T) {
	a, err := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := New(a.Array(), 4, 4)
	if !a.Equals(b) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixReshape(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	b, err := Reshape(a, 8, 2)
	if err != nil {
		t.Fatal(err)
	}
	c, _ := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 8, 2)
	if !b.Equals(c) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixSlice(t *testing.T) {
	a, err := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	if err != nil {
		t.Fatal(err)
	}
	b, err := a.Slice(1, 1, 3, 3)
	if err != nil {
		t.Fatal(err)
	}
	expected, _ := New([]float64{6, 7, 8, 10, 11, 12, 14, 15, 16}, 3, 3)
	if !b.Equals(expected) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixSlice_error(t *testing.T) {
	a, err := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3)
	if err != nil {
		t.Fatal(err)
	}
	_, err = a.Slice(3, 1, 3, 2)
	if err == nil {
		t.Fatal("Expected dimensions error")
	}
}

func TestMatrixArrays(t *testing.T) {
	a, err := New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	if err != nil {
		t.Fatal(err)
	}
	b := a.Arrays()
	expected := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	if fmt.Sprintf("%v", b) != fmt.Sprintf("%v", expected) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixEquals(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	if !a.Equals(b) {
		t.Fatal("Expected equal")
	}
	a, _ = New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	b, _ = New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	if a.Equals(b) {
		t.Fatal("Expected not equal")
	}
	a, _ = New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ = New([]float64{0, 2, 3, 4, 5, 6}, 3, 2)
	if a.Equals(b) {
		t.Fatal("Expected not equal")
	}
}

func TestMatrixDot_expected1(t *testing.T) {
	a, _ := New([]float64{1, 0, 0, 1}, 2, 2)
	b, _ := New([]float64{4, 1, 2, 2}, 2, 2)
	c, err := Dot(a, b)
	expected, _ := New([]float64{4, 1, 2, 2}, 2, 2)
	if err != nil || !c.Equals(expected) {
		t.Fatal("Got incorrect results for dot", c, expected)
	}
}

func TestMatrixDot_expected2(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	c, err := Dot(a, b)
	expected, _ := New([]float64{9, 12, 15, 19, 26, 33, 29, 40, 51}, 3, 3)
	if err != nil || !c.Equals(expected) {
		t.Fatal("Got incorrect results for dot", c, expected)
	}
}

func TestMatrixDot_error(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	_, err := Dot(a, b)
	if err == nil {
		t.Fatal("Expected dimensions error")
	}
}

func TestMatrixTranspose(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 2, 3)
	if !Transpose(a).Equals(b) {
		t.Fatal("Got incorrectly transposed matrix", a, b)
	}
}

func TestMatrixGreater(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	c := Greater(a, b)
	expected, _ := New([]float64{0, 0, 0, 1, 1, 0}, 3, 2)

	if !c.Equals(expected) {
		t.Fatal("Returned unexpected results", c)
	}
}

func TestMatrixGreater_mismatch(t *testing.T) {
	a, _ := New([]float64{1, 9, 2, 5, 4, 5}, 3, 2)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 2, 3)

	c := Greater(a, b)
	expected, _ := New([]float64{0, 1, 0, 1}, 2, 2)

	if !c.Equals(expected) {
		t.Fatal("Returned unexpected results", c)
	}

	a, _ = New([]float64{1, 9, 2, 5, 4, 5}, 2, 3)
	b, _ = New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	c = Greater(a, b)
	expected, _ = New([]float64{0, 1, 0, 1}, 2, 2)

	if !c.Equals(expected) {
		t.Fatal("Returned unexpected results", c)
	}
}

func TestMatrixSetRow(t *testing.T) {
	m, _ := New([]float64{1, 9, 2, 5, 4, 5}, 2, 3)
	m.SetRow(0, 1)

	expected, _ := New([]float64{1, 1, 1, 5, 4, 5}, 2, 3)
	if !m.Equals(expected) {
		t.Fatal("Returned unexpected results", m)
	}
}

func TestMatrixSetRow_error(t *testing.T) {
	m, _ := New([]float64{1, 9, 2, 5, 4, 5}, 2, 3)
	err := m.SetRow(2, 1)

	if err == nil {
		t.Fatal("Expected index error")
	}
}

func TestMatrixSetCol(t *testing.T) {
	m, _ := New([]float64{1, 9, 2, 5, 4, 5}, 3, 2)
	m.SetCol(0, 1)

	expected, _ := New([]float64{1, 9, 1, 5, 1, 5}, 3, 2)
	if !m.Equals(expected) {
		t.Fatal("Returned unexpected results", m)
	}
}

func TestMatrixSetCol_error(t *testing.T) {
	m, _ := New([]float64{1, 9, 2, 5, 4, 5}, 3, 2)
	err := m.SetCol(2, 1)

	if err == nil {
		t.Fatal("Expected index error")
	}
}

func TestMatrixSub(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	c, _ := Sub(a, b)
	expected, _ := New([]float64{0, -1, -2, 2, 1, 0}, 3, 2)

	if !c.Equals(expected) {
		t.Fatal("Returned unexpected results", c)
	}
}

func TestMatrixSub_error(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	_, err := Sub(a, b)

	if err == nil {
		t.Fatal("Expected dimensions error")
	}
}

func TestMatrixAdd(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	c, _ := Add(a, b)
	expected, _ := New([]float64{2, 5, 8, 6, 9, 12}, 3, 2)

	if !c.Equals(expected) {
		t.Fatal("Returned unexpected results", c)
	}
}

func TestMatrixAdd_error(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	b, _ := New([]float64{1, 3, 5, 2, 4, 6}, 3, 2)

	_, err := Add(a, b)

	if err == nil {
		t.Fatal("Expected dimensions error")
	}
}

func TestMatrixSum(t *testing.T) {
	m, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	s := m.Sum()
	if s != 21 {
		t.Fatal("Returned unexpected results", s)
	}
}

func TestMatrixSet(t *testing.T) {
	m, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	err := m.Set(0, 2, -1)
	expected, _ := New([]float64{1, 2, -1, 4, 5, 6}, 2, 3)
	if err != nil {
		t.Fatal("Got an error setting row/col", err)
	}
	if !m.Equals(expected) {
		t.Fatal("Returned unexpected results", m)
	}
	err = m.Set(1, 3, 0)
	if err == nil {
		t.Fatal("Expected index err")
	}
}

func TestMatrixGet(t *testing.T) {
	m, _ := New([]float64{1, 2, 3, 4, 5, 6}, 2, 3)
	v, err := m.Get(0, 2)
	if err != nil {
		t.Fatal("Got an error setting row/col", err)
	}
	if v != 3 {
		t.Fatal("Returned unexpected results", v)
	}
	_, err = m.Get(1, 3)
	if err == nil {
		t.Fatal("Expected index err")
	}
}

func TestMatrixNewRand(t *testing.T) {
	m := NewRand(3, 3)
	if m.Size() != 9 {
		t.Fatal("Matrix returned incorrect length")
	}
}

func TestMatrixMap(t *testing.T) {
	m := NewOnes(3, 3)
	n := Map(m, func(x float64) float64 {
		return x * 2
	})

	for _, x := range n.elements {
		if x != 2 {
			t.Fatal("Got invalid results for matrix", n)
		}
	}
}

func TestMatrixMultiplyScalar(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	expected, _ := New([]float64{2, 4, 6, 8, 10, 12}, 3, 2)

	b := MultiplyScalar(a, 2)

	if !b.Equals(expected) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixDivideScalar(t *testing.T) {
	a, _ := New([]float64{2, 4, 6, 8, 10, 12}, 3, 2)
	expected, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)

	b := DivideScalar(a, 2)

	if !b.Equals(expected) {
		t.Fatal("Returned unexpected results", b)
	}
}

func TestMatrixPow(t *testing.T) {
	a, _ := New([]float64{2, 4, 6, 8, 10, 12}, 3, 2)
	a.Pow(2)

	expected, _ := New([]float64{4, 16, 36, 64, 100, 144}, 3, 2)

	if !a.Equals(expected) {
		t.Fatal("Returned unexpected results", a)
	}
}

func TestMatrixAugment(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{7, 7, 7}, 3, 1)
	c, err := Augment(a, b)
	if err != nil {
		t.Fatal("Got unexpected error", err)
	}
	expected, _ := New([]float64{1, 2, 7, 3, 4, 7, 5, 6, 7}, 3, 3)
	if !c.Equals(expected) {
		t.Fatal("Got unexpected result", c)
	}
	b, _ = New([]float64{7, 7}, 2, 1)
	_, err = Augment(a, b)
	if err == nil {
		t.Fatal("Expected dimensions error", err)
	}
}

func TestMatrixStack(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{7, 7}, 1, 2)
	c, err := Stack(a, b)
	if err != nil {
		t.Fatal("Got unexpected error", err)
	}
	expected, _ := New([]float64{1, 2, 3, 4, 5, 6, 7, 7}, 4, 2)
	if !c.Equals(expected) {
		t.Fatal("Got unexpected result", c)
	}
	b, _ = New([]float64{7, 7}, 2, 1)
	_, err = Stack(a, b)
	if err == nil {
		t.Fatal("Expected dimensions error", err)
	}
}

func TestMatrixSetSlice(t *testing.T) {
	a, _ := New([]float64{1, 2, 3, 4, 5, 6}, 3, 2)
	b, _ := New([]float64{7, 7, 7}, 3, 1)
	err := a.SetSlice(0, 0, b)
	if err != nil {
		t.Fatal("Got unexpected error", err)
	}
	expected, _ := New([]float64{7, 2, 7, 4, 7, 6}, 3, 2)
	if !a.Equals(expected) {
		t.Fatal("Got unexpected result", a)
	}
	err = a.SetSlice(0, 1, b)
	if err != nil {
		t.Fatal("Got unexpected error", err)
	}
	expected, _ = New([]float64{7, 7, 7, 7, 7, 7}, 3, 2)
	if !a.Equals(expected) {
		t.Fatal("Got unexpected result", a)
	}
	err = a.SetSlice(0, 2, b)
	if err == nil {
		t.Fatal("Expected index error", err)
	}
}

func ExampleMatrixStringer() {
	m, _ := New([]float64{
		1, 2, 3, 4, 5, 6, 7, 8,
	}, 4, 2)
	fmt.Println(m.String())
	// Output:
	// [1.000000,2.000000
	//  3.000000,4.000000
	//  5.000000,6.000000
	//  7.000000,8.000000]
}

func BenchmarkMatrixDot(b *testing.B) {
	b.StopTimer()
	dim := 100
	a := NewRand(dim, dim)
	c := NewRand(dim, dim)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Dot(a, c)
	}
}
