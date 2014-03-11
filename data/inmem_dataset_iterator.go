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
)

type inmemDatasetIterator struct {
	dateset *inmemDataset
	currIndex int
}

fucn (it *inmemDatasetIterator) Start() {
	it.dataset.CheckFinalized(true)
	it.currIndex = 0
}

func (it *inmemDatasetIterator) End() bool {
	it.dateset.CheckFinalized(true)
	if it.currIndex >= len(it.dataset.instances) {
		return true
	}
	return false
}

func (it *inmemDatasetIterator) Next() {
	it.dateset.CheckFinalized(true)
	if !it.End() {
		it.currIndex++
	}
}

func (it *inmemDatasetIterator) Skip(n int) {
	it.dateset.CheckFinalized(true)
	if n < 0 {
		log.Fatal("Skip step must be non-negative.")
	}
	it.currIndex += n
}

func (it *inmemDatasetIterator) GetInstance() *Instance {
	it.dateset.CheckFinalized(true)
	if it.End() {
		return nil
	}
	return it.dateset.instances[it.currIndex]
}
