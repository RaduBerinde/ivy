// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package value

import (
	"fmt"
	"strconv"
)

// Int is not only the simplest representation, it provides the operands that mix
// types upward. That is, BigInt.Add(Int) will be done by rewriting as Int.Add(BigInt).

type Int struct {
	x int64
}

const (
	intBits = 32
	minInt  = -(1 << (intBits - 1))
	maxInt  = 1<<(intBits-1) - 1
)

func SetIntString(s string) (Int, error) {
	i, err := strconv.ParseInt(s, 0, intBits)
	return Int{x: i}, err
}

func (i Int) String() string {
	return fmt.Sprint(i.x)
}

func (i Int) Eval() Value {
	return i
}

func (i Int) ToType(which valueType) Value {
	switch which {
	case intType:
		return i
	case bigIntType:
		return bigInt64(i.x)
	case bigRatType:
		return bigRatInt64(i.x)
	case vectorType:
		return ValueSlice([]Value{i})
	}
	panic("Int.ToType")
}