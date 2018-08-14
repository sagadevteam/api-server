<<<<<<< HEAD
// Copyright 2012 The Go Authors. All rights reserved.
=======
// Copyright 2012 The Go Authors.  All rights reserved.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build plan9,race

package plan9

import (
	"runtime"
	"unsafe"
)

const raceenabled = true

func raceAcquire(addr unsafe.Pointer) {
	runtime.RaceAcquire(addr)
}

func raceReleaseMerge(addr unsafe.Pointer) {
	runtime.RaceReleaseMerge(addr)
}

func raceReadRange(addr unsafe.Pointer, len int) {
	runtime.RaceReadRange(addr, len)
}

func raceWriteRange(addr unsafe.Pointer, len int) {
	runtime.RaceWriteRange(addr, len)
}
