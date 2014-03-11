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
	"log"
	"github.com/numb3r3/gorec/utils"
)

type inmemDataset struct {

	instances []*Instance

	finalized bool

	options DatasetOptions

	featureDIct, labelDict *utils.Dictionary

	userFeatureDict, userLabelDict bool
}

func NewInmemDataset() *inmemDataset {
	dataset := new(inmemDataset)
	return dataset
}

func (dataset *inmemDataset) NumInstance() int {
	dataset.CheckFinalized(true)
	return len(dataset.instances)
}

func (dataset *inmemDataset) CreateIterator() DatasetIterator {
	return &inmemDatasetIterator{dataset: dataset}
}

func (dataset *inmemDataset) GetFeatureDictionary() *utils.Dictionary {
	return dataset.featureDIct
}

func (dataset *inmemDataset) GetLabelDictionary() *utils.Dictionary {
	return dataset.labelDict
}

func (dataset *inmemDataset) GetOptions() DatasetOptions {
	return dataset.options
}

func (dataset *inmemDataset) AddInstance(instance *Instance) bool {
	
}

func (dataset *inmemDataset) Finalize() {
	dataset.CheckFinalized(false)
	dataset.finalized = true
}

func (dataset *inmemDataset) CheckFinalized(stat bool) {
	if dataset.finalized != stat {
		if stat {
			log.Fatal("The dataset must be freezeen before iterating")
		} else {
			log.Fatal("The data in the dataset cannot be modified after freezen.")
		}
	}
}










