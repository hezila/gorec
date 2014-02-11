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

package data

// one sample item

type Instance struct {

	// Sample features
	Features []float64

	// Features indexing by "Name"
	NamedFeatures map[string]float64

	// Label/Output
	// Only used for fixing the supervised model
	// it is assinged with nil in the unsupervised fashion
	Output *InstanceOutput

	// the name of the sample
	// Can be nil
	Name string

	// Addtional information
	Attachement interface{}
}
