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

package route

import (
	"hash/crc32"
	"strconv"
	"strings"

	"github.com/berkaroad/saashard/errors"
)

// ShardAlgorithm shard algorithm
type ShardAlgorithm func(val string, dataNodeCount int, params ...interface{}) (int, error)

// ParseShardAlgorithm parse ShardAlgorithm
func ParseShardAlgorithm(name string) ShardAlgorithm {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	var algo ShardAlgorithm
	switch name {
	case "mod":
		algo = ModShardAlgo
	default:
		algo = HashShardAlgo
	}
	return algo
}

// HashShardAlgo hash shard algorithm
func HashShardAlgo(val string, dataNodeCount int, params ...interface{}) (int, error) {
	val = strings.Trim(val, "'")
	hashCode := crc32.ChecksumIEEE([]byte(val))
	index := int(hashCode) % dataNodeCount
	return index, nil
}

// ModShardAlgo mod shard algorithm
func ModShardAlgo(val string, dataNodeCount int, params ...interface{}) (int, error) {
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.ErrMustPositiveIntegerInModShard
	}
	if num <= 0 {
		return 0, errors.ErrMustPositiveIntegerInModShard
	}
	index := (num - 1) % dataNodeCount
	return index, nil
}
