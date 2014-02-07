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
	"fmt"
	"strings"
)

// matrix interface defining matrix operations
type Matrix interface {
	// Return true if the matrix is nil.
	Nil() bool

	// Return the number of rows of this matrix
	Rows() int

	// Return the number of columns of this matrix
	Cols() int

	// Return the number of elements contained in this matrix
	NumElements() int

	// Size of the leading index
	LeadingIndex() int

	// Return the dimension of the matrix
	Dimension() (int, int)

	// Get the value in the ith row and jth column
	Get(int, int) float64

	// Set the value in the ith row and jth column
	Set(int, int, float64, error)

	// Returns an array of slices referencing the matrix data. 
	// Changes to the slices effect changes to the matrix.
	Arrays() [][]float64

	// Returns the contents of this matrix stored into a flat array (row-major).
	Array() []float64

	// Return the i-th row elements
	RowSlice(row int) []float64

	// The pretty-print string
	String() string
}

type matrix struct {
	rows int
	cols int
}

func (M *matrix) Nil() bool { return M == nil }

func (M *matrix) Rows() int { return M.rows }

func (M *matrix) Cols() int { return M.cols }

func (M *matrix) NumElements() int { return M.rows * M.cols }

func (M *matrix) Dimension() (rows, cols int) {
	rows = M.rows
	cols = M.cols
	return
}


func String(A Matrix) string {
	condense := func(vs string) string {
		if strings.Index(vs, ".") != -1 {
			for vs[len(vs)-1] == '0' {
				vs = vs[0 : len(vs)-1]
			}
		}
		if vs[len(vs)-1] == '.' {
			vs = vs[0 : len(vs)-1]
		}
		return vs
	}

	if A == nil {
		return "{nil}"
	}
	s := "{"

	maxLen := 0
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			v := A.Get(i, j)
			vs := condense(fmt.Sprintf("%f", v))

			maxLen = maxInt(maxLen, len(vs))
		}
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			v := A.Get(i, j)

			vs := condense(fmt.Sprintf("%f", v))

			for len(vs) < maxLen {
				vs = " " + vs
			}
			s += vs
			if i != A.Rows()-1 || j != A.Cols()-1 {
				s += ","
			}
			if j != A.Cols()-1 {
				s += " "
			}
		}
		if i != A.Rows()-1 {
			s += "\n "
		}
	}
	s += "}"
	return s
}
