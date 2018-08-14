<<<<<<< HEAD
// Copyright (c) 2012-2018 Ugorji Nwoke. All rights reserved.
=======
// Copyright (c) 2012-2015 Ugorji Nwoke. All rights reserved.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

import (
<<<<<<< HEAD
	"bufio"
	"encoding"
	"errors"
=======
	"encoding"
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	"fmt"
	"io"
	"reflect"
	"sort"
<<<<<<< HEAD
	"strconv"
	"sync"
	"time"
)

const defEncByteBufSize = 1 << 6 // 4:16, 6:64, 8:256, 10:1024

var errEncoderNotInitialized = errors.New("Encoder not initialized")
=======
	"sync"
)

const (
	defEncByteBufSize = 1 << 6 // 4:16, 6:64, 8:256, 10:1024
)

// AsSymbolFlag defines what should be encoded as symbols.
type AsSymbolFlag uint8

const (
	// AsSymbolDefault is default.
	// Currently, this means only encode struct field names as symbols.
	// The default is subject to change.
	AsSymbolDefault AsSymbolFlag = iota

	// AsSymbolAll means encode anything which could be a symbol as a symbol.
	AsSymbolAll = 0xfe

	// AsSymbolNone means do not encode anything as a symbol.
	AsSymbolNone = 1 << iota

	// AsSymbolMapStringKeys means encode keys in map[string]XXX as symbols.
	AsSymbolMapStringKeysFlag

	// AsSymbolStructFieldName means encode struct field names as symbols.
	AsSymbolStructFieldNameFlag
)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

// encWriter abstracts writing to a byte array or to an io.Writer.
type encWriter interface {
	writeb([]byte)
	writestr(string)
	writen1(byte)
	writen2(byte, byte)
	atEndOfEncode()
}

// encDriver abstracts the actual codec (binc vs msgpack, etc)
type encDriver interface {
<<<<<<< HEAD
=======
	IsBuiltinType(rt uintptr) bool
	EncodeBuiltin(rt uintptr, v interface{})
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	EncodeNil()
	EncodeInt(i int64)
	EncodeUint(i uint64)
	EncodeBool(b bool)
	EncodeFloat32(f float32)
	EncodeFloat64(f float64)
	// encodeExtPreamble(xtag byte, length int)
	EncodeRawExt(re *RawExt, e *Encoder)
	EncodeExt(v interface{}, xtag uint64, ext Ext, e *Encoder)
<<<<<<< HEAD
	EncodeString(c charEncoding, v string)
	// EncodeSymbol(v string)
	EncodeStringBytes(c charEncoding, v []byte)
	EncodeTime(time.Time)
	//encBignum(f *big.Int)
	//encStringRunes(c charEncoding, v []rune)
	WriteArrayStart(length int)
	WriteArrayElem()
	WriteArrayEnd()
	WriteMapStart(length int)
	WriteMapElemKey()
	WriteMapElemValue()
	WriteMapEnd()

	reset()
	atEndOfEncode()
}

