<<<<<<< HEAD
// Copyright 2012 The Go Authors. All rights reserved.
=======
// Copyright 2012 The Go Authors.  All rights reserved.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows,!race

package windows

import (
	"unsafe"
)

const raceenabled = false

func raceAcquire(addr unsafe.Pointer) {
}

func raceReleaseMerge(addr unsafe.Pointer) {
}

func raceReadRange(addr unsafe.Pointer, len int) {
}

func raceWriteRange(addr unsafe.Pointer, len int) {
}
