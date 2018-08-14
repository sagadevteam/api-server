<<<<<<< HEAD
// Copyright (c) 2012-2018 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

/*
Package codec provides a
High Performance, Feature-Rich Idiomatic Go 1.4+ codec/encoding library
for binc, msgpack, cbor, json.
=======
// Copyright (c) 2012-2015 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

/*
High Performance, Feature-Rich Idiomatic Go codec/encoding library for 
binc, msgpack, cbor, json.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

Supported Serialization formats are:

  - msgpack: https://github.com/msgpack/msgpack
  - binc:    http://github.com/ugorji/binc
  - cbor:    http://cbor.io http://tools.ietf.org/html/rfc7049
  - json:    http://json.org http://tools.ietf.org/html/rfc7159
<<<<<<< HEAD
  - simple:
=======
  - simple: 
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

To install:

    go get github.com/ugorji/go/codec

<<<<<<< HEAD
This package will carefully use 'unsafe' for performance reasons in specific places.
You can build without unsafe use by passing the safe or appengine tag
i.e. 'go install -tags=safe ...'. Note that unsafe is only supported for the last 3
go sdk versions e.g. current go release is go 1.9, so we support unsafe use only from
go 1.7+ . This is because supporting unsafe requires knowledge of implementation details.
=======
This package understands the 'unsafe' tag, to allow using unsafe semantics:

  - When decoding into a struct, you need to read the field name as a string 
    so you can find the struct field it is mapped to.
    Using `unsafe` will bypass the allocation and copying overhead of []byte->string conversion.

To install using unsafe, pass the 'unsafe' tag:

    go get -tags=unsafe github.com/ugorji/go/codec
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

For detailed usage information, read the primer at http://ugorji.net/blog/go-codec-primer .

The idiomatic Go support is as seen in other encoding packages in
the standard library (ie json, xml, gob, etc).

Rich Feature Set includes:

  - Simple but extremely powerful and feature-rich API
<<<<<<< HEAD
  - Support for go1.4 and above, while selectively using newer APIs for later releases
  - Excellent code coverage ( > 90% )
  - Very High Performance.
    Our extensive benchmarks show us outperforming Gob, Json, Bson, etc by 2-4X.
  - Careful selected use of 'unsafe' for targeted performance gains.
    100% mode exists where 'unsafe' is not used at all.
  - Lock-free (sans mutex) concurrency for scaling to 100's of cores
  - Coerce types where appropriate
    e.g. decode an int in the stream into a float, decode numbers from formatted strings, etc
  - Corner Cases:
=======
  - Very High Performance.
    Our extensive benchmarks show us outperforming Gob, Json, Bson, etc by 2-4X.
  - Multiple conversions:
    Package coerces types where appropriate 
    e.g. decode an int in the stream into a float, etc.
  - Corner Cases: 
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
    Overflows, nil maps/slices, nil values in streams are handled correctly
  - Standard field renaming via tags
  - Support for omitting empty fields during an encoding
  - Encoding from any value and decoding into pointer to any value
    (struct, slice, map, primitives, pointers, interface{}, etc)
  - Extensions to support efficient encoding/decoding of any named types
  - Support encoding.(Binary|Text)(M|Unm)arshaler interfaces
<<<<<<< HEAD
  - Support IsZero() bool to determine if a value is a zero value.
    Analogous to time.Time.IsZero() bool.
  - Decoding without a schema (into a interface{}).
    Includes Options to configure what specific map or slice type to use
    when decoding an encoded list or map into a nil interface{}
  - Mapping a non-interface type to an interface, so we can decode appropriately
    into any interface type with a correctly configured non-interface value.
  - Encode a struct as an array, and decode struct from an array in the data stream
  - Option to encode struct keys as numbers (instead of strings)
    (to support structured streams with fields encoded as numeric codes)
=======
  - Decoding without a schema (into a interface{}).
    Includes Options to configure what specific map or slice type to use
    when decoding an encoded list or map into a nil interface{}
  - Encode a struct as an array, and decode struct from an array in the data stream
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
  - Comprehensive support for anonymous fields
  - Fast (no-reflection) encoding/decoding of common maps and slices
  - Code-generation for faster performance.
  - Support binary (e.g. messagepack, cbor) and text (e.g. json) formats
<<<<<<< HEAD
  - Support indefinite-length formats to enable true streaming
=======
  - Support indefinite-length formats to enable true streaming 
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
    (for formats which support it e.g. json, cbor)
  - Support canonical encoding, where a value is ALWAYS encoded as same sequence of bytes.
    This mostly applies to maps, where iteration order is non-deterministic.
  - NIL in data stream decoded as zero value
  - Never silently skip data when decoding.
    User decides whether to return an error or silently skip data when keys or indexes
    in the data stream do not map to fields in the struct.
  - Detect and error when encoding a cyclic reference (instead of stack overflow shutdown)
  - Encode/Decode from/to chan types (for iterative streaming support)
  - Drop-in replacement for encoding/json. `json:` key in struct tag supported.
  - Provides a RPC Server and Client Codec for net/rpc communication protocol.
<<<<<<< HEAD
  - Handle unique idiosyncrasies of codecs e.g.
    - For messagepack, configure how ambiguities in handling raw bytes are resolved
    - For messagepack, provide rpc server/client codec to support
      msgpack-rpc protocol defined at:
      https://github.com/msgpack-rpc/msgpack-rpc/blob/master/spec.md

=======
  - Handle unique idiosyncrasies of codecs e.g. 
    - For messagepack, configure how ambiguities in handling raw bytes are resolved 
    - For messagepack, provide rpc server/client codec to support 
      msgpack-rpc protocol defined at:
      https://github.com/msgpack-rpc/msgpack-rpc/blob/master/spec.md
  
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
Extension Support

Users can register a function to handle the encoding or decoding of
their custom types.

There are no restrictions on what the custom type can be. Some examples:

    type BisSet   []int
    type BitSet64 uint64
    type UUID     string
    type MyStructWithUnexportedFields struct { a int; b bool; c []int; }
    type GifImage struct { ... }

As an illustration, MyStructWithUnexportedFields would normally be
encoded as an empty map because it has no exported fields, while UUID
would be encoded as a string. However, with extension support, you can
encode any of these however you like.

<<<<<<< HEAD
Custom Encoding and Decoding

This package maintains symmetry in the encoding and decoding halfs.
We determine how to encode or decode by walking this decision tree

  - is type a codec.Selfer?
  - is there an extension registered for the type?
  - is format binary, and is type a encoding.BinaryMarshaler and BinaryUnmarshaler?
  - is format specifically json, and is type a encoding/json.Marshaler and Unmarshaler?
  - is format text-based, and type an encoding.TextMarshaler?
  - else we use a pair of functions based on the "kind" of the type e.g. map, slice, int64, etc

This symmetry is important to reduce chances of issues happening because the
encoding and decoding sides are out of sync e.g. decoded via very specific
encoding.TextUnmarshaler but encoded via kind-specific generalized mode.

Consequently, if a type only defines one-half of the symmetry
(e.g. it implements UnmarshalJSON() but not MarshalJSON() ),
then that type doesn't satisfy the check and we will continue walking down the
decision tree.

=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
RPC

RPC Client and Server Codecs are implemented, so the codecs can be used
with the standard net/rpc package.

Usage

The Handle is SAFE for concurrent READ, but NOT SAFE for concurrent modification.

The Encoder and Decoder are NOT safe for concurrent use.

Consequently, the usage model is basically:

    - Create and initialize the Handle before any use.
      Once created, DO NOT modify it.
    - Multiple Encoders or Decoders can now use the Handle concurrently.
      They only read information off the Handle (never write).
    - However, each Encoder or Decoder MUST not be used concurrently
    - To re-use an Encoder/Decoder, call Reset(...) on it first.
      This allows you use state maintained on the Encoder/Decoder.

Sample usage model:

    // create and configure Handle
    var (
      bh codec.BincHandle
      mh codec.MsgpackHandle
      ch codec.CborHandle
    )

    mh.MapType = reflect.TypeOf(map[string]interface{}(nil))

    // configure extensions
    // e.g. for msgpack, define functions and enable Time support for tag 1
    // mh.SetExt(reflect.TypeOf(time.Time{}), 1, myExt)

    // create and use decoder/encoder
    var (
      r io.Reader
      w io.Writer
      b []byte
      h = &bh // or mh to use msgpack
    )

    dec = codec.NewDecoder(r, h)
    dec = codec.NewDecoderBytes(b, h)
    err = dec.Decode(&v)

    enc = codec.NewEncoder(w, h)
    enc = codec.NewEncoderBytes(&b, h)
    err = enc.Encode(v)

    //RPC Server
    go func() {
        for {
            conn, err := listener.Accept()
            rpcCodec := codec.GoRpc.ServerCodec(conn, h)
            //OR rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, h)
            rpc.ServeCodec(rpcCodec)
        }
    }()

    //RPC Communication (client side)
    conn, err = net.Dial("tcp", "localhost:5555")
    rpcCodec := codec.GoRpc.ClientCodec(conn, h)
    //OR rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, h)
    client := rpc.NewClientWithCodec(rpcCodec)

<<<<<<< HEAD
Running Tests

To run tests, use the following:

    go test

To run the full suite of tests, use the following:

    go test -tags alltests -run Suite

You can run the tag 'safe' to run tests or build in safe mode. e.g.

    go test -tags safe -run Json
    go test -tags "alltests safe" -run Suite

Running Benchmarks

Please see http://github.com/ugorji/go-codec-bench .

Caveats

Struct fields matching the following are ignored during encoding and decoding
    - struct tag value set to -
    - func, complex numbers, unsafe pointers
    - unexported and not embedded
    - unexported and embedded and not struct kind
    - unexported and embedded pointers (from go1.10)

Every other field in a struct will be encoded/decoded.

Embedded fields are encoded as if they exist in the top-level struct,
with some caveats. See Encode documentation.

*/
package codec

