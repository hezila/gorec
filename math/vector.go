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
	values []float64
}

// Creat a new vector
func NewVector(length int) *Vector {
	v := new(Vector)
	v.values = make([]float64, length)
	v.Clear()
	return v
}



func (v *Vector) Clear() {
	for i := 0; i < len(v.values); i++ {
		v.values[i] = 0.0
	}
}

func (v *Vector) IsSameSize(o *Vector) bool {
	if len(v.values) == len(o.values) {
		return true
	}
	return false
}

func (v *Vector) Copy(from *Vector) {
	if !v.IsSameSize(from) {
		log.Fatal("cannot perform the copy operation on two different size vectors")
	}

	for i := 0; i < len(v.values); i++ {
		v.values[i] = from.values[i]
	}
}

func (v *Vector) Get(index int) float64 {
	if index >= len(v.values) {
		log.Fatal("index out of bounds in vector")
	}
	return v.values[index]
}

func (v *Vector) Set(index int, value float64) {
	if index >= len(v.values) {
		log.Fatal("index out of bounds in vector")
	}
	v.values[index] = value
}

func (v *Vector) SetAll(value float64) {
	for i := 0; i < len(v.values); i++ {
		v.values[i] = value
	}
}

func (v *Vector) SetValues(values []float64) {
	if len(v.values) != len(values) {
		log.Fatal("not same dimensions")
	}
	for i, k := range values {
		v.values[i] = k
	}
}

// v_i = v_i + alpha * o_i
func (v *Vector) Increament(o *Vector, alpha float64) {
	if !v.IsSameSize(o) {
		log.Fatal("cannot perform the increament operation on two different size vector")
	}

	for i, k := range v.values {
		v.values[i] = k + alpha*o.values[i]
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
		log.Fatal("not same dimensions")
	}
	for i := 0; i < len(v.values); i++ {
		v.values[i] = a*va.values[i] + b*vb.values[i]
	}
}
