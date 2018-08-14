<<<<<<< HEAD
// Copyright 2010 The Go Authors. All rights reserved.
=======
// Copyright 2010 The Go Authors.  All rights reserved.
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows environment variables.

package windows

import "syscall"

func Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}

func Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}

func Clearenv() {
	syscall.Clearenv()
}

func Environ() []string {
	return syscall.Environ()
}
<<<<<<< HEAD

func Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
