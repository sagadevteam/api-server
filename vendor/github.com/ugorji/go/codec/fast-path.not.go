<<<<<<< HEAD
// Copyright (c) 2012-2018 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// +build notfastpath

package codec

import "reflect"

const fastpathEnabled = false

// The generated fast-path code is very large, and adds a few seconds to the build time.
// This causes test execution, execution of small tools which use codec, etc
// to take a long time.
//
// To mitigate, we now support the notfastpath tag.
// This tag disables fastpath during build, allowing for faster build, test execution,
// short-program runs, etc.

func fastpathDecodeTypeSwitch(iv interface{}, d *Decoder) bool      { return false }
func fastpathEncodeTypeSwitch(iv interface{}, e *Encoder) bool      { return false }
func fastpathEncodeTypeSwitchSlice(iv interface{}, e *Encoder) bool { return false }
func fastpathEncodeTypeSwitchMap(iv interface{}, e *Encoder) bool   { return false }
<<<<<<< HEAD
func fastpathDecodeSetZeroTypeSwitch(iv interface{}) bool           { return false }
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

type fastpathT struct{}
type fastpathE struct {
	rtid  uintptr
	rt    reflect.Type
<<<<<<< HEAD
	encfn func(*Encoder, *codecFnInfo, reflect.Value)
	decfn func(*Decoder, *codecFnInfo, reflect.Value)
=======
	encfn func(*encFnInfo, reflect.Value)
	decfn func(*decFnInfo, reflect.Value)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}
type fastpathA [0]fastpathE

func (x fastpathA) index(rtid uintptr) int { return -1 }

<<<<<<< HEAD
func (_ fastpathT) DecSliceUint8V(v []uint8, canChange bool, d *Decoder) (_ []uint8, changed bool) {
	fn := d.cfer().get(uint8SliceTyp, true, true)
	d.kSlice(&fn.i, reflect.ValueOf(&v).Elem())
	return v, true
}

var fastpathAV fastpathA
var fastpathTV fastpathT

// ----
type TestMammoth2Wrapper struct{} // to allow testMammoth work in notfastpath mode
=======
var fastpathAV fastpathA
var fastpathTV fastpathT
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
