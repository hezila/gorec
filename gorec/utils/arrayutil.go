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

import "fmt"

type filterInt func(int) bool

func IntFilter(slice []int, f filterInt) []int {
     var result []int
     for _, value := range slice {
     	 if f(value) {
	    result = append(result, value)
	 }
     }
     return result
}

type filterFloat func(float) bool

func FloatFilter(slice []float, f filterFloat) []float {
     var result []float;
     for _, value := range slice {
     	 if f(value) {
	    result = append(result, value)
	 }
     }
     return result
}