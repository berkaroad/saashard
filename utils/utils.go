// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"reflect"
)

// Contains is a convenience function that returns
// true if str matches any of the values.
func Contains(arr interface{}, item interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrValue.Len(); i++ {
			if arrValue.Index(i).Interface() == item {
				return true
			}
		}
	case reflect.Map:
		if arrValue.MapIndex(reflect.ValueOf(item)).IsValid() {
			return true
		}
	}
	return false
}

// CollectionContains that large collection contains small collection.
func CollectionContains(largeArr interface{}, smallArr interface{}) bool {
	largeArrKind := reflect.TypeOf(largeArr).Kind()
	smallArrKind := reflect.TypeOf(smallArr).Kind()
	if (largeArrKind == reflect.Slice || largeArrKind == reflect.Array) &&
		(smallArrKind == reflect.Slice || smallArrKind == reflect.Array) {
		smallArrValue := reflect.ValueOf(smallArr)
		for i := 0; i < smallArrValue.Len(); i++ {
			smallArrItemValue := smallArrValue.Index(i).Interface()
			return Contains(largeArr, smallArrItemValue)
		}
	}
	return false
}

// StringCollectionUnion union with string collection.
func StringCollectionUnion(leftArr []string, rightArr []string) []string {
	coll := []string{}
	if leftArr != nil && len(leftArr) > 0 {
		for _, leftArrItem := range leftArr {
			if !Contains(coll, leftArrItem) {
				coll = append(coll, leftArrItem)
			}
		}
	}
	if rightArr != nil && len(rightArr) > 0 {
		for _, rightArrItem := range rightArr {
			if !Contains(coll, rightArrItem) {
				coll = append(coll, rightArrItem)
			}
		}
	}
	return coll
}

// StringCollectionIntersection intersection with string collection.
func StringCollectionIntersection(leftArr []string, rightArr []string) []string {
	coll := []string{}
	if leftArr != nil && len(leftArr) > 0 &&
		rightArr != nil && len(rightArr) > 0 {
		for _, leftArrItem := range leftArr {
			if !Contains(coll, leftArrItem) && Contains(rightArr, leftArrItem) {
				coll = append(coll, leftArrItem)
			}
		}
	}
	return coll
}
