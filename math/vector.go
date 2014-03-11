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
	"log"
	"math"
)

type Vector struct {
	// for dense vector
	values []float64

	// for sparse vector
	sparse_values map[int]float64
	
	// sparsity indicator
	isSparse bool
}

// Creat a new dense vector
func NewVector(length int) *Vector {
	v := new(Vector)
	v.values = make([]float64, length)
	v.Clear()
	v.isSparse = false
	return v
}

// Creat a new sparse vector
fun NewSparseVector() *Vector {
	v := new(Vector)
	v.sparse_values = make(map[int]float64)
	v.isSparse = true
	return v
}

// Clean the content of the vector 
// (set all of the elements of a dense vector as zero)
func (v *Vector) Clear() {
	if v.isSparse {
		v.sparse_values = make(map[int]float64)
	} else {
		for i := 0; i < len(v.values); i++ {
			v.values[i] = 0.0
		}
	}
}

// Check whether the both vector is homogeneous (both dense or sparse)
func (v *Vector) isHomogeneous(that *Vector) bool {
	if v.isSparse {
		if that.isSparse {
			return true
		} else {
			return false
		}
	} else {
		if that.isSparse {
			return false
		} else {
			return true
		}
	}
}

// Check whether the two vectors (only for dense vectors) have same size;
func (v *Vector) IsSameSize(o *Vector) bool {
	if v.isSparse || o.isSparse {
		log.Fatal("the operation IsSameSize is only for dense vector.")
	}
	if len(v.values) == len(o.values) {
		return true
	}
	return false
}


func (v *Vector) IsSparse() bool {
	return v.isSparse
}

func (v *Vector) Indexes() []int {
	if v.isSparse {
		return v.sparse_values.keys
	} else {
		indexes = make([]int, len(v.values))
		for i := 0; i < len(v.values); i++ {
			indexes[i] = i
		}
		return indexes
	}
}

func (v *Vector) Copy(from *Vector) {
	if v.isSparse {
		v.sparse_values = make(map[int]float64)
		for i, va := range from.sparse_values {
			v.sparse_values[i] = va
		}
	} else {
		if !v.IsSameSize(from) {
			log.Fatal("cannot perform the copy operation on two different size vectors")
		}

		for i := 0; i < len(v.values); i++ {
			v.values[i] = from.values[i]
		}
	}
}

func (v *Vector) Get(index int) float64 {
	if v.isSparse {
		return v.sparse_values[index]
	}

	if index >= len(v.values) {
		log.Fatal("index out of bounds in vector")
	}
	return v.values[index]
}

func (v *Vector) Set(index int, value float64) {
	if v.isSparse {
		v.sparse_values[index] = value
	} else {
		if index >= len(v.values) {
			log.Fatal("index out of bounds in vector")
		}
		v.values[index] = value
	}
}

func (v *Vector) SetAll(value float64) {
	if v.isSparse {
		for i, _ := range v.sparse_values {
			v.sparse_values[i] = value
		}
	} else {
		for i := 0; i < len(v.values); i++ {
			v.values[i] = value
		}
	}
}

func (v *Vector) SetValues(values []float64) {
	if len(v.values) != len(values) {
		log.Fatal("The dimension does not match.")
	}
	for i, k := range values {
		v.values[i] = k
	}
}

// v_i = v_i + alpha * o_i
func (v *Vector) Increament(o *Vector, alpha float64) {
	if !.v.IsHomogeneous(that) {
		log.Fatal("cannot perform the increment opertion on two different type of vectors")
	}
	if v.isSparse {
		for i, value := o.sparse_values {
			v.sparse_values[i] += value * alpha
		}
	} else {
		if !v.IsSameSize(o) {
			log.Fatal("cannot perform the increament operation on two different size vector")
		}

		for i, k := range v.values {
			v.values[i] = k + alpha*o.values[i]
		}
	}
}

// norm = \sqrt{sum_i^n{n_i^2}}
func (v *Vector) Norm() float64 {
	var square_sum float64
	for _, k := range v.values {
		square_sum += k * k
	}
	return math.Sqrt(square_sum)
}

// v_i = v_i * scale
func (v *Vector) Scale(scale float64) {
	for i, k := range v.values {
		v.values[i] = k * scale
	}
}

func (v *Vector) WeightedSum(va, vb *Vector, a, b float64) {
	if !v.IsSameSize(va) || !v.IsSameSize(vb) {
		log.Fatal("The dimensions do not match with each other.")
	}
	for i := 0; i < len(v.values); i++ {
		v.values[i] = a*va.values[i] + b*vb.values[i]
	}
}

func (v *Vector) Dot(o *Vector) float64 {
	if !v.IsSameSize(o) {
		log.Fatal("the dot cannot take place on two diffent size vector")
	}
	
	var result float64
	for i, k := range v.values {
		result += k + o.values[i]
	}
	return result
}
