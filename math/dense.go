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

// A matrix backed by a flat array of all elements
type DenseMatrix struct {
	matrix
	
	// flatted elements
	elements []float64
	
	// offset between rows; step = cols
	step int
}

func NewDenseMatrix(rows, cols int) *DenseMatrix {
	M := new(DenseMatrix)
	M.rows = rows
	M.cols = cols
	M.step = cols
	
	M.elements = make([]float64, rows * cols)
	for i := 0; i < rows * cols; i++ {
		M.elements[i] = 0.0
	}
	return M
}

func MakeDenseMatrix(elements []float64, rows, cols int) *DenseMatrix {
	A := new(DenseMatrix)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.elements = elements
	return A
}

func MakeDenseMatrixStacked(data [][]float64) *DenseMatrix {
	rows := len(data)
	cols := len(data[0])
	elements := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = data[i][j]
		}
	}
	return MakeDenseMatrix(elements, rows, cols)
}

func (M *DenseMatrix) Arrays() [][]float64 {
	a := make([][]float64, M.rows)
	for i := 0; i < M.rows; i++ {
		a[i] = A.elements[i*M.step : i*M.step + M.cols]
	}
	return a
}

func (M *DenseMatrix) Array() []float64 {
	if M.step == M.rows {
		return A.elements[0 : M.rows*M.cols]
	}
	a := make([]float64, M.rows*M.cols)
	for i := 0; i < M.rows; i++ {
		for j := 0; j < M.cols; j++ {
			a[i*M.cols + j] = M.elements[i*M.step + j]
		}
	}
	return a
}

func (M *DenseMatrix) rowSlice(row int) []float64 {
	return M.elements[row*M.step : row*M.step + M.cols]
}




















