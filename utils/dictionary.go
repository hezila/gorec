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

type Dictionary struct {
	nameToId map[string]int
	idToName map[int]string
	maxId int
	minId int
}

func NewDictionary(minId int) *Dictionary {
	dict := new(Dictionary)
	dict.nameToId = make(map[string]int)
	dict.idToName = make(map[int]string)
	dict.maxId = minId
	dict.minId = minId
	return dict
}

func (d *Dictionary) GetIdFromName(name string) int {
	id, ok := d.nameToId[name]
	if ok {
		return id
	}

	return -1
}

func (d *Dictionary) GetNameFromId(id int) string {
	return d.idToName[id]
}

func (d *Dictionary) AddName(name string) int {
	id, ok := d.nameToId[name]
	if ok {
		return id
	}
	d.nameToId[name] = d.maxId
	d.idToName[d.maxId] = name
	d.maxId ++
	return d.maxId - 1
}





















