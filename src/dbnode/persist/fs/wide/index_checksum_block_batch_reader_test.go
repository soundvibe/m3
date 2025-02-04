// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE

package wide

import (
	"testing"

	"github.com/m3db/m3/src/x/ident"

	"github.com/stretchr/testify/assert"
)

func TestIndexChecksumBlockBatchReader(t *testing.T) {
	ch := make(chan ident.IndexChecksumBlockBatch)
	buf := NewIndexChecksumBlockBatchReader(ch)
	bl := ident.IndexChecksumBlockBatch{EndMarker: []byte("foo")}
	bl2 := ident.IndexChecksumBlockBatch{
		Checksums: []int64{1, 2, 3},
		EndMarker: []byte("bar"),
	}

	go func() {
		ch <- bl
		ch <- bl2
		close(ch)
	}()

	assert.True(t, buf.Next())
	assert.Equal(t, bl, buf.Current())
	assert.True(t, buf.Next())
	assert.Equal(t, bl2, buf.Current())
	assert.False(t, buf.Next())
}
