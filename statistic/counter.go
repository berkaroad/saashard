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

package statistic

import (
	"sync/atomic"
)

// Counter is a performance counter.
type Counter struct {
	OldClientQPS    int64
	OldErrLogTotal  int64
	OldSlowLogTotal int64

	ClientConns  int64
	ClientQPS    int64
	ErrLogTotal  int64
	SlowLogTotal int64
}

// IncrClientConns is to increase client conns.
func (c *Counter) IncrClientConns() {
	atomic.AddInt64(&c.ClientConns, 1)
}

// DecrClientConns is to decrease client conns.
func (c *Counter) DecrClientConns() {
	atomic.AddInt64(&c.ClientConns, -1)
}

// IncrClientQPS is to increase client qps.
func (c *Counter) IncrClientQPS() {
	atomic.AddInt64(&c.ClientQPS, 1)
}

// IncrErrLogTotal is to increase err log total.
func (c *Counter) IncrErrLogTotal() {
	atomic.AddInt64(&c.ErrLogTotal, 1)
}

// IncrSlowLogTotal is to increase slow log total.
func (c *Counter) IncrSlowLogTotal() {
	atomic.AddInt64(&c.SlowLogTotal, 1)
}

// FlushCounter is to flush the count per second
func (c *Counter) FlushCounter() {
	atomic.StoreInt64(&c.OldClientQPS, c.ClientQPS)
	atomic.StoreInt64(&c.OldErrLogTotal, c.ErrLogTotal)
	atomic.StoreInt64(&c.OldSlowLogTotal, c.SlowLogTotal)
	atomic.StoreInt64(&c.ClientQPS, 0)
}
