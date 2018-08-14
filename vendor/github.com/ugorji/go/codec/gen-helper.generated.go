/* // +build ignore */

// Copyright (c) 2012-2015 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

<<<<<<< HEAD
// Code generated from gen-helper.go.tmpl - DO NOT EDIT.
=======
// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED from gen-helper.go.tmpl
// ************************************************************
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

package codec

import (
	"encoding"
	"reflect"
)

<<<<<<< HEAD
// GenVersion is the current version of codecgen.
const GenVersion = 8

=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// This file is used to generate helper code for codecgen.
// The values here i.e. genHelper(En|De)coder are not to be used directly by
// library users. They WILL change continuously and without notice.
//
// To help enforce this, we create an unexported type with exported members.
// The only way to get the type is via the one exported type that we control (somewhat).
//
// When static codecs are created for types, they will use this value
// to perform encoding or decoding of primitives or known slice or map types.

// GenHelperEncoder is exported so that it can be used externally by codecgen.
<<<<<<< HEAD
//
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperEncoder(e *Encoder) (ge genHelperEncoder, ee genHelperEncDriver) {
	ge = genHelperEncoder{e: e}
	ee = genHelperEncDriver{encDriver: e.e}
	return
}

// GenHelperDecoder is exported so that it can be used externally by codecgen.
//
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperDecoder(d *Decoder) (gd genHelperDecoder, dd genHelperDecDriver) {
	gd = genHelperDecoder{d: d}
	dd = genHelperDecDriver{decDriver: d.d}
	return
}

type genHelperEncDriver struct {
	encDriver
}

func (x genHelperEncDriver) EncodeBuiltin(rt uintptr, v interface{}) {}
func (x genHelperEncDriver) EncStructFieldKey(keyType valueType, s string) {
	encStructFieldKey(x.encDriver, keyType, s)
}
func (x genHelperEncDriver) EncodeSymbol(s string) {
	x.encDriver.EncodeString(cUTF8, s)
}

type genHelperDecDriver struct {
	decDriver
	C checkOverflow
}

func (x genHelperDecDriver) DecodeBuiltin(rt uintptr, v interface{}) {}
func (x genHelperDecDriver) DecStructFieldKey(keyType valueType, buf *[decScratchByteArrayLen]byte) []byte {
	return decStructFieldKey(x.decDriver, keyType, buf)
}
func (x genHelperDecDriver) DecodeInt(bitsize uint8) (i int64) {
	return x.C.IntV(x.decDriver.DecodeInt64(), bitsize)
}
func (x genHelperDecDriver) DecodeUint(bitsize uint8) (ui uint64) {
	return x.C.UintV(x.decDriver.DecodeUint64(), bitsize)
}
func (x genHelperDecDriver) DecodeFloat(chkOverflow32 bool) (f float64) {
	f = x.DecodeFloat64()
	if chkOverflow32 && chkOvf.Float32(f) {
		panicv.errorf("float32 overflow: %v", f)
	}
	return
}
func (x genHelperDecDriver) DecodeFloat32As64() (f float64) {
	f = x.DecodeFloat64()
	if chkOvf.Float32(f) {
		panicv.errorf("float32 overflow: %v", f)
	}
	return
=======
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperEncoder(e *Encoder) (genHelperEncoder, encDriver) {
	return genHelperEncoder{e: e}, e.e
}

// GenHelperDecoder is exported so that it can be used externally by codecgen.
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperDecoder(d *Decoder) (genHelperDecoder, decDriver) {
	return genHelperDecoder{d: d}, d.d
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
type genHelperEncoder struct {
<<<<<<< HEAD
	M must
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	e *Encoder
	F fastpathT
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
type genHelperDecoder struct {
<<<<<<< HEAD
	C checkOverflow
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	d *Decoder
	F fastpathT
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBasicHandle() *BasicHandle {
	return f.e.h
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBinary() bool {
	return f.e.be // f.e.hh.isBinaryEncoding()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
func (f genHelperEncoder) IsJSONHandle() bool {
	return f.e.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncFallback(iv interface{}) {
	// println(">>>>>>>>> EncFallback")
	// f.e.encodeI(iv, false, false)
	f.e.encodeValue(reflect.ValueOf(iv), nil, false)
=======
func (f genHelperEncoder) EncFallback(iv interface{}) {
	// println(">>>>>>>>> EncFallback")
	f.e.encodeI(iv, false, false)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncTextMarshal(iv encoding.TextMarshaler) {
	bs, fnerr := iv.MarshalText()
<<<<<<< HEAD
	f.e.marshal(bs, fnerr, false, cUTF8)
=======
	f.e.marshal(bs, fnerr, false, c_UTF8)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncJSONMarshal(iv jsonMarshaler) {
	bs, fnerr := iv.MarshalJSON()
<<<<<<< HEAD
	f.e.marshal(bs, fnerr, true, cUTF8)
=======
	f.e.marshal(bs, fnerr, true, c_UTF8)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBinaryMarshal(iv encoding.BinaryMarshaler) {
	bs, fnerr := iv.MarshalBinary()
<<<<<<< HEAD
	f.e.marshal(bs, fnerr, false, cRAW)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncRaw(iv Raw) { f.e.rawBytes(iv) }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: builtin no longer supported - so we make this method a no-op,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperEncoder) TimeRtidIfBinc() (v uintptr) { return }

// func (f genHelperEncoder) TimeRtidIfBinc() uintptr {
// 	if _, ok := f.e.hh.(*BincHandle); ok {
// 		return timeTypId
// 	}
// }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) I2Rtid(v interface{}) uintptr {
	return i2rtid(v)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) Extension(rtid uintptr) (xfn *extTypeTagFn) {
	return f.e.h.getExt(rtid)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncExtension(v interface{}, xfFn *extTypeTagFn) {
	f.e.e.EncodeExt(v, xfFn.tag, xfFn.ext, f.e)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
=======
	f.e.marshal(bs, fnerr, false, c_RAW)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncRaw(iv Raw) {
	f.e.raw(iv)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) TimeRtidIfBinc() uintptr {
	if _, ok := f.e.hh.(*BincHandle); ok {
		return timeTypId
	}
	return 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) IsJSONHandle() bool {
	return f.e.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
func (f genHelperEncoder) HasExtensions() bool {
	return len(f.e.h.extHandle) != 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperEncoder) EncExt(v interface{}) (r bool) {
	if xfFn := f.e.h.getExt(i2rtid(v)); xfFn != nil {
=======
func (f genHelperEncoder) EncExt(v interface{}) (r bool) {
	rt := reflect.TypeOf(v)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	rtid := reflect.ValueOf(rt).Pointer()
	if xfFn := f.e.h.getExt(rtid); xfFn != nil {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		f.e.e.EncodeExt(v, xfFn.tag, xfFn.ext, f.e)
		return true
	}
	return false
}

<<<<<<< HEAD
=======
// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncSendContainerState(c containerState) {
	if f.e.cr != nil {
		f.e.cr.sendContainerState(c)
	}
}

>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// ---------------- DECODER FOLLOWS -----------------

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBasicHandle() *BasicHandle {
	return f.d.h
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBinary() bool {
	return f.d.be // f.d.hh.isBinaryEncoding()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
func (f genHelperDecoder) DecSwallow() { f.d.swallow() }
=======
func (f genHelperDecoder) DecSwallow() {
	f.d.swallow()
}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecScratchBuffer() []byte {
	return f.d.b[:]
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
func (f genHelperDecoder) DecScratchArrayBuffer() *[decScratchByteArrayLen]byte {
	return &f.d.b
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecFallback(iv interface{}, chkPtr bool) {
	// println(">>>>>>>>> DecFallback")
	rv := reflect.ValueOf(iv)
	if chkPtr {
		rv = f.d.ensureDecodeable(rv)
	}
	f.d.decodeValue(rv, nil, false)
	// f.d.decodeValueFallback(rv)
=======
func (f genHelperDecoder) DecFallback(iv interface{}, chkPtr bool) {
	// println(">>>>>>>>> DecFallback")
	f.d.decodeI(iv, chkPtr, false, false, false)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecSliceHelperStart() (decSliceHelper, int) {
	return f.d.decSliceHelperStart()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecStructFieldNotFound(index int, name string) {
	f.d.structFieldNotFound(index, name)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecArrayCannotExpand(sliceLen, streamLen int) {
	f.d.arrayCannotExpand(sliceLen, streamLen)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecTextUnmarshal(tm encoding.TextUnmarshaler) {
<<<<<<< HEAD
	fnerr := tm.UnmarshalText(f.d.d.DecodeStringAsBytes())
=======
	fnerr := tm.UnmarshalText(f.d.d.DecodeBytes(f.d.b[:], true, true))
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecJSONUnmarshal(tm jsonUnmarshaler) {
<<<<<<< HEAD
	// bs := f.dd.DecodeStringAsBytes()
=======
	// bs := f.dd.DecodeBytes(f.d.b[:], true, true)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	// grab the bytes to be read, as UnmarshalJSON needs the full JSON so as to unmarshal it itself.
	fnerr := tm.UnmarshalJSON(f.d.nextValueBytes())
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBinaryUnmarshal(bm encoding.BinaryUnmarshaler) {
<<<<<<< HEAD
	fnerr := bm.UnmarshalBinary(f.d.d.DecodeBytes(nil, true))
=======
	fnerr := bm.UnmarshalBinary(f.d.d.DecodeBytes(nil, false, true))
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
func (f genHelperDecoder) DecRaw() []byte { return f.d.rawBytes() }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: builtin no longer supported - so we make this method a no-op,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) TimeRtidIfBinc() (v uintptr) { return }

// func (f genHelperDecoder) TimeRtidIfBinc() uintptr {
// 	// Note: builtin is no longer supported - so make this a no-op
// 	if _, ok := f.d.hh.(*BincHandle); ok {
// 		return timeTypId
// 	}
// 	return 0
// }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) IsJSONHandle() bool {
	return f.d.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) I2Rtid(v interface{}) uintptr {
	return i2rtid(v)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) Extension(rtid uintptr) (xfn *extTypeTagFn) {
	return f.d.h.getExt(rtid)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecExtension(v interface{}, xfFn *extTypeTagFn) {
	f.d.d.DecodeExt(v, xfFn.tag, xfFn.ext)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
=======
func (f genHelperDecoder) DecRaw() []byte {
	return f.d.raw()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) TimeRtidIfBinc() uintptr {
	if _, ok := f.d.hh.(*BincHandle); ok {
		return timeTypId
	}
	return 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) IsJSONHandle() bool {
	return f.d.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
func (f genHelperDecoder) HasExtensions() bool {
	return len(f.d.h.extHandle) != 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) DecExt(v interface{}) (r bool) {
	if xfFn := f.d.h.getExt(i2rtid(v)); xfFn != nil {
=======
func (f genHelperDecoder) DecExt(v interface{}) (r bool) {
	rt := reflect.TypeOf(v).Elem()
	rtid := reflect.ValueOf(rt).Pointer()
	if xfFn := f.d.h.getExt(rtid); xfFn != nil {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		f.d.d.DecodeExt(v, xfFn.tag, xfFn.ext)
		return true
	}
	return false
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
func (f genHelperDecoder) DecInferLen(clen, maxlen, unit int) (rvlen int) {
=======
func (f genHelperDecoder) DecInferLen(clen, maxlen, unit int) (rvlen int, truncated bool) {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	return decInferLen(clen, maxlen, unit)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
<<<<<<< HEAD
//
// Deprecated: no longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) StringView(v []byte) string { return stringView(v) }
=======
func (f genHelperDecoder) DecSendContainerState(c containerState) {
	if f.d.cr != nil {
		f.d.cr.sendContainerState(c)
	}
}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