// TODO:
//   - For Go 1.11, when mid-stack inlining is enabled,
//     we should use committed functions for writeXXX and readXXX calls.
//     This involves uncommenting the methods for decReaderSwitch and encWriterSwitch
//     and using those (decReaderSwitch and encWriterSwitch) in all handles
//     instead of encWriter and decReader.
//     The benefit is that, for the (En|De)coder over []byte, the encWriter/decReader
//     will be inlined, giving a performance bump for that typical case.
//     However, it will only  be inlined if mid-stack inlining is enabled,
//     as we call panic to raise errors, and panic currently prevents inlining.
//
// PUNTED:
//   - To make Handle comparable, make extHandle in BasicHandle a non-embedded pointer,
//     and use overlay methods on *BasicHandle to call through to extHandle after initializing
//     the "xh *extHandle" to point to a real slice.
//
// BEFORE EACH RELEASE:
//   - Look through and fix padding for each type, to eliminate false sharing
//     - critical shared objects that are read many times
//       TypeInfos
//     - pooled objects:
//       decNaked, decNakedContainers, codecFner, typeInfoLoadArray, 
//     - small objects allocated independently, that we read/use much across threads:
//       codecFn, typeInfo
//     - Objects allocated independently and used a lot
//       Decoder, Encoder,
//       xxxHandle, xxxEncDriver, xxxDecDriver (xxx = json, msgpack, cbor, binc, simple)
//     - In all above, arrange values modified together to be close to each other.
//
//     For all of these, either ensure that they occupy full cache lines,
//     or ensure that the things just past the cache line boundary are hardly read/written
//     e.g. JsonHandle.RawBytesExt - which is copied into json(En|De)cDriver at init
//
//     Occupying full cache lines means they occupy 8*N words (where N is an integer).
//     Check this out by running: ./run.sh -z
//     - look at those tagged ****, meaning they are not occupying full cache lines
//     - look at those tagged <<<<, meaning they are larger than 32 words (something to watch)
//   - Run "golint -min_confidence 0.81"
=======
*/
package codec

