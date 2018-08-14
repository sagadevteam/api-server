// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build s390x
// +build linux
// +build !gccgo

#include "textflag.h"

//
// System calls for s390x, Linux
//

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT ·Syscall(SB),NOSPLIT,$0-56
	BR	syscall·Syscall(SB)

TEXT ·Syscall6(SB),NOSPLIT,$0-80
	BR	syscall·Syscall6(SB)

<<<<<<< HEAD
TEXT ·SyscallNoError(SB),NOSPLIT,$0-48
	BL	runtime·entersyscall(SB)
	MOVD	a1+8(FP), R2
	MOVD	a2+16(FP), R3
	MOVD	a3+24(FP), R4
	MOVD	$0, R5
	MOVD	$0, R6
	MOVD	$0, R7
	MOVD	trap+0(FP), R1	// syscall entry
	SYSCALL
	MOVD	R2, r1+32(FP)
	MOVD	R3, r2+40(FP)
	BL	runtime·exitsyscall(SB)
	RET

=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
TEXT ·RawSyscall(SB),NOSPLIT,$0-56
	BR	syscall·RawSyscall(SB)

TEXT ·RawSyscall6(SB),NOSPLIT,$0-80
	BR	syscall·RawSyscall6(SB)
<<<<<<< HEAD

TEXT ·RawSyscallNoError(SB),NOSPLIT,$0-48
	MOVD	a1+8(FP), R2
	MOVD	a2+16(FP), R3
	MOVD	a3+24(FP), R4
	MOVD	$0, R5
	MOVD	$0, R6
	MOVD	$0, R7
	MOVD	trap+0(FP), R1	// syscall entry
	SYSCALL
	MOVD	R2, r1+32(FP)
	MOVD	R3, r2+40(FP)
	RET
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
