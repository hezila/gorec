// Copyright (c) 2014 Feng Wang <wffrank1987@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language

package math

import (
	"math/rand"
)

type SparseMatrix struct {
	matrix

	// flatted elements
	elements map[int]float64

	// offset to start of matrix s.t. idx = i * cols + j + offset
	// offset  = staring row * step + staring col
	offset int
	// analogous to dense step
	step int
}


func MakeSparseMatrix(elements map[int]float64, rows, cols int) *SparseMatrix {
	M := new(SparseMatrix)
	M.rows = rows
	M.cols = cols
	M.offset = 0
	M.step = cols

	M.elements = elements
	return M
}


func (M *SparseMatrix) GetRowColIndex(index int) (i, j int) {
	i = (index - M.offset) / M.step
	j = (index - M.offset) % M.step
	return
}

func (M *SparseMatrix) GetRowIndex(index int) (i int) {
	i = (index - M.offset) / M.step
	return
}

func (M *SparseMatrix) GetRowIndex(index int) (j int) {
	j = (index - M.offset) % M.step
	return
}

func (M *SparseMatrix) Get(i, j int) float64 {
	i = i % M.rows
	if i < 0 {
		i = M.rows + i
	}

	j = j % M.cols
	if j < 0 {
		j = M.cols + j
	}

	v, ok := M.elements[i * M.step + j + M.offset]
	if !ok {
		return 0
	}
	return v
}

// Looks up an element given its element index
func (M *SparseMatrix) GetValue(index int) float64 {
	x, ok := M.elements[index]
	if !ok {
		return 0
	}
	return x
}

func (M *SparseMatrix) Set(i, j int, v float64) {
	i = i % M.rows
	if i < 0 {
		i = M.rows + i
	}

	j = j % M.cols
	if j < 0 {
		j = M.cols + j
	}

	if v == 0 {
		delete(M.elements, i * M.step + j + M.offset)
	} else {
		M.elements[i * M.step + j + M.offset] = v
	}

}

func (M *SparseMatrix) SetValue(index int, v float64) {
	if v == 0 {
		delete(M.elements, index)
	} else {
		M.elements[index] = v
	}
}

func (M *SparseMatrix) Indices() (out chan int) {
	// maybe thread the populating?
	out = make(chan int)
	go func(o chan int) {
		for index := range M.elements {
			i, j := M.GetRowColIndex(index)
			if 0 <= i && i < M.rows && 0 <= j && j < M.cols {
				o <- index
			}
		}
		close(o)
	}(out)
	return
}


func (M *SparseMatrix) SubMatrix(i, j, row, cols int) *SparseMatrix {
	if i < 0 || j < 0 || i + rows > M.rows || j + cols > M.cols {
		i = maxInt(0, i)
		j = maxInt(0, j)
		rows = minInt(M.rows - i, rows)
		cols = minInt(M.cols - j, cols)
	}
	S := ZerosSparse(rows, cols)

	for index, value := range M.elements {
		r, c := M.GetRowColIndex(index)
		if r < i + row && c < j + cols {
			S.Set(r-i, c-j, value)
		}
	}
	
	return S
}

func (M *SparseMatrix) ColVector(j int) *SparseMatrix {
	return M.SubMatrix(0, j, M.rows, 1)
}

func (M *SparseMatrix) RowVector(i int) *SparseMatrix {
	return M.SubMatrix(i, 0, 1, M.cols)
}

// Create a new matrix [A B]
func (A *SparseMatrix) Augment(B *SparseMatrix) (*SparseMatrix, error) {
	if A.rows != B.rows {
		return nil, ErrorDimensionMismatch
	}

	C := ZerosSparse(A.rows, A.cols + B.cols)

	for index, value := range A.elements {
		i, j = A.GetRowColIndex(index)
		C.Set(i, j, value)
	}

	for index, value := range B.elements {
		i, j = B.GetRowColIndex(index)
		C.Set(i, j+A.cols, value)
	}
	
	return C, nil
}

func (A *SparseMatrix) Stack(B *SparseMatrix) (*SparseMatrix, error) {
	if A.cols != B.cols {
		return nil, ErrorDimensionMismatch
	}
	
	C := ZerosSparse(A.rows + B.rows, A.cols)

	for index, value := range A.elements {
		i, j := A.GetRowColIndex(index)
		C.Set(i, j, value)
	}

	for index, value := range B.elements {
		i, j := B.GetRowColIndex(index)
		C.Set(i, j, value)
	}

	return C, nil
}

func (M *SparseMatrix) L() *SparseMatrix {
	B := ZerosSparse(M.rows, M.cols)
	for index, value := range M.elements {
		i, j := M.GetRowColIndex(index)
		if i >= j {
			B.Set(i, j, value)
		}
	}
	return B
}

func (M *SparseMatrix) U() *SparseMatrix {
	U := ZerosSparse(M.rows, M.cols)
	for index, value := range M.elements {
		i, j := M.GetRowColIndex(index)
		if i <= j {
			U.Set(i, j, value)
		}
	}
	return U
}

func (M *SparseMatrix) Copy() *SparseMatrix {
	C := ZerosSparse(M.rows, M.cols)
	for index, value := range M.elements {
		C.elements[index] = value
	}
	return C
}

// Convert this sparse matrix into a dense matrix
func (M *SparseMatrix) DenseMatrix() *DenseMatrix {
	D := Zeros(M.rows, M.cols)
	for index, value := range M.elements {
		i, j = M.GetRowColInddex(index)
		D.Set(i, j, value)
	}
	return D
}

func (M *SparseMatrix) String() string {return String(M)}

func ZerosSparse(rows, cols int) *SparseMatrix {
	M := new(SparseMatrix)
	M.rows = rows
	M.cols = cols
	M.offset = 0
	M.step = cols
	M.elements = map[int]float64{}
	return M
}

func OnesSparse(rows, cols int) *SparseMatrix {
	O := new(SparseMatrix)
	O.rows = rows
	O.cols = cols
	O.step = cols
	O.elements = map[int]float64{}
	for i := 0; i < cols*cols; i++ {
		O.elements[i] = 1
	}
	return O
}

func EyeSparse(size int) *SparseMatrix {
	E := ZerosSparse(size, size)

	for i := 0; i < size; i++ {
		E.Set(i, i, 1)
	}
	return E
}

func NormalsSparse(rows, cols int) *SparseMatrix {
	N := ZerosSparse(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			N.Set(i, j, rand.NormFloat64())
		}
	}
	return N
}

func Diagonal(d []float64) *SparseMatrix {
	n := len(d)
	D := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		D.Set(i, i, d[i])
	}
	return D
}

func MakeSparseCopy(M Matrix) *SparseMatrix {
	A := ZerosSparse(M.Rows(), M.Cols())
	for i := 0; i < M.Rows(); i++ {
		for j := 0; j < M.Cols(); j++ {
			A.Set(i, j, M.Get(i, j))
		}
	}
	return A
}