type ioEncStringWriter interface {
	WriteString(s string) (n int, err error)
=======
	EncodeArrayStart(length int)
	EncodeMapStart(length int)
	EncodeString(c charEncoding, v string)
	EncodeSymbol(v string)
	EncodeStringBytes(c charEncoding, v []byte)
	//TODO
	//encBignum(f *big.Int)
	//encStringRunes(c charEncoding, v []rune)

	reset()
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

type encDriverAsis interface {
	EncodeAsis(v []byte)
}

<<<<<<< HEAD
type encDriverNoopContainerWriter struct{}

func (encDriverNoopContainerWriter) WriteArrayStart(length int) {}
func (encDriverNoopContainerWriter) WriteArrayElem()            {}
func (encDriverNoopContainerWriter) WriteArrayEnd()             {}
func (encDriverNoopContainerWriter) WriteMapStart(length int)   {}
func (encDriverNoopContainerWriter) WriteMapElemKey()           {}
func (encDriverNoopContainerWriter) WriteMapElemValue()         {}
func (encDriverNoopContainerWriter) WriteMapEnd()               {}
func (encDriverNoopContainerWriter) atEndOfEncode()             {}

type encDriverTrackContainerWriter struct {
	c containerState
}

func (e *encDriverTrackContainerWriter) WriteArrayStart(length int) { e.c = containerArrayStart }
func (e *encDriverTrackContainerWriter) WriteArrayElem()            { e.c = containerArrayElem }
func (e *encDriverTrackContainerWriter) WriteArrayEnd()             { e.c = containerArrayEnd }
func (e *encDriverTrackContainerWriter) WriteMapStart(length int)   { e.c = containerMapStart }
func (e *encDriverTrackContainerWriter) WriteMapElemKey()           { e.c = containerMapKey }
func (e *encDriverTrackContainerWriter) WriteMapElemValue()         { e.c = containerMapValue }
func (e *encDriverTrackContainerWriter) WriteMapEnd()               { e.c = containerMapEnd }
func (e *encDriverTrackContainerWriter) atEndOfEncode()             {}

// type ioEncWriterWriter interface {
// 	WriteByte(c byte) error
// 	WriteString(s string) (n int, err error)
// 	Write(p []byte) (n int, err error)
// }

// EncodeOptions captures configuration options during encode.
type EncodeOptions struct {
	// WriterBufferSize is the size of the buffer used when writing.
	//
	// if > 0, we use a smart buffer internally for performance purposes.
	WriterBufferSize int

	// ChanRecvTimeout is the timeout used when selecting from a chan.
	//
	// Configuring this controls how we receive from a chan during the encoding process.
	//   - If ==0, we only consume the elements currently available in the chan.
	//   - if  <0, we consume until the chan is closed.
	//   - If  >0, we consume until this timeout.
	ChanRecvTimeout time.Duration

	// StructToArray specifies to encode a struct as an array, and not as a map
=======
type encNoSeparator struct{}

func (_ encNoSeparator) EncodeEnd() {}

type ioEncWriterWriter interface {
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
	Write(p []byte) (n int, err error)
}

type ioEncStringWriter interface {
	WriteString(s string) (n int, err error)
}

type EncodeOptions struct {
	// Encode a struct as an array, and not as a map
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	StructToArray bool

	// Canonical representation means that encoding a value will always result in the same
	// sequence of bytes.
	//
	// This only affects maps, as the iteration order for maps is random.
	//
	// The implementation MAY use the natural sort order for the map keys if possible:
	//
	//     - If there is a natural sort order (ie for number, bool, string or []byte keys),
	//       then the map keys are first sorted in natural order and then written
	//       with corresponding map values to the strema.
	//     - If there is no natural sort order, then the map keys will first be
	//       encoded into []byte, and then sorted,
	//       before writing the sorted keys and the corresponding map values to the stream.
	//
	Canonical bool

	// CheckCircularRef controls whether we check for circular references
	// and error fast during an encode.
	//
	// If enabled, an error is received if a pointer to a struct
	// references itself either directly or through one of its fields (iteratively).
	//
	// This is opt-in, as there may be a performance hit to checking circular references.
	CheckCircularRef bool

	// RecursiveEmptyCheck controls whether we descend into interfaces, structs and pointers
	// when checking if a value is empty.
	//
	// Note that this may make OmitEmpty more expensive, as it incurs a lot more reflect calls.
	RecursiveEmptyCheck bool

	// Raw controls whether we encode Raw values.
	// This is a "dangerous" option and must be explicitly set.
	// If set, we blindly encode Raw values as-is, without checking
	// if they are a correct representation of a value in that format.
	// If unset, we error out.
	Raw bool

<<<<<<< HEAD
	// // AsSymbols defines what should be encoded as symbols.
	// //
	// // Encoding as symbols can reduce the encoded size significantly.
	// //
	// // However, during decoding, each string to be encoded as a symbol must
	// // be checked to see if it has been seen before. Consequently, encoding time
	// // will increase if using symbols, because string comparisons has a clear cost.
	// //
	// // Sample values:
	// //   AsSymbolNone
	// //   AsSymbolAll
	// //   AsSymbolMapStringKeys
	// //   AsSymbolMapStringKeysFlag | AsSymbolStructFieldNameFlag
	// AsSymbols AsSymbolFlag
=======
	// AsSymbols defines what should be encoded as symbols.
	//
	// Encoding as symbols can reduce the encoded size significantly.
	//
	// However, during decoding, each string to be encoded as a symbol must
	// be checked to see if it has been seen before. Consequently, encoding time
	// will increase if using symbols, because string comparisons has a clear cost.
	//
	// Sample values:
	//   AsSymbolNone
	//   AsSymbolAll
	//   AsSymbolMapStringKeys
	//   AsSymbolMapStringKeysFlag | AsSymbolStructFieldNameFlag
	AsSymbols AsSymbolFlag
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// ---------------------------------------------

<<<<<<< HEAD
// ioEncWriter implements encWriter and can write to an io.Writer implementation
type ioEncWriter struct {
	w  io.Writer
	ww io.Writer
	bw io.ByteWriter
	sw ioEncStringWriter
	fw ioFlusher
	b  [8]byte
}

func (z *ioEncWriter) WriteByte(b byte) (err error) {
	z.b[0] = b
	_, err = z.w.Write(z.b[:1])
	return
}

func (z *ioEncWriter) WriteString(s string) (n int, err error) {
	return z.w.Write(bytesView(s))
}

func (z *ioEncWriter) writeb(bs []byte) {
	if _, err := z.ww.Write(bs); err != nil {
		panic(err)
	}
}

func (z *ioEncWriter) writestr(s string) {
	if _, err := z.sw.WriteString(s); err != nil {
		panic(err)
	}
}

func (z *ioEncWriter) writen1(b byte) {
	if err := z.bw.WriteByte(b); err != nil {
=======
type simpleIoEncWriterWriter struct {
	w  io.Writer
	bw io.ByteWriter
	sw ioEncStringWriter
	bs [1]byte
}

func (o *simpleIoEncWriterWriter) WriteByte(c byte) (err error) {
	if o.bw != nil {
		return o.bw.WriteByte(c)
	}
	// _, err = o.w.Write([]byte{c})
	o.bs[0] = c
	_, err = o.w.Write(o.bs[:])
	return
}

func (o *simpleIoEncWriterWriter) WriteString(s string) (n int, err error) {
	if o.sw != nil {
		return o.sw.WriteString(s)
	}
	// return o.w.Write([]byte(s))
	return o.w.Write(bytesView(s))
}

func (o *simpleIoEncWriterWriter) Write(p []byte) (n int, err error) {
	return o.w.Write(p)
}

// ----------------------------------------

// ioEncWriter implements encWriter and can write to an io.Writer implementation
type ioEncWriter struct {
	w ioEncWriterWriter
	s simpleIoEncWriterWriter
	// x [8]byte // temp byte array re-used internally for efficiency
}

func (z *ioEncWriter) writeb(bs []byte) {
	if len(bs) == 0 {
		return
	}
	n, err := z.w.Write(bs)
	if err != nil {
		panic(err)
	}
	if n != len(bs) {
		panic(fmt.Errorf("incorrect num bytes written. Expecting: %v, Wrote: %v", len(bs), n))
	}
}

func (z *ioEncWriter) writestr(s string) {
	n, err := z.w.WriteString(s)
	if err != nil {
		panic(err)
	}
	if n != len(s) {
		panic(fmt.Errorf("incorrect num bytes written. Expecting: %v, Wrote: %v", len(s), n))
	}
}

func (z *ioEncWriter) writen1(b byte) {
	if err := z.w.WriteByte(b); err != nil {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		panic(err)
	}
}

<<<<<<< HEAD
func (z *ioEncWriter) writen2(b1, b2 byte) {
	var err error
	if err = z.bw.WriteByte(b1); err == nil {
		if err = z.bw.WriteByte(b2); err == nil {
			return
		}
	}
	panic(err)
}

// func (z *ioEncWriter) writen5(b1, b2, b3, b4, b5 byte) {
// 	z.b[0], z.b[1], z.b[2], z.b[3], z.b[4] = b1, b2, b3, b4, b5
// 	if _, err := z.ww.Write(z.b[:5]); err != nil {
// 		panic(err)
// 	}
// }

func (z *ioEncWriter) atEndOfEncode() {
	if z.fw != nil {
		if err := z.fw.Flush(); err != nil {
			panic(err)
		}
	}
=======
func (z *ioEncWriter) writen2(b1 byte, b2 byte) {
	z.writen1(b1)
	z.writen1(b2)
}

func (z *ioEncWriter) atEndOfEncode() {}

// ----------------------------------------

// bytesEncWriter implements encWriter and can write to an byte slice.
// It is used by Marshal function.
type bytesEncWriter struct {
	b   []byte
	c   int     // cursor
	out *[]byte // write out on atEndOfEncode
}

func (z *bytesEncWriter) writeb(s []byte) {
	if len(s) == 0 {
		return
	}
	oc, a := z.growNoAlloc(len(s))
	if a {
		z.growAlloc(len(s), oc)
	}
	copy(z.b[oc:], s)
}

func (z *bytesEncWriter) writestr(s string) {
	if len(s) == 0 {
		return
	}
	oc, a := z.growNoAlloc(len(s))
	if a {
		z.growAlloc(len(s), oc)
	}
	copy(z.b[oc:], s)
}

func (z *bytesEncWriter) writen1(b1 byte) {
	oc, a := z.growNoAlloc(1)
	if a {
		z.growAlloc(1, oc)
	}
	z.b[oc] = b1
}

func (z *bytesEncWriter) writen2(b1 byte, b2 byte) {
	oc, a := z.growNoAlloc(2)
	if a {
		z.growAlloc(2, oc)
	}
	z.b[oc+1] = b2
	z.b[oc] = b1
}

func (z *bytesEncWriter) atEndOfEncode() {
	*(z.out) = z.b[:z.c]
}

// have a growNoalloc(n int), which can be inlined.
// if allocation is needed, then call growAlloc(n int)

func (z *bytesEncWriter) growNoAlloc(n int) (oldcursor int, allocNeeded bool) {
	oldcursor = z.c
	z.c = z.c + n
	if z.c > len(z.b) {
		if z.c > cap(z.b) {
			allocNeeded = true
		} else {
			z.b = z.b[:cap(z.b)]
		}
	}
	return
}

func (z *bytesEncWriter) growAlloc(n int, oldcursor int) {
	// appendslice logic (if cap < 1024, *2, else *1.25): more expensive. many copy calls.
	// bytes.Buffer model (2*cap + n): much better
	// bs := make([]byte, 2*cap(z.b)+n)
	bs := make([]byte, growCap(cap(z.b), 1, n))
	copy(bs, z.b[:oldcursor])
	z.b = bs
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// ---------------------------------------------

<<<<<<< HEAD
// bytesEncAppender implements encWriter and can write to an byte slice.
type bytesEncAppender struct {
	b   []byte
	out *[]byte
}

func (z *bytesEncAppender) writeb(s []byte) {
	z.b = append(z.b, s...)
}
func (z *bytesEncAppender) writestr(s string) {
	z.b = append(z.b, s...)
}
func (z *bytesEncAppender) writen1(b1 byte) {
	z.b = append(z.b, b1)
}
func (z *bytesEncAppender) writen2(b1, b2 byte) {
	z.b = append(z.b, b1, b2)
}
func (z *bytesEncAppender) atEndOfEncode() {
	*(z.out) = z.b
}
func (z *bytesEncAppender) reset(in []byte, out *[]byte) {
	z.b = in[:0]
	z.out = out
}

// ---------------------------------------------

func (e *Encoder) rawExt(f *codecFnInfo, rv reflect.Value) {
	e.e.EncodeRawExt(rv2i(rv).(*RawExt), e)
}

func (e *Encoder) ext(f *codecFnInfo, rv reflect.Value) {
	e.e.EncodeExt(rv2i(rv), f.xfTag, f.xfFn, e)
}

func (e *Encoder) selferMarshal(f *codecFnInfo, rv reflect.Value) {
	rv2i(rv).(Selfer).CodecEncodeSelf(e)
}

func (e *Encoder) binaryMarshal(f *codecFnInfo, rv reflect.Value) {
	bs, fnerr := rv2i(rv).(encoding.BinaryMarshaler).MarshalBinary()
	e.marshal(bs, fnerr, false, cRAW)
}

func (e *Encoder) textMarshal(f *codecFnInfo, rv reflect.Value) {
	bs, fnerr := rv2i(rv).(encoding.TextMarshaler).MarshalText()
	e.marshal(bs, fnerr, false, cUTF8)
}

func (e *Encoder) jsonMarshal(f *codecFnInfo, rv reflect.Value) {
	bs, fnerr := rv2i(rv).(jsonMarshaler).MarshalJSON()
	e.marshal(bs, fnerr, true, cUTF8)
}

func (e *Encoder) raw(f *codecFnInfo, rv reflect.Value) {
	e.rawBytes(rv2i(rv).(Raw))
}

func (e *Encoder) kInvalid(f *codecFnInfo, rv reflect.Value) {
	e.e.EncodeNil()
}

func (e *Encoder) kErr(f *codecFnInfo, rv reflect.Value) {
	e.errorf("unsupported kind %s, for %#v", rv.Kind(), rv)
}

func (e *Encoder) kSlice(f *codecFnInfo, rv reflect.Value) {
	ti := f.ti
	ee := e.e
=======
type encFnInfo struct {
	e     *Encoder
	ti    *typeInfo
	xfFn  Ext
	xfTag uint64
	seq   seqType
}

func (f *encFnInfo) builtin(rv reflect.Value) {
	f.e.e.EncodeBuiltin(f.ti.rtid, rv.Interface())
}

func (f *encFnInfo) raw(rv reflect.Value) {
	f.e.raw(rv.Interface().(Raw))
}

func (f *encFnInfo) rawExt(rv reflect.Value) {
	// rev := rv.Interface().(RawExt)
	// f.e.e.EncodeRawExt(&rev, f.e)
	var re *RawExt
	if rv.CanAddr() {
		re = rv.Addr().Interface().(*RawExt)
	} else {
		rev := rv.Interface().(RawExt)
		re = &rev
	}
	f.e.e.EncodeRawExt(re, f.e)
}

func (f *encFnInfo) ext(rv reflect.Value) {
	// if this is a struct|array and it was addressable, then pass the address directly (not the value)
	if k := rv.Kind(); (k == reflect.Struct || k == reflect.Array) && rv.CanAddr() {
		rv = rv.Addr()
	}
	f.e.e.EncodeExt(rv.Interface(), f.xfTag, f.xfFn, f.e)
}

func (f *encFnInfo) getValueForMarshalInterface(rv reflect.Value, indir int8) (v interface{}, proceed bool) {
	if indir == 0 {
		v = rv.Interface()
	} else if indir == -1 {
		// If a non-pointer was passed to Encode(), then that value is not addressable.
		// Take addr if addressable, else copy value to an addressable value.
		if rv.CanAddr() {
			v = rv.Addr().Interface()
		} else {
			rv2 := reflect.New(rv.Type())
			rv2.Elem().Set(rv)
			v = rv2.Interface()
			// fmt.Printf("rv.Type: %v, rv2.Type: %v, v: %v\n", rv.Type(), rv2.Type(), v)
		}
	} else {
		for j := int8(0); j < indir; j++ {
			if rv.IsNil() {
				f.e.e.EncodeNil()
				return
			}
			rv = rv.Elem()
		}
		v = rv.Interface()
	}
	return v, true
}

func (f *encFnInfo) selferMarshal(rv reflect.Value) {
	if v, proceed := f.getValueForMarshalInterface(rv, f.ti.csIndir); proceed {
		v.(Selfer).CodecEncodeSelf(f.e)
	}
}

func (f *encFnInfo) binaryMarshal(rv reflect.Value) {
	if v, proceed := f.getValueForMarshalInterface(rv, f.ti.bmIndir); proceed {
		bs, fnerr := v.(encoding.BinaryMarshaler).MarshalBinary()
		f.e.marshal(bs, fnerr, false, c_RAW)
	}
}

func (f *encFnInfo) textMarshal(rv reflect.Value) {
	if v, proceed := f.getValueForMarshalInterface(rv, f.ti.tmIndir); proceed {
		// debugf(">>>> encoding.TextMarshaler: %T", rv.Interface())
		bs, fnerr := v.(encoding.TextMarshaler).MarshalText()
		f.e.marshal(bs, fnerr, false, c_UTF8)
	}
}

func (f *encFnInfo) jsonMarshal(rv reflect.Value) {
	if v, proceed := f.getValueForMarshalInterface(rv, f.ti.jmIndir); proceed {
		bs, fnerr := v.(jsonMarshaler).MarshalJSON()
		f.e.marshal(bs, fnerr, true, c_UTF8)
	}
}

func (f *encFnInfo) kBool(rv reflect.Value) {
	f.e.e.EncodeBool(rv.Bool())
}

func (f *encFnInfo) kString(rv reflect.Value) {
	f.e.e.EncodeString(c_UTF8, rv.String())
}

func (f *encFnInfo) kFloat64(rv reflect.Value) {
	f.e.e.EncodeFloat64(rv.Float())
}

func (f *encFnInfo) kFloat32(rv reflect.Value) {
	f.e.e.EncodeFloat32(float32(rv.Float()))
}

func (f *encFnInfo) kInt(rv reflect.Value) {
	f.e.e.EncodeInt(rv.Int())
}

func (f *encFnInfo) kUint(rv reflect.Value) {
	f.e.e.EncodeUint(rv.Uint())
}

func (f *encFnInfo) kInvalid(rv reflect.Value) {
	f.e.e.EncodeNil()
}

func (f *encFnInfo) kErr(rv reflect.Value) {
	f.e.errorf("unsupported kind %s, for %#v", rv.Kind(), rv)
}

func (f *encFnInfo) kSlice(rv reflect.Value) {
	ti := f.ti
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	// array may be non-addressable, so we have to manage with care
	//   (don't call rv.Bytes, rv.Slice, etc).
	// E.g. type struct S{B [2]byte};
	//   Encode(S{}) will bomb on "panic: slice of unaddressable array".
<<<<<<< HEAD
	if f.seq != seqTypeArray {
		if rv.IsNil() {
			ee.EncodeNil()
=======
	e := f.e
	if f.seq != seqTypeArray {
		if rv.IsNil() {
			e.e.EncodeNil()
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
			return
		}
		// If in this method, then there was no extension function defined.
		// So it's okay to treat as []byte.
		if ti.rtid == uint8SliceTypId {
<<<<<<< HEAD
			ee.EncodeStringBytes(cRAW, rv.Bytes())
			return
		}
	}
	if f.seq == seqTypeChan && ti.chandir&uint8(reflect.RecvDir) == 0 {
		e.errorf("send-only channel cannot be encoded")
	}
	elemsep := e.esep
	rtelem := ti.elem
	rtelemIsByte := uint8TypId == rt2id(rtelem) // NOT rtelem.Kind() == reflect.Uint8
	var l int
	// if a slice, array or chan of bytes, treat specially
	if rtelemIsByte {
		switch f.seq {
		case seqTypeSlice:
			ee.EncodeStringBytes(cRAW, rv.Bytes())
		case seqTypeArray:
			l = rv.Len()
			if rv.CanAddr() {
				ee.EncodeStringBytes(cRAW, rv.Slice(0, l).Bytes())
=======
			e.e.EncodeStringBytes(c_RAW, rv.Bytes())
			return
		}
	}
	cr := e.cr
	rtelem := ti.rt.Elem()
	l := rv.Len()
	if ti.rtid == uint8SliceTypId || rtelem.Kind() == reflect.Uint8 {
		switch f.seq {
		case seqTypeArray:
			// if l == 0 { e.e.encodeStringBytes(c_RAW, nil) } else
			if rv.CanAddr() {
				e.e.EncodeStringBytes(c_RAW, rv.Slice(0, l).Bytes())
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
			} else {
				var bs []byte
				if l <= cap(e.b) {
					bs = e.b[:l]
				} else {
					bs = make([]byte, l)
				}
				reflect.Copy(reflect.ValueOf(bs), rv)
<<<<<<< HEAD
				ee.EncodeStringBytes(cRAW, bs)
			}
		case seqTypeChan:
			// do not use range, so that the number of elements encoded
			// does not change, and encoding does not hang waiting on someone to close chan.
			// for b := range rv2i(rv).(<-chan byte) { bs = append(bs, b) }
			// ch := rv2i(rv).(<-chan byte) // fix error - that this is a chan byte, not a <-chan byte.

			if rv.IsNil() {
				ee.EncodeNil()
				break
			}
			bs := e.b[:0]
			irv := rv2i(rv)
			ch, ok := irv.(<-chan byte)
			if !ok {
				ch = irv.(chan byte)
			}

		L1:
			switch timeout := e.h.ChanRecvTimeout; {
			case timeout == 0: // only consume available
				for {
					select {
					case b := <-ch:
						bs = append(bs, b)
					default:
						break L1
					}
				}
			case timeout > 0: // consume until timeout
				tt := time.NewTimer(timeout)
				for {
					select {
					case b := <-ch:
						bs = append(bs, b)
					case <-tt.C:
						// close(tt.C)
						break L1
					}
				}
			default: // consume until close
				for b := range ch {
					bs = append(bs, b)
				}
			}

			ee.EncodeStringBytes(cRAW, bs)
=======
				// TODO: Test that reflect.Copy works instead of manual one-by-one
				// for i := 0; i < l; i++ {
				// 	bs[i] = byte(rv.Index(i).Uint())
				// }
				e.e.EncodeStringBytes(c_RAW, bs)
			}
		case seqTypeSlice:
			e.e.EncodeStringBytes(c_RAW, rv.Bytes())
		case seqTypeChan:
			bs := e.b[:0]
			// do not use range, so that the number of elements encoded
			// does not change, and encoding does not hang waiting on someone to close chan.
			// for b := range rv.Interface().(<-chan byte) {
			// 	bs = append(bs, b)
			// }
			ch := rv.Interface().(<-chan byte)
			for i := 0; i < l; i++ {
				bs = append(bs, <-ch)
			}
			e.e.EncodeStringBytes(c_RAW, bs)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		}
		return
	}

<<<<<<< HEAD
	// if chan, consume chan into a slice, and work off that slice.
	var rvcs reflect.Value
	if f.seq == seqTypeChan {
		rvcs = reflect.Zero(reflect.SliceOf(rtelem))
		timeout := e.h.ChanRecvTimeout
		if timeout < 0 { // consume until close
			for {
				recv, recvOk := rv.Recv()
				if !recvOk {
					break
				}
				rvcs = reflect.Append(rvcs, recv)
			}
		} else {
			cases := make([]reflect.SelectCase, 2)
			cases[0] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: rv}
			if timeout == 0 {
				cases[1] = reflect.SelectCase{Dir: reflect.SelectDefault}
			} else {
				tt := time.NewTimer(timeout)
				cases[1] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(tt.C)}
			}
			for {
				chosen, recv, recvOk := reflect.Select(cases)
				if chosen == 1 || !recvOk {
					break
				}
				rvcs = reflect.Append(rvcs, recv)
			}
		}
		rv = rvcs // TODO: ensure this doesn't mess up anywhere that rv of kind chan is expected
	}

	l = rv.Len()
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	if ti.mbs {
		if l%2 == 1 {
			e.errorf("mapBySlice requires even slice length, but got %v", l)
			return
		}
<<<<<<< HEAD
		ee.WriteMapStart(l / 2)
	} else {
		ee.WriteArrayStart(l)
	}

	if l > 0 {
		var fn *codecFn
=======
		e.e.EncodeMapStart(l / 2)
	} else {
		e.e.EncodeArrayStart(l)
	}

	if l > 0 {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		for rtelem.Kind() == reflect.Ptr {
			rtelem = rtelem.Elem()
		}
		// if kind is reflect.Interface, do not pre-determine the
		// encoding type, because preEncodeValue may break it down to
		// a concrete type and kInterface will bomb.
<<<<<<< HEAD
		if rtelem.Kind() != reflect.Interface {
			fn = e.cfer().get(rtelem, true, true)
		}
		for j := 0; j < l; j++ {
			if elemsep {
				if ti.mbs {
					if j%2 == 0 {
						ee.WriteMapElemKey()
					} else {
						ee.WriteMapElemValue()
					}
				} else {
					ee.WriteArrayElem()
				}
			}
			e.encodeValue(rv.Index(j), fn, true)
		}
	}

	if ti.mbs {
		ee.WriteMapEnd()
	} else {
		ee.WriteArrayEnd()
	}
}

func (e *Encoder) kStructNoOmitempty(f *codecFnInfo, rv reflect.Value) {
	fti := f.ti
	elemsep := e.esep
	tisfi := fti.sfiSrc
	toMap := !(fti.toArray || e.h.StructToArray)
	if toMap {
		tisfi = fti.sfiSort
	}
	ee := e.e

	sfn := structFieldNode{v: rv, update: false}
	if toMap {
		ee.WriteMapStart(len(tisfi))
		if elemsep {
			for _, si := range tisfi {
				ee.WriteMapElemKey()
				// ee.EncodeString(cUTF8, si.encName)
				encStructFieldKey(ee, fti.keyType, si.encName)
				ee.WriteMapElemValue()
				e.encodeValue(sfn.field(si), nil, true)
			}
		} else {
			for _, si := range tisfi {
				// ee.EncodeString(cUTF8, si.encName)
				encStructFieldKey(ee, fti.keyType, si.encName)
				e.encodeValue(sfn.field(si), nil, true)
			}
		}
		ee.WriteMapEnd()
	} else {
		ee.WriteArrayStart(len(tisfi))
		if elemsep {
			for _, si := range tisfi {
				ee.WriteArrayElem()
				e.encodeValue(sfn.field(si), nil, true)
			}
		} else {
			for _, si := range tisfi {
				e.encodeValue(sfn.field(si), nil, true)
			}
		}
		ee.WriteArrayEnd()
	}
}

func encStructFieldKey(ee encDriver, keyType valueType, s string) {
	var m must

	// use if-else-if, not switch (which compiles to binary-search)
	// since keyType is typically valueTypeString, branch prediction is pretty good.

	if keyType == valueTypeString {
		ee.EncodeString(cUTF8, s)
	} else if keyType == valueTypeInt {
		ee.EncodeInt(m.Int(strconv.ParseInt(s, 10, 64)))
	} else if keyType == valueTypeUint {
		ee.EncodeUint(m.Uint(strconv.ParseUint(s, 10, 64)))
	} else if keyType == valueTypeFloat {
		ee.EncodeFloat64(m.Float(strconv.ParseFloat(s, 64)))
	} else {
		ee.EncodeString(cUTF8, s)
	}
}

func (e *Encoder) kStruct(f *codecFnInfo, rv reflect.Value) {
	fti := f.ti
	elemsep := e.esep
	tisfi := fti.sfiSrc
	toMap := !(fti.toArray || e.h.StructToArray)
	// if toMap, use the sorted array. If toArray, use unsorted array (to match sequence in struct)
	if toMap {
		tisfi = fti.sfiSort
	}
	newlen := len(fti.sfiSort)
	ee := e.e

	// Use sync.Pool to reduce allocating slices unnecessarily.
	// The cost of sync.Pool is less than the cost of new allocation.
	//
	// Each element of the array pools one of encStructPool(8|16|32|64).
	// It allows the re-use of slices up to 64 in length.
	// A performance cost of encoding structs was collecting
	// which values were empty and should be omitted.
	// We needed slices of reflect.Value and string to collect them.
	// This shared pool reduces the amount of unnecessary creation we do.
	// The cost is that of locking sometimes, but sync.Pool is efficient
	// enough to reduce thread contention.

	var spool *sync.Pool
	var poolv interface{}
	var fkvs []stringRv
	// fmt.Printf(">>>>>>>>>>>>>> encode.kStruct: newlen: %d\n", newlen)
	if newlen <= 8 {
		spool, poolv = pool.stringRv8()
		fkvs = poolv.(*[8]stringRv)[:newlen]
	} else if newlen <= 16 {
		spool, poolv = pool.stringRv16()
		fkvs = poolv.(*[16]stringRv)[:newlen]
	} else if newlen <= 32 {
		spool, poolv = pool.stringRv32()
		fkvs = poolv.(*[32]stringRv)[:newlen]
	} else if newlen <= 64 {
		spool, poolv = pool.stringRv64()
		fkvs = poolv.(*[64]stringRv)[:newlen]
	} else if newlen <= 128 {
		spool, poolv = pool.stringRv128()
		fkvs = poolv.(*[128]stringRv)[:newlen]
	} else {
		fkvs = make([]stringRv, newlen)
	}

	newlen = 0
	var kv stringRv
	recur := e.h.RecursiveEmptyCheck
	sfn := structFieldNode{v: rv, update: false}
	for _, si := range tisfi {
		// kv.r = si.field(rv, false)
		kv.r = sfn.field(si)
		if toMap {
			if si.omitEmpty() && isEmptyValue(kv.r, e.h.TypeInfos, recur, recur) {
=======
		var fn *encFn
		if rtelem.Kind() != reflect.Interface {
			rtelemid := reflect.ValueOf(rtelem).Pointer()
			fn = e.getEncFn(rtelemid, rtelem, true, true)
		}
		// TODO: Consider perf implication of encoding odd index values as symbols if type is string
		for j := 0; j < l; j++ {
			if cr != nil {
				if ti.mbs {
					if j%2 == 0 {
						cr.sendContainerState(containerMapKey)
					} else {
						cr.sendContainerState(containerMapValue)
					}
				} else {
					cr.sendContainerState(containerArrayElem)
				}
			}
			if f.seq == seqTypeChan {
				if rv2, ok2 := rv.Recv(); ok2 {
					e.encodeValue(rv2, fn)
				} else {
					e.encode(nil) // WE HAVE TO DO SOMETHING, so nil if nothing received.
				}
			} else {
				e.encodeValue(rv.Index(j), fn)
			}
		}
	}

	if cr != nil {
		if ti.mbs {
			cr.sendContainerState(containerMapEnd)
		} else {
			cr.sendContainerState(containerArrayEnd)
		}
	}
}

func (f *encFnInfo) kStruct(rv reflect.Value) {
	fti := f.ti
	e := f.e
	cr := e.cr
	tisfi := fti.sfip
	toMap := !(fti.toArray || e.h.StructToArray)
	newlen := len(fti.sfi)

	// Use sync.Pool to reduce allocating slices unnecessarily.
	// The cost of sync.Pool is less than the cost of new allocation.
	pool, poolv, fkvs := encStructPoolGet(newlen)

	// if toMap, use the sorted array. If toArray, use unsorted array (to match sequence in struct)
	if toMap {
		tisfi = fti.sfi
	}
	newlen = 0
	var kv stringRv
	recur := e.h.RecursiveEmptyCheck
	for _, si := range tisfi {
		kv.r = si.field(rv, false)
		if toMap {
			if si.omitEmpty && isEmptyValue(kv.r, recur, recur) {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
				continue
			}
			kv.v = si.encName
		} else {
			// use the zero value.
			// if a reference or struct, set to nil (so you do not output too much)
<<<<<<< HEAD
			if si.omitEmpty() && isEmptyValue(kv.r, e.h.TypeInfos, recur, recur) {
=======
			if si.omitEmpty && isEmptyValue(kv.r, recur, recur) {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
				switch kv.r.Kind() {
				case reflect.Struct, reflect.Interface, reflect.Ptr, reflect.Array, reflect.Map, reflect.Slice:
					kv.r = reflect.Value{} //encode as nil
				}
			}
		}
		fkvs[newlen] = kv
		newlen++
	}

<<<<<<< HEAD
	if toMap {
		ee.WriteMapStart(newlen)
		if elemsep {
			for j := 0; j < newlen; j++ {
				kv = fkvs[j]
				ee.WriteMapElemKey()
				// ee.EncodeString(cUTF8, kv.v)
				encStructFieldKey(ee, fti.keyType, kv.v)
				ee.WriteMapElemValue()
				e.encodeValue(kv.r, nil, true)
			}
		} else {
			for j := 0; j < newlen; j++ {
				kv = fkvs[j]
				// ee.EncodeString(cUTF8, kv.v)
				encStructFieldKey(ee, fti.keyType, kv.v)
				e.encodeValue(kv.r, nil, true)
			}
		}
		ee.WriteMapEnd()
	} else {
		ee.WriteArrayStart(newlen)
		if elemsep {
			for j := 0; j < newlen; j++ {
				ee.WriteArrayElem()
				e.encodeValue(fkvs[j].r, nil, true)
			}
		} else {
			for j := 0; j < newlen; j++ {
				e.encodeValue(fkvs[j].r, nil, true)
			}
		}
		ee.WriteArrayEnd()
=======
	// debugf(">>>> kStruct: newlen: %v", newlen)
	// sep := !e.be
	ee := e.e //don't dereference every time

	if toMap {
		ee.EncodeMapStart(newlen)
		// asSymbols := e.h.AsSymbols&AsSymbolStructFieldNameFlag != 0
		asSymbols := e.h.AsSymbols == AsSymbolDefault || e.h.AsSymbols&AsSymbolStructFieldNameFlag != 0
		for j := 0; j < newlen; j++ {
			kv = fkvs[j]
			if cr != nil {
				cr.sendContainerState(containerMapKey)
			}
			if asSymbols {
				ee.EncodeSymbol(kv.v)
			} else {
				ee.EncodeString(c_UTF8, kv.v)
			}
			if cr != nil {
				cr.sendContainerState(containerMapValue)
			}
			e.encodeValue(kv.r, nil)
		}
		if cr != nil {
			cr.sendContainerState(containerMapEnd)
		}
	} else {
		ee.EncodeArrayStart(newlen)
		for j := 0; j < newlen; j++ {
			kv = fkvs[j]
			if cr != nil {
				cr.sendContainerState(containerArrayElem)
			}
			e.encodeValue(kv.r, nil)
		}
		if cr != nil {
			cr.sendContainerState(containerArrayEnd)
		}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	}

	// do not use defer. Instead, use explicit pool return at end of function.
	// defer has a cost we are trying to avoid.
	// If there is a panic and these slices are not returned, it is ok.
<<<<<<< HEAD
	if spool != nil {
		spool.Put(poolv)
	}
}

func (e *Encoder) kMap(f *codecFnInfo, rv reflect.Value) {
	ee := e.e
=======
	if pool != nil {
		pool.Put(poolv)
	}
}

// func (f *encFnInfo) kPtr(rv reflect.Value) {
// 	debugf(">>>>>>> ??? encode kPtr called - shouldn't get called")
// 	if rv.IsNil() {
// 		f.e.e.encodeNil()
// 		return
// 	}
// 	f.e.encodeValue(rv.Elem())
// }

// func (f *encFnInfo) kInterface(rv reflect.Value) {
// 	println("kInterface called")
// 	debug.PrintStack()
// 	if rv.IsNil() {
// 		f.e.e.EncodeNil()
// 		return
// 	}
// 	f.e.encodeValue(rv.Elem(), nil)
// }

func (f *encFnInfo) kMap(rv reflect.Value) {
	ee := f.e.e
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	if rv.IsNil() {
		ee.EncodeNil()
		return
	}

	l := rv.Len()
<<<<<<< HEAD
	ee.WriteMapStart(l)
	elemsep := e.esep
	if l == 0 {
		ee.WriteMapEnd()
		return
	}
	// var asSymbols bool
=======
	ee.EncodeMapStart(l)
	e := f.e
	cr := e.cr
	if l == 0 {
		if cr != nil {
			cr.sendContainerState(containerMapEnd)
		}
		return
	}
	var asSymbols bool
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	// determine the underlying key and val encFn's for the map.
	// This eliminates some work which is done for each loop iteration i.e.
	// rv.Type(), ref.ValueOf(rt).Pointer(), then check map/list for fn.
	//
	// However, if kind is reflect.Interface, do not pre-determine the
	// encoding type, because preEncodeValue may break it down to
	// a concrete type and kInterface will bomb.
<<<<<<< HEAD
	var keyFn, valFn *codecFn
	ti := f.ti
	rtkey0 := ti.key
	rtkey := rtkey0
	rtval0 := ti.elem
	rtval := rtval0
	// rtkeyid := rt2id(rtkey0)
=======
	var keyFn, valFn *encFn
	ti := f.ti
	rtkey := ti.rt.Key()
	rtval := ti.rt.Elem()
	rtkeyid := reflect.ValueOf(rtkey).Pointer()
	// keyTypeIsString := f.ti.rt.Key().Kind() == reflect.String
	var keyTypeIsString = rtkeyid == stringTypId
	if keyTypeIsString {
		asSymbols = e.h.AsSymbols&AsSymbolMapStringKeysFlag != 0
	} else {
		for rtkey.Kind() == reflect.Ptr {
			rtkey = rtkey.Elem()
		}
		if rtkey.Kind() != reflect.Interface {
			rtkeyid = reflect.ValueOf(rtkey).Pointer()
			keyFn = e.getEncFn(rtkeyid, rtkey, true, true)
		}
	}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	for rtval.Kind() == reflect.Ptr {
		rtval = rtval.Elem()
	}
	if rtval.Kind() != reflect.Interface {
<<<<<<< HEAD
		valFn = e.cfer().get(rtval, true, true)
	}
	mks := rv.MapKeys()

	if e.h.Canonical {
		e.kMapCanonical(rtkey, rv, mks, valFn)
		ee.WriteMapEnd()
		return
	}

	var keyTypeIsString = stringTypId == rt2id(rtkey0) // rtkeyid
	if !keyTypeIsString {
		for rtkey.Kind() == reflect.Ptr {
			rtkey = rtkey.Elem()
		}
		if rtkey.Kind() != reflect.Interface {
			// rtkeyid = rt2id(rtkey)
			keyFn = e.cfer().get(rtkey, true, true)
		}
	}

	// for j, lmks := 0, len(mks); j < lmks; j++ {
	for j := range mks {
		if elemsep {
			ee.WriteMapElemKey()
		}
		if keyTypeIsString {
			ee.EncodeString(cUTF8, mks[j].String())
		} else {
			e.encodeValue(mks[j], keyFn, true)
		}
		if elemsep {
			ee.WriteMapElemValue()
		}
		e.encodeValue(rv.MapIndex(mks[j]), valFn, true)

	}
	ee.WriteMapEnd()
}

func (e *Encoder) kMapCanonical(rtkey reflect.Type, rv reflect.Value, mks []reflect.Value, valFn *codecFn) {
	ee := e.e
	elemsep := e.esep
	// we previously did out-of-band if an extension was registered.
	// This is not necessary, as the natural kind is sufficient for ordering.

	switch rtkey.Kind() {
	case reflect.Bool:
		mksv := make([]boolRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Bool()
		}
		sort.Sort(boolRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeBool(mksv[i].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.String:
		mksv := make([]stringRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.String()
		}
		sort.Sort(stringRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeString(cUTF8, mksv[i].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint, reflect.Uintptr:
		mksv := make([]uintRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Uint()
		}
		sort.Sort(uintRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeUint(mksv[i].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		mksv := make([]intRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Int()
		}
		sort.Sort(intRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeInt(mksv[i].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.Float32:
		mksv := make([]floatRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Float()
		}
		sort.Sort(floatRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeFloat32(float32(mksv[i].v))
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.Float64:
		mksv := make([]floatRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Float()
		}
		sort.Sort(floatRvSlice(mksv))
		for i := range mksv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			ee.EncodeFloat64(mksv[i].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
		}
	case reflect.Struct:
		if rv.Type() == timeTyp {
			mksv := make([]timeRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = rv2i(k).(time.Time)
			}
			sort.Sort(timeRvSlice(mksv))
			for i := range mksv {
				if elemsep {
					ee.WriteMapElemKey()
				}
				ee.EncodeTime(mksv[i].v)
				if elemsep {
					ee.WriteMapElemValue()
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn, true)
			}
			break
		}
		fallthrough
	default:
		// out-of-band
		// first encode each key to a []byte first, then sort them, then record
		var mksv []byte = make([]byte, 0, len(mks)*16) // temporary byte slice for the encoding
		e2 := NewEncoderBytes(&mksv, e.hh)
		mksbv := make([]bytesRv, len(mks))
		for i, k := range mks {
			v := &mksbv[i]
			l := len(mksv)
			e2.MustEncode(k)
			v.r = k
			v.v = mksv[l:]
		}
		sort.Sort(bytesRvSlice(mksbv))
		for j := range mksbv {
			if elemsep {
				ee.WriteMapElemKey()
			}
			e.asis(mksbv[j].v)
			if elemsep {
				ee.WriteMapElemValue()
			}
			e.encodeValue(rv.MapIndex(mksbv[j].r), valFn, true)
=======
		rtvalid := reflect.ValueOf(rtval).Pointer()
		valFn = e.getEncFn(rtvalid, rtval, true, true)
	}
	mks := rv.MapKeys()
	// for j, lmks := 0, len(mks); j < lmks; j++ {

	if e.h.Canonical {
		e.kMapCanonical(rtkeyid, rtkey, rv, mks, valFn, asSymbols)
	} else {
		for j := range mks {
			if cr != nil {
				cr.sendContainerState(containerMapKey)
			}
			if keyTypeIsString {
				if asSymbols {
					ee.EncodeSymbol(mks[j].String())
				} else {
					ee.EncodeString(c_UTF8, mks[j].String())
				}
			} else {
				e.encodeValue(mks[j], keyFn)
			}
			if cr != nil {
				cr.sendContainerState(containerMapValue)
			}
			e.encodeValue(rv.MapIndex(mks[j]), valFn)
		}
	}
	if cr != nil {
		cr.sendContainerState(containerMapEnd)
	}
}

func (e *Encoder) kMapCanonical(rtkeyid uintptr, rtkey reflect.Type, rv reflect.Value, mks []reflect.Value, valFn *encFn, asSymbols bool) {
	ee := e.e
	cr := e.cr
	// we previously did out-of-band if an extension was registered.
	// This is not necessary, as the natural kind is sufficient for ordering.

	if rtkeyid == uint8SliceTypId {
		mksv := make([]bytesRv, len(mks))
		for i, k := range mks {
			v := &mksv[i]
			v.r = k
			v.v = k.Bytes()
		}
		sort.Sort(bytesRvSlice(mksv))
		for i := range mksv {
			if cr != nil {
				cr.sendContainerState(containerMapKey)
			}
			ee.EncodeStringBytes(c_RAW, mksv[i].v)
			if cr != nil {
				cr.sendContainerState(containerMapValue)
			}
			e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
		}
	} else {
		switch rtkey.Kind() {
		case reflect.Bool:
			mksv := make([]boolRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.Bool()
			}
			sort.Sort(boolRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				ee.EncodeBool(mksv[i].v)
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		case reflect.String:
			mksv := make([]stringRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.String()
			}
			sort.Sort(stringRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				if asSymbols {
					ee.EncodeSymbol(mksv[i].v)
				} else {
					ee.EncodeString(c_UTF8, mksv[i].v)
				}
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint, reflect.Uintptr:
			mksv := make([]uintRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.Uint()
			}
			sort.Sort(uintRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				ee.EncodeUint(mksv[i].v)
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			mksv := make([]intRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.Int()
			}
			sort.Sort(intRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				ee.EncodeInt(mksv[i].v)
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		case reflect.Float32:
			mksv := make([]floatRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.Float()
			}
			sort.Sort(floatRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				ee.EncodeFloat32(float32(mksv[i].v))
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		case reflect.Float64:
			mksv := make([]floatRv, len(mks))
			for i, k := range mks {
				v := &mksv[i]
				v.r = k
				v.v = k.Float()
			}
			sort.Sort(floatRvSlice(mksv))
			for i := range mksv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				ee.EncodeFloat64(mksv[i].v)
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksv[i].r), valFn)
			}
		default:
			// out-of-band
			// first encode each key to a []byte first, then sort them, then record
			var mksv []byte = make([]byte, 0, len(mks)*16) // temporary byte slice for the encoding
			e2 := NewEncoderBytes(&mksv, e.hh)
			mksbv := make([]bytesRv, len(mks))
			for i, k := range mks {
				v := &mksbv[i]
				l := len(mksv)
				e2.MustEncode(k)
				v.r = k
				v.v = mksv[l:]
				// fmt.Printf(">>>>> %s\n", mksv[l:])
			}
			sort.Sort(bytesRvSlice(mksbv))
			for j := range mksbv {
				if cr != nil {
					cr.sendContainerState(containerMapKey)
				}
				e.asis(mksbv[j].v)
				if cr != nil {
					cr.sendContainerState(containerMapValue)
				}
				e.encodeValue(rv.MapIndex(mksbv[j].r), valFn)
			}
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		}
	}
}

<<<<<<< HEAD
// // --------------------------------------------------

type encWriterSwitch struct {
	wi *ioEncWriter
	// wb bytesEncWriter
	wb   bytesEncAppender
	wx   bool // if bytes, wx=true
	esep bool // whether it has elem separators
	isas bool // whether e.as != nil
}

// // TODO: Uncomment after mid-stack inlining enabled in go 1.11

// func (z *encWriterSwitch) writeb(s []byte) {
// 	if z.wx {
// 		z.wb.writeb(s)
// 	} else {
// 		z.wi.writeb(s)
// 	}
// }
// func (z *encWriterSwitch) writestr(s string) {
// 	if z.wx {
// 		z.wb.writestr(s)
// 	} else {
// 		z.wi.writestr(s)
// 	}
// }
// func (z *encWriterSwitch) writen1(b1 byte) {
// 	if z.wx {
// 		z.wb.writen1(b1)
// 	} else {
// 		z.wi.writen1(b1)
// 	}
// }
// func (z *encWriterSwitch) writen2(b1, b2 byte) {
// 	if z.wx {
// 		z.wb.writen2(b1, b2)
// 	} else {
// 		z.wi.writen2(b1, b2)
// 	}
// }

// An Encoder writes an object to an output stream in the codec format.
type Encoder struct {
	panicHdl
=======
// --------------------------------------------------

// encFn encapsulates the captured variables and the encode function.
// This way, we only do some calculations one times, and pass to the
// code block that should be called (encapsulated in a function)
// instead of executing the checks every time.
type encFn struct {
	i encFnInfo
	f func(*encFnInfo, reflect.Value)
}

// --------------------------------------------------

type encRtidFn struct {
	rtid uintptr
	fn   encFn
}

// An Encoder writes an object to an output stream in the codec format.
type Encoder struct {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	// hopefully, reduce derefencing cost by laying the encWriter inside the Encoder
	e encDriver
	// NOTE: Encoder shouldn't call it's write methods,
	// as the handler MAY need to do some coordination.
<<<<<<< HEAD
	w encWriter

	h  *BasicHandle
	bw *bufio.Writer
	as encDriverAsis

	// ---- cpu cache line boundary?

	// ---- cpu cache line boundary?
	encWriterSwitch
	err error

	// ---- cpu cache line boundary?
	codecFnPooler
	ci set
	js bool    // here, so that no need to piggy back on *codecFner for this
	be bool    // here, so that no need to piggy back on *codecFner for this
	_  [6]byte // padding

	// ---- writable fields during execution --- *try* to keep in sep cache line

	// ---- cpu cache line boundary?
	// b [scratchByteArrayLen]byte
	// _ [cacheLineSize - scratchByteArrayLen]byte // padding
	b [cacheLineSize - 0]byte // used for encoding a chan or (non-addressable) array of bytes
=======
	w  encWriter
	s  []encRtidFn
	ci set
	be bool // is binary encoding
	js bool // is json handle

	wi ioEncWriter
	wb bytesEncWriter

	h  *BasicHandle
	hh Handle

	cr containerStateRecv
	as encDriverAsis

	f map[uintptr]*encFn
	b [scratchByteArrayLen]byte
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

// NewEncoder returns an Encoder for encoding into an io.Writer.
//
// For efficiency, Users are encouraged to pass in a memory buffered writer
// (eg bufio.Writer, bytes.Buffer).
func NewEncoder(w io.Writer, h Handle) *Encoder {
	e := newEncoder(h)
	e.Reset(w)
	return e
}

// NewEncoderBytes returns an encoder for encoding directly and efficiently
// into a byte slice, using zero-copying to temporary slices.
//
// It will potentially replace the output byte slice pointed to.
// After encoding, the out parameter contains the encoded contents.
func NewEncoderBytes(out *[]byte, h Handle) *Encoder {
	e := newEncoder(h)
	e.ResetBytes(out)
	return e
}

func newEncoder(h Handle) *Encoder {
<<<<<<< HEAD
	e := &Encoder{h: h.getBasicHandle(), err: errEncoderNotInitialized}
	e.hh = h
	e.esep = h.hasElemSeparators()
	return e
}

func (e *Encoder) resetCommon() {
	if e.e == nil || e.hh.recreateEncDriver(e.e) {
		e.e = e.hh.newEncDriver(e)
		e.as, e.isas = e.e.(encDriverAsis)
		// e.cr, _ = e.e.(containerStateRecv)
	}
	e.be = e.hh.isBinary()
	_, e.js = e.hh.(*JsonHandle)
	e.e.reset()
	e.err = nil
}

// Reset resets the Encoder with a new output stream.
=======
	e := &Encoder{hh: h, h: h.getBasicHandle(), be: h.isBinary()}
	_, e.js = h.(*JsonHandle)
	e.e = h.newEncDriver(e)
	e.as, _ = e.e.(encDriverAsis)
	e.cr, _ = e.e.(containerStateRecv)
	return e
}

// Reset the Encoder with a new output stream.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
//
// This accommodates using the state of the Encoder,
// where it has "cached" information about sub-engines.
func (e *Encoder) Reset(w io.Writer) {
<<<<<<< HEAD
	if w == nil {
		return
	}
	if e.wi == nil {
		e.wi = new(ioEncWriter)
	}
	var ok bool
	e.wx = false
	e.wi.w = w
	if e.h.WriterBufferSize > 0 {
		e.bw = bufio.NewWriterSize(w, e.h.WriterBufferSize)
		e.wi.bw = e.bw
		e.wi.sw = e.bw
		e.wi.fw = e.bw
		e.wi.ww = e.bw
	} else {
		if e.wi.bw, ok = w.(io.ByteWriter); !ok {
			e.wi.bw = e.wi
		}
		if e.wi.sw, ok = w.(ioEncStringWriter); !ok {
			e.wi.sw = e.wi
		}
		e.wi.fw, _ = w.(ioFlusher)
		e.wi.ww = w
	}
	e.w = e.wi
	e.resetCommon()
}

// ResetBytes resets the Encoder with a new destination output []byte.
func (e *Encoder) ResetBytes(out *[]byte) {
	if out == nil {
		return
	}
	var in []byte
	if out != nil {
		in = *out
	}
	if in == nil {
		in = make([]byte, defEncByteBufSize)
	}
	e.wx = true
	e.wb.reset(in, out)
	e.w = &e.wb
	e.resetCommon()
}

// Encode writes an object into a stream.
//
// Encoding can be configured via the struct tag for the fields.
// The key (in the struct tags) that we look at is configurable.
//
// By default, we look up the "codec" key in the struct field's tags,
// and fall bak to the "json" key if "codec" is absent.
// That key in struct field's tag value is the key name,
// followed by an optional comma and options.
//
// To set an option on all fields (e.g. omitempty on all fields), you
// can create a field called _struct, and set flags on it. The options
// which can be set on _struct are:
//    - omitempty: so all fields are omitted if empty
//    - toarray: so struct is encoded as an array
//    - int: so struct key names are encoded as signed integers (instead of strings)
//    - uint: so struct key names are encoded as unsigned integers (instead of strings)
//    - float: so struct key names are encoded as floats (instead of strings)
// More details on these below.
=======
	ww, ok := w.(ioEncWriterWriter)
	if ok {
		e.wi.w = ww
	} else {
		sww := &e.wi.s
		sww.w = w
		sww.bw, _ = w.(io.ByteWriter)
		sww.sw, _ = w.(ioEncStringWriter)
		e.wi.w = sww
		//ww = bufio.NewWriterSize(w, defEncByteBufSize)
	}
	e.w = &e.wi
	e.e.reset()
}

func (e *Encoder) ResetBytes(out *[]byte) {
	in := *out
	if in == nil {
		in = make([]byte, defEncByteBufSize)
	}
	e.wb.b, e.wb.out, e.wb.c = in, out, 0
	e.w = &e.wb
	e.e.reset()
}

// func (e *Encoder) sendContainerState(c containerState) {
// 	if e.cr != nil {
// 		e.cr.sendContainerState(c)
// 	}
// }

// Encode writes an object into a stream.
//
// Encoding can be configured via the struct tag for the fields.
// The "codec" key in struct field's tag value is the key name,
// followed by an optional comma and options.
// Note that the "json" key is used in the absence of the "codec" key.
//
// To set an option on all fields (e.g. omitempty on all fields), you
// can create a field called _struct, and set flags on it.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
//
// Struct values "usually" encode as maps. Each exported struct field is encoded unless:
//    - the field's tag is "-", OR
//    - the field is empty (empty or the zero value) and its tag specifies the "omitempty" option.
//
// When encoding as a map, the first string in the tag (before the comma)
// is the map key string to use when encoding.
<<<<<<< HEAD
// ...
// This key is typically encoded as a string.
// However, there are instances where the encoded stream has mapping keys encoded as numbers.
// For example, some cbor streams have keys as integer codes in the stream, but they should map
// to fields in a structured object. Consequently, a struct is the natural representation in code.
// For these, configure the struct to encode/decode the keys as numbers (instead of string).
// This is done with the int,uint or float option on the _struct field (see above).
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
//
// However, struct values may encode as arrays. This happens when:
//    - StructToArray Encode option is set, OR
//    - the tag on the _struct field sets the "toarray" option
// Note that omitempty is ignored when encoding struct values as arrays,
// as an entry must be encoded for each field, to maintain its position.
//
// Values with types that implement MapBySlice are encoded as stream maps.
//
// The empty values (for omitempty option) are false, 0, any nil pointer
// or interface value, and any array, slice, map, or string of length zero.
//
// Anonymous fields are encoded inline except:
//    - the struct tag specifies a replacement name (first value)
//    - the field is of an interface type
//
// Examples:
//
//      // NOTE: 'json:' can be used as struct tag key, in place 'codec:' below.
//      type MyStruct struct {
//          _struct bool    `codec:",omitempty"`   //set omitempty for every field
//          Field1 string   `codec:"-"`            //skip this field
//          Field2 int      `codec:"myName"`       //Use key "myName" in encode stream
//          Field3 int32    `codec:",omitempty"`   //use key "Field3". Omit if empty.
//          Field4 bool     `codec:"f4,omitempty"` //use key "f4". Omit if empty.
//          io.Reader                              //use key "Reader".
//          MyStruct        `codec:"my1"           //use key "my1".
//          MyStruct                               //inline it
//          ...
//      }
//
//      type MyStruct struct {
<<<<<<< HEAD
//          _struct bool    `codec:",toarray"`     //encode struct as an array
//      }
//
//      type MyStruct struct {
//          _struct bool    `codec:",uint"`        //encode struct with "unsigned integer" keys
//          Field1 string   `codec:"1"`            //encode Field1 key using: EncodeInt(1)
//          Field2 string   `codec:"2"`            //encode Field2 key using: EncodeInt(2)
=======
//          _struct bool    `codec:",toarray"`   //encode struct as an array
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
//      }
//
// The mode of encoding is based on the type of the value. When a value is seen:
//   - If a Selfer, call its CodecEncodeSelf method
//   - If an extension is registered for it, call that extension function
<<<<<<< HEAD
//   - If implements encoding.(Binary|Text|JSON)Marshaler, call Marshal(Binary|Text|JSON) method
=======
//   - If it implements encoding.(Binary|Text|JSON)Marshaler, call its Marshal(Binary|Text|JSON) method
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
//   - Else encode it based on its reflect.Kind
//
// Note that struct field names and keys in map[string]XXX will be treated as symbols.
// Some formats support symbols (e.g. binc) and will properly encode the string
// only once in the stream, and use a tag to refer to it thereafter.
func (e *Encoder) Encode(v interface{}) (err error) {
<<<<<<< HEAD
	defer e.deferred(&err)
	e.MustEncode(v)
=======
	defer panicToErr(&err)
	e.encode(v)
	e.w.atEndOfEncode()
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	return
}

// MustEncode is like Encode, but panics if unable to Encode.
// This provides insight to the code location that triggered the error.
func (e *Encoder) MustEncode(v interface{}) {
<<<<<<< HEAD
	if e.err != nil {
		panic(e.err)
	}
	e.encode(v)
	e.e.atEndOfEncode()
	e.w.atEndOfEncode()
	e.alwaysAtEnd()
}

func (e *Encoder) deferred(err1 *error) {
	e.alwaysAtEnd()
	if recoverPanicToErr {
		if x := recover(); x != nil {
			panicValToErr(e, x, err1)
			panicValToErr(e, x, &e.err)
		}
	}
}

// func (e *Encoder) alwaysAtEnd() {
// 	e.codecFnPooler.alwaysAtEnd()
// }

func (e *Encoder) encode(iv interface{}) {
	if iv == nil || definitelyNil(iv) {
		e.e.EncodeNil()
		return
	}
	if v, ok := iv.(Selfer); ok {
		v.CodecEncodeSelf(e)
		return
	}

	// a switch with only concrete types can be optimized.
	// consequently, we deal with nil and interfaces outside.

	switch v := iv.(type) {
	case Raw:
		e.rawBytes(v)
	case reflect.Value:
		e.encodeValue(v, nil, true)

	case string:
		e.e.EncodeString(cUTF8, v)
=======
	e.encode(v)
	e.w.atEndOfEncode()
}

func (e *Encoder) encode(iv interface{}) {
	// if ics, ok := iv.(Selfer); ok {
	// 	ics.CodecEncodeSelf(e)
	// 	return
	// }

	switch v := iv.(type) {
	case nil:
		e.e.EncodeNil()
	case Selfer:
		v.CodecEncodeSelf(e)
	case Raw:
		e.raw(v)
	case reflect.Value:
		e.encodeValue(v, nil)

	case string:
		e.e.EncodeString(c_UTF8, v)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	case bool:
		e.e.EncodeBool(v)
	case int:
		e.e.EncodeInt(int64(v))
	case int8:
		e.e.EncodeInt(int64(v))
	case int16:
		e.e.EncodeInt(int64(v))
	case int32:
		e.e.EncodeInt(int64(v))
	case int64:
		e.e.EncodeInt(v)
	case uint:
		e.e.EncodeUint(uint64(v))
	case uint8:
		e.e.EncodeUint(uint64(v))
	case uint16:
		e.e.EncodeUint(uint64(v))
	case uint32:
		e.e.EncodeUint(uint64(v))
	case uint64:
		e.e.EncodeUint(v)
<<<<<<< HEAD
	case uintptr:
		e.e.EncodeUint(uint64(v))
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	case float32:
		e.e.EncodeFloat32(v)
	case float64:
		e.e.EncodeFloat64(v)
<<<<<<< HEAD
	case time.Time:
		e.e.EncodeTime(v)
	case []uint8:
		e.e.EncodeStringBytes(cRAW, v)

	case *Raw:
		e.rawBytes(*v)

	case *string:
		e.e.EncodeString(cUTF8, *v)
=======

	case []uint8:
		e.e.EncodeStringBytes(c_RAW, v)

	case *string:
		e.e.EncodeString(c_UTF8, *v)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	case *bool:
		e.e.EncodeBool(*v)
	case *int:
		e.e.EncodeInt(int64(*v))
	case *int8:
		e.e.EncodeInt(int64(*v))
	case *int16:
		e.e.EncodeInt(int64(*v))
	case *int32:
		e.e.EncodeInt(int64(*v))
	case *int64:
		e.e.EncodeInt(*v)
	case *uint:
		e.e.EncodeUint(uint64(*v))
	case *uint8:
		e.e.EncodeUint(uint64(*v))
	case *uint16:
		e.e.EncodeUint(uint64(*v))
	case *uint32:
		e.e.EncodeUint(uint64(*v))
	case *uint64:
		e.e.EncodeUint(*v)
<<<<<<< HEAD
	case *uintptr:
		e.e.EncodeUint(uint64(*v))
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	case *float32:
		e.e.EncodeFloat32(*v)
	case *float64:
		e.e.EncodeFloat64(*v)
<<<<<<< HEAD
	case *time.Time:
		e.e.EncodeTime(*v)

	case *[]uint8:
		e.e.EncodeStringBytes(cRAW, *v)

	default:
		if !fastpathEncodeTypeSwitch(iv, e) {
			// checkfastpath=true (not false), as underlying slice/map type may be fast-path
			e.encodeValue(reflect.ValueOf(iv), nil, true)
=======

	case *[]uint8:
		e.e.EncodeStringBytes(c_RAW, *v)

	default:
		const checkCodecSelfer1 = true // in case T is passed, where *T is a Selfer, still checkCodecSelfer
		if !fastpathEncodeTypeSwitch(iv, e) {
			e.encodeI(iv, false, checkCodecSelfer1)
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		}
	}
}

<<<<<<< HEAD
func (e *Encoder) encodeValue(rv reflect.Value, fn *codecFn, checkFastpath bool) {
	// if a valid fn is passed, it MUST BE for the dereferenced type of rv
	var sptr uintptr
	var rvp reflect.Value
	var rvpValid bool
=======
func (e *Encoder) preEncodeValue(rv reflect.Value) (rv2 reflect.Value, sptr uintptr, proceed bool) {
	// use a goto statement instead of a recursive function for ptr/interface.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
TOP:
	switch rv.Kind() {
	case reflect.Ptr:
		if rv.IsNil() {
			e.e.EncodeNil()
			return
		}
<<<<<<< HEAD
		rvpValid = true
		rvp = rv
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
		rv = rv.Elem()
		if e.h.CheckCircularRef && rv.Kind() == reflect.Struct {
			// TODO: Movable pointers will be an issue here. Future problem.
			sptr = rv.UnsafeAddr()
			break TOP
		}
		goto TOP
	case reflect.Interface:
		if rv.IsNil() {
			e.e.EncodeNil()
			return
		}
		rv = rv.Elem()
		goto TOP
	case reflect.Slice, reflect.Map:
		if rv.IsNil() {
			e.e.EncodeNil()
			return
		}
	case reflect.Invalid, reflect.Func:
		e.e.EncodeNil()
		return
	}

<<<<<<< HEAD
	if sptr != 0 && (&e.ci).add(sptr) {
		e.errorf("circular reference found: # %d", sptr)
	}

	if fn == nil {
		rt := rv.Type()
		// always pass checkCodecSelfer=true, in case T or ****T is passed, where *T is a Selfer
		fn = e.cfer().get(rt, checkFastpath, true)
	}
	if fn.i.addrE {
		if rvpValid {
			fn.fe(e, &fn.i, rvp)
		} else if rv.CanAddr() {
			fn.fe(e, &fn.i, rv.Addr())
		} else {
			rv2 := reflect.New(rv.Type())
			rv2.Elem().Set(rv)
			fn.fe(e, &fn.i, rv2)
		}
	} else {
		fn.fe(e, &fn.i, rv)
	}
	if sptr != 0 {
		(&e.ci).remove(sptr)
	}
=======
	proceed = true
	rv2 = rv
	return
}

func (e *Encoder) doEncodeValue(rv reflect.Value, fn *encFn, sptr uintptr,
	checkFastpath, checkCodecSelfer bool) {
	if sptr != 0 {
		if (&e.ci).add(sptr) {
			e.errorf("circular reference found: # %d", sptr)
		}
	}
	if fn == nil {
		rt := rv.Type()
		rtid := reflect.ValueOf(rt).Pointer()
		// fn = e.getEncFn(rtid, rt, true, true)
		fn = e.getEncFn(rtid, rt, checkFastpath, checkCodecSelfer)
	}
	fn.f(&fn.i, rv)
	if sptr != 0 {
		(&e.ci).remove(sptr)
	}
}

func (e *Encoder) encodeI(iv interface{}, checkFastpath, checkCodecSelfer bool) {
	if rv, sptr, proceed := e.preEncodeValue(reflect.ValueOf(iv)); proceed {
		e.doEncodeValue(rv, nil, sptr, checkFastpath, checkCodecSelfer)
	}
}

func (e *Encoder) encodeValue(rv reflect.Value, fn *encFn) {
	// if a valid fn is passed, it MUST BE for the dereferenced type of rv
	if rv, sptr, proceed := e.preEncodeValue(rv); proceed {
		e.doEncodeValue(rv, fn, sptr, true, true)
	}
}

func (e *Encoder) getEncFn(rtid uintptr, rt reflect.Type, checkFastpath, checkCodecSelfer bool) (fn *encFn) {
	// rtid := reflect.ValueOf(rt).Pointer()
	var ok bool
	if useMapForCodecCache {
		fn, ok = e.f[rtid]
	} else {
		for i := range e.s {
			v := &(e.s[i])
			if v.rtid == rtid {
				fn, ok = &(v.fn), true
				break
			}
		}
	}
	if ok {
		return
	}

	if useMapForCodecCache {
		if e.f == nil {
			e.f = make(map[uintptr]*encFn, initCollectionCap)
		}
		fn = new(encFn)
		e.f[rtid] = fn
	} else {
		if e.s == nil {
			e.s = make([]encRtidFn, 0, initCollectionCap)
		}
		e.s = append(e.s, encRtidFn{rtid: rtid})
		fn = &(e.s[len(e.s)-1]).fn
	}

	ti := e.h.getTypeInfo(rtid, rt)
	fi := &(fn.i)
	fi.e = e
	fi.ti = ti

	if checkCodecSelfer && ti.cs {
		fn.f = (*encFnInfo).selferMarshal
	} else if rtid == rawTypId {
		fn.f = (*encFnInfo).raw
	} else if rtid == rawExtTypId {
		fn.f = (*encFnInfo).rawExt
	} else if e.e.IsBuiltinType(rtid) {
		fn.f = (*encFnInfo).builtin
	} else if xfFn := e.h.getExt(rtid); xfFn != nil {
		fi.xfTag, fi.xfFn = xfFn.tag, xfFn.ext
		fn.f = (*encFnInfo).ext
	} else if supportMarshalInterfaces && e.be && ti.bm {
		fn.f = (*encFnInfo).binaryMarshal
	} else if supportMarshalInterfaces && !e.be && e.js && ti.jm {
		//If JSON, we should check JSONMarshal before textMarshal
		fn.f = (*encFnInfo).jsonMarshal
	} else if supportMarshalInterfaces && !e.be && ti.tm {
		fn.f = (*encFnInfo).textMarshal
	} else {
		rk := rt.Kind()
		if fastpathEnabled && checkFastpath && (rk == reflect.Map || rk == reflect.Slice) {
			if rt.PkgPath() == "" { // un-named slice or map
				if idx := fastpathAV.index(rtid); idx != -1 {
					fn.f = fastpathAV[idx].encfn
				}
			} else {
				ok = false
				// use mapping for underlying type if there
				var rtu reflect.Type
				if rk == reflect.Map {
					rtu = reflect.MapOf(rt.Key(), rt.Elem())
				} else {
					rtu = reflect.SliceOf(rt.Elem())
				}
				rtuid := reflect.ValueOf(rtu).Pointer()
				if idx := fastpathAV.index(rtuid); idx != -1 {
					xfnf := fastpathAV[idx].encfn
					xrt := fastpathAV[idx].rt
					fn.f = func(xf *encFnInfo, xrv reflect.Value) {
						xfnf(xf, xrv.Convert(xrt))
					}
				}
			}
		}
		if fn.f == nil {
			switch rk {
			case reflect.Bool:
				fn.f = (*encFnInfo).kBool
			case reflect.String:
				fn.f = (*encFnInfo).kString
			case reflect.Float64:
				fn.f = (*encFnInfo).kFloat64
			case reflect.Float32:
				fn.f = (*encFnInfo).kFloat32
			case reflect.Int, reflect.Int8, reflect.Int64, reflect.Int32, reflect.Int16:
				fn.f = (*encFnInfo).kInt
			case reflect.Uint8, reflect.Uint64, reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uintptr:
				fn.f = (*encFnInfo).kUint
			case reflect.Invalid:
				fn.f = (*encFnInfo).kInvalid
			case reflect.Chan:
				fi.seq = seqTypeChan
				fn.f = (*encFnInfo).kSlice
			case reflect.Slice:
				fi.seq = seqTypeSlice
				fn.f = (*encFnInfo).kSlice
			case reflect.Array:
				fi.seq = seqTypeArray
				fn.f = (*encFnInfo).kSlice
			case reflect.Struct:
				fn.f = (*encFnInfo).kStruct
				// reflect.Ptr and reflect.Interface are handled already by preEncodeValue
				// case reflect.Ptr:
				// 	fn.f = (*encFnInfo).kPtr
				// case reflect.Interface:
				// 	fn.f = (*encFnInfo).kInterface
			case reflect.Map:
				fn.f = (*encFnInfo).kMap
			default:
				fn.f = (*encFnInfo).kErr
			}
		}
	}

	return
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
}

func (e *Encoder) marshal(bs []byte, fnerr error, asis bool, c charEncoding) {
	if fnerr != nil {
		panic(fnerr)
	}
	if bs == nil {
		e.e.EncodeNil()
	} else if asis {
		e.asis(bs)
	} else {
		e.e.EncodeStringBytes(c, bs)
	}
}

func (e *Encoder) asis(v []byte) {
<<<<<<< HEAD
	if e.isas {
		e.as.EncodeAsis(v)
	} else {
		e.w.writeb(v)
	}
}

func (e *Encoder) rawBytes(vv Raw) {
=======
	if e.as == nil {
		e.w.writeb(v)
	} else {
		e.as.EncodeAsis(v)
	}
}

func (e *Encoder) raw(vv Raw) {
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
	v := []byte(vv)
	if !e.h.Raw {
		e.errorf("Raw values cannot be encoded: %v", v)
	}
<<<<<<< HEAD
	e.asis(v)
}

func (e *Encoder) wrapErrstr(v interface{}, err *error) {
	*err = fmt.Errorf("%s encode error: %v", e.hh.Name(), v)
}
=======
	if e.as == nil {
		e.w.writeb(v)
	} else {
		e.as.EncodeAsis(v)
	}
}

func (e *Encoder) errorf(format string, params ...interface{}) {
	err := fmt.Errorf(format, params...)
	panic(err)
}

// ----------------------------------------

const encStructPoolLen = 5

// encStructPool is an array of sync.Pool.
// Each element of the array pools one of encStructPool(8|16|32|64).
// It allows the re-use of slices up to 64 in length.
// A performance cost of encoding structs was collecting
// which values were empty and should be omitted.
// We needed slices of reflect.Value and string to collect them.
// This shared pool reduces the amount of unnecessary creation we do.
// The cost is that of locking sometimes, but sync.Pool is efficient
// enough to reduce thread contention.
var encStructPool [encStructPoolLen]sync.Pool

func init() {
	encStructPool[0].New = func() interface{} { return new([8]stringRv) }
	encStructPool[1].New = func() interface{} { return new([16]stringRv) }
	encStructPool[2].New = func() interface{} { return new([32]stringRv) }
	encStructPool[3].New = func() interface{} { return new([64]stringRv) }
	encStructPool[4].New = func() interface{} { return new([128]stringRv) }
}

func encStructPoolGet(newlen int) (p *sync.Pool, v interface{}, s []stringRv) {
	// if encStructPoolLen != 5 { // constant chec, so removed at build time.
	// 	panic(errors.New("encStructPoolLen must be equal to 4")) // defensive, in case it is changed
	// }
	// idxpool := newlen / 8
	if newlen <= 8 {
		p = &encStructPool[0]
		v = p.Get()
		s = v.(*[8]stringRv)[:newlen]
	} else if newlen <= 16 {
		p = &encStructPool[1]
		v = p.Get()
		s = v.(*[16]stringRv)[:newlen]
	} else if newlen <= 32 {
		p = &encStructPool[2]
		v = p.Get()
		s = v.(*[32]stringRv)[:newlen]
	} else if newlen <= 64 {
		p = &encStructPool[3]
		v = p.Get()
		s = v.(*[64]stringRv)[:newlen]
	} else if newlen <= 128 {
		p = &encStructPool[4]
		v = p.Get()
		s = v.(*[128]stringRv)[:newlen]
	} else {
		s = make([]stringRv, newlen)
	}
	return
}

// ----------------------------------------

// func encErr(format string, params ...interface{}) {
// 	doPanic(msgTagEnc, format, params...)
// }
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
