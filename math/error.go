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
)

const (
	// The matrix returned was nil
	errorNilMatrix = iota
	// The dimensions of the inputs do not make sense
	errorDimensionMismatch
	// The indices provided are out of bounds
	errorIllegalIndex
	// The matrix provided has a singularity
	exceptionSingular
	// The matrix provided is not positive semi-definite.
	exceptionNotSPD
)

type error_ int

func (e error_) Error() string {
	switch e {
	case errorNilMatrix:
		return "Matrix is nil"
	case errorDimensionMismatch:
		return "Input dimensions do not match"
	case errorIllegalIndex:
		return "Index out of bounds"
	case exceptionSingular:
		return "Matrix is singular"
	case exceptionNotSPD:
		return "Matrix is not positive semidefinite"
	}
	return fmt.Sprintf("Unknown error code %d", e)
}

func (e error_) String() string {
	return e.Error()
}

var (
	//The matrix returned was nil.
	ErrorNilMatrix error_ = error_(errorNilMatrix)
	//The dimensions of the inputs do not make sense for this operation.
	ErrorDimensionMismatch error_ = error_(errorDimensionMismatch)
	//The indices provided are out of bounds.
	ErrorIllegalIndex error_ = error_(errorIllegalIndex)
	//The matrix provided has a singularity.
	ExceptionSingular error_ = error_(exceptionSingular)
	//The matrix provided is not positive semi-definite.
	ExceptionNotSPD error_ = error_(exceptionNotSPD)
)


















