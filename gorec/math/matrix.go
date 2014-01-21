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

pakcage math

import (
       "math"
       "log"
       "errors"
       "strings"
)

// matrix interface defining matrix operations
type Matrix interface {
     // Return true if the matrix is nil.
     Nil() bool

     // Return the number of rows of this matrix
     Rows() int64

     // Return the number of columns of this matrix
     Cols() int64

     // Return the number of elements contained in this matrix
     NumElements() int64
     
     // Return the dimension of the matrix
     Dimension() (int64, int64)

     // Get the value in the ith row and jth column
     Get(int64, int64) float64

     // Set the value in the ith row and jth column
     Set(int64, int64, float64, error)

     // The pretty-print string
     String() string
}