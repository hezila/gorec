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

package utils

import (
	"reflect"
)

// Return sequence integers with the begining value as x
func IntSequence(begin int) func() int {
	i := begin
	return func() int {
		i += 1
		return i
	}
}

// from: https://gist.github.com/rafkhan/6501567

// map function types
type mapf func(interface{}) interface{}

// reduce function types
type reducef func(interface{}, interface{}) interface{}

// filter function types
type filterf func(interface{}) bool

// Map(slice, func)
//
// Usages:
//     a := []int{1, 2, 3, 4}
//     b := Map(a, func(val interface{}) interface{} {
//     	 return val.(int) * 2
//     })
// should be [2, 4, 6, 8]
func Map(in interface{}, fn mapf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, val.Len())

	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface())
	}
	return out
}

//Reduce(slice, starting value, func)
//
//Usage:
// c := Reduce(b, 0, func(val interface{}, memo interface{}) interface{} {
//   return memo.(int) + val.(int)
// })
func Reduce(in interface{}, memo interface{}, fn reducef) interface{} {
	val := reflect.ValueOf(in)

	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}
	return memo
}

//Filter(slice, predicate func)
//Usage:
// d := Filter(b, func(val interface{}) bool {
//   return val.(int) % 4 == 0
//})
func Filter(in interface{}, fn filterf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()
		if fn(current) {
			out = append(out, current)
		}
	}
	return out
}
