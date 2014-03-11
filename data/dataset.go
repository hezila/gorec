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

import (
	"github.com/numb3r3/gorec/utils"
)

// The interface to access the dataset
type Dataset interface {

	// Get the number of samples/insances in the dataset
	NumInstance() int
	
	// Create a iterator to access each instance one by one
	CreateIterator() DatasetIterator

	// Get the options of the dataset
	GetOptions() DatasetOptions
	
	// Get the feature identification name dictionary
	GetFeatureDictionary() *utils.Dictionary

	// Ge the target label dictionary
	GetLabelDictionary() *utils.Dictionary
}

// The structure of dataset options
type DatasetOptions struct {

	// Wether the feature is stored in sparse vector
	FeatureIsSparse bool
	
	// The dimenion of the feature vector
	FeatureDimension int
	

	// Wether it is a supervised learning problem
	IsSupervisedLearning bool

	// The number of target labels
	NumLabels int
	
	// Other options
	Options interface()
}

func ConvertNamedFeatures(instance *Instance, dict *utils.Dictionary) {
	if instance.Features != nil {
		return
	}

	instance.Features = NewSparseVector()
	
	// The first element value is asways 1.0
	instance.Features.Set(0, 1.0)

	for k, v := range instance.NamedFeatures {
		id : = dict.GetIdFromName(k)
		instance.Features.Set(id, v)
	}
}
