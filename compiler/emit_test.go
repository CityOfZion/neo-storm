package compiler

import (
	"bytes"
	"reflect"
	"testing"
)

func TestArrayReverse(t *testing.T) {
	var cases = []struct {
		actual []byte
		expect []byte
	}{
		{
			actual: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expect: []byte{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			actual: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8},
			expect: []byte{8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			actual: []byte{0, 1, 2, 3, 4, 5, 6, 7},
			expect: []byte{7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			actual: []byte{0},
			expect: []byte{0},
		},
	}

	for _, item := range cases {
		res := arrayReverse(item.actual)
		if !bytes.Equal(res, item.expect) || !reflect.DeepEqual(res, item.expect) {
			t.Fatalf("arrayReverse works wrong:\n \t actual: %#v \n \t expect: %#v", res, item.expect)
		}
	}
}