// Benefits of go-codec:
//
//    - encoding/json always reads whole file into memory first.
//      This makes it unsuitable for parsing very large files.
//    - encoding/xml cannot parse into a map[string]interface{}
//      I found this out on reading https://github.com/clbanning/mxj

// TODO:
//
//   - optimization for codecgen:
//     if len of entity is <= 3 words, then support a value receiver for encode.
//   - (En|De)coder should store an error when it occurs.
//     Until reset, subsequent calls return that error that was stored.
//     This means that free panics must go away.
//     All errors must be raised through errorf method.
//   - Decoding using a chan is good, but incurs concurrency costs.
//     This is because there's no fast way to use a channel without it
//     having to switch goroutines constantly.
//     Callback pattern is still the best. Maybe consider supporting something like:
//        type X struct {
//             Name string
//             Ys []Y
//             Ys chan <- Y
//             Ys func(Y) -> call this function for each entry
//        }
//    - Consider adding a isZeroer interface { isZero() bool }
//      It is used within isEmpty, for omitEmpty support.
//    - Consider making Handle used AS-IS within the encoding/decoding session.
//      This means that we don't cache Handle information within the (En|De)coder,
//      except we really need it at Reset(...)
//    - Consider adding math/big support
//    - Consider reducing the size of the generated functions:
//      Maybe use one loop, and put the conditionals in the loop.
//      for ... { if cLen > 0 { if j == cLen { break } } else if dd.CheckBreak() { break } }
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
