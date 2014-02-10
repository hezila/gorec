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
	"github.com/numb3r3/gorec/utils"
	"testing"
	"fmt"
)

func TestBase(t *testing.T) {
	A := NewDenseMatrix(3, 3)
	fmt.Println(A.String())

	B := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	fmt.Println(String(B))

	C := MakeDenseMatrixStacked([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}, []float64{7, 8, 9}})
	fmt.Println(C.String())

}

func TestArrays(t *testing.T) {
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	fmt.Println(String(M))

	arrays := M.Arrays()
	arrays[1][1] = 111
	utils.Expect(t, "111", M.Get(1, 1))

	array := M.Array()
	array[2] = 100
	utils.Expect(t, "100", M.Get(0, 2))

	fmt.Println(M.Arrays())
	fmt.Println(M.Array())

	row_slice := M.RowSlice(1)
	row_slice[1] = 99
	utils.Expect(t, "99", M.Get(1, 1))
	fmt.Println(M.String())
	
}

func TestGetSet(t *testing.T) {
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)

	M.Set(1, 2, 33)
	fmt.Println(M.String())
	utils.Expect(t, "33", M.Get(1, 2))
}

func TestSubMatrix(t *testing.T) {
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	sub_mat := M.SubMatrix(1, 1, 2, 2)
	sub_mat.Set(1, 1, 0.0)
	sub_mat.Set(0, 0, 100.0)
	fmt.Println(sub_mat)
	fmt.Println(M.String())
}

func TestColVector(t *testing.T) {
	fmt.Println("TestColVector")
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	fmt.Println(M)

	col_vec := M.ColVector(0)
	fmt.Println(col_vec)

	col_vec.Set(1, 0, 100)
	utils.Expect(t, "5", M.Get(1, 0))
}

func TestRowVector(t *testing.T) {
	fmt.Println("TestColVector")
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	fmt.Println(M)

	row_vec := M.RowVector(0)
	fmt.Println(row_vec)

	row_vec.Set(0, 1, 200)
	utils.Expect(t, "2", M.Get(0, 1))
}

func TestMatrixCopy(t *testing.T) {
	M := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4, 4)
	C := M.Copy()

	utils.Expect(t, "7", C.Get(1, 2))

	C.Set(1, 2, 77)
	utils.Expect(t, "7", M.Get(1, 2))
}














