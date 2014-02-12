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

func ZerosSparse(rows, cols int) *SparseMatrix {
	M := new(SparseMatrix)
	M.rows = rows
	M.cols = cols
	M.offset = 0
	M.step = cols
	M.elements = map[int]float64{}
	return M
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







