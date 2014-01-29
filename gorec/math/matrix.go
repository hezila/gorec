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
	"errors"
	"log"
	"math"
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

	// The pretty-print string
	String() string
}

// Stacking direction for matrix constructor
type Stacking int

const StackDown = Stacking(0)
const stackRight = Stacking(1)

// Matrix constructor data order
type DataOrder int

const RowOrder = DataOrder(0)
const ColumnOrder = DataOrder(1)

// Tridiagonal matrix type
type Tridiagonal int

const Symetric = Tridiagonal(0)
const Lower = Tridiagonal(1)
const Upper = Tridiagonal(2)

// Matrix dimensions, rows, cols and leading index.
// leading index is equal to row count.
type dimensions struct {
	rows int
	cols int
	// actual offset between leading index
	step int
}
