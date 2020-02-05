// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for AMD64, OpenBSD
// /usr/src/sys/kern/syscalls.master for syscall numbers.
//

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

#define CLOCK_MONOTONIC	$3

TEXT runtime·newthread(SB),NOSPLIT,$0
    MOVQ    mp+0(FP), R8    // arg 1 - m
    MOVQ    m_g0(R8), R9    // g

    // hook up g & m
    get_tls(CX)
    MOVQ    R9, g(CX)
    MOVQ    R8, g_m(R9)

    MOVQ    SP, AX
    MOVQ    AX, (g_stack+stack_hi)(R9)
    SUBQ    $(64*1024), AX
    MOVQ    AX, (g_stack+stack_lo)(R9)
    ADDQ    $const__StackGuard, AX
    MOVQ    AX, g_stackguard0(R9)
    MOVQ    AX, g_stackguard1(R9)

    CLD

    CALL    runtime·stackcheck(SB)
    CALL    runtime·mstart(SB)

    XORL    AX, AX
    MOVL    AX, ret+8(FP)
    RET

TEXT runtime·setitimer(SB),NOSPLIT,$-8
	MOVL	mode+0(FP), DI		// arg 1 - which
	MOVQ	new+8(FP), SI		// arg 2 - itv
	MOVQ	old+16(FP), DX		// arg 3 - oitv
	MOVL	$69, AX			// sys_setitimer
	SYSCALL
	RET

// func walltime1() (sec int64, nsec int32)
TEXT runtime·walltime1(SB), NOSPLIT, $32
	MOVQ	$0, DI			// arg 1 - clock_id
	LEAQ	8(SP), SI		// arg 2 - tp
	MOVL	$87, AX			// sys_clock_gettime
	SYSCALL
	MOVQ	8(SP), AX		// sec
	MOVQ	16(SP), DX		// nsec

	// sec is in AX, nsec in DX
	MOVQ	AX, sec+0(FP)
	MOVL	DX, nsec+8(FP)
	RET

TEXT runtime·nanotime1(SB),NOSPLIT,$24
	MOVQ	CLOCK_MONOTONIC, DI	// arg 1 - clock_id
	LEAQ	8(SP), SI		// arg 2 - tp
	MOVL	$87, AX			// sys_clock_gettime
	SYSCALL
	MOVQ	8(SP), AX		// sec
	MOVQ	16(SP), DX		// nsec

	// sec is in AX, nsec in DX
	// return nsec in AX
	IMULQ	$1000000000, AX
	ADDQ	DX, AX
	MOVQ	AX, ret+0(FP)
	RET

TEXT runtime·sigaction(SB),NOSPLIT,$-8
	MOVL	sig+0(FP), DI		// arg 1 - signum
	MOVQ	new+8(FP), SI		// arg 2 - nsa
	MOVQ	old+16(FP), DX		// arg 3 - osa
	MOVL	$46, AX
	SYSCALL
	JCC	2(PC)
	MOVL	$0xf1, 0xf1		// crash
	RET

TEXT runtime·obsdsigprocmask(SB),NOSPLIT,$0
	MOVL	how+0(FP), DI		// arg 1 - how
	MOVL	new+4(FP), SI		// arg 2 - set
	MOVL	$48, AX			// sys_sigprocmask
	SYSCALL
	JCC	2(PC)
	MOVL	$0xf1, 0xf1		// crash
	MOVL	AX, ret+8(FP)
	RET

TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVQ	fn+0(FP),    AX
	MOVL	sig+8(FP),   DI
	MOVQ	info+16(FP), SI
	MOVQ	ctx+24(FP),  DX
	PUSHQ	BP
	MOVQ	SP, BP
	ANDQ	$~15, SP     // alignment for x86_64 ABI
	CALL	AX
	MOVQ	BP, SP
	POPQ	BP
	RET

TEXT runtime·sigtramp(SB),NOSPLIT,$72
	// Save callee-saved C registers, since the caller may be a C signal handler.
	MOVQ	BX,  bx-8(SP)
	MOVQ	BP,  bp-16(SP)  // save in case GOEXPERIMENT=noframepointer is set
	MOVQ	R12, r12-24(SP)
	MOVQ	R13, r13-32(SP)
	MOVQ	R14, r14-40(SP)
	MOVQ	R15, r15-48(SP)
	// We don't save mxcsr or the x87 control word because sigtrampgo doesn't
	// modify them.

	MOVQ	DX, ctx-56(SP)
	MOVQ	SI, info-64(SP)
	MOVQ	DI, signum-72(SP)
	CALL	runtime·sigtrampgo(SB)

	MOVQ	r15-48(SP), R15
	MOVQ	r14-40(SP), R14
	MOVQ	r13-32(SP), R13
	MOVQ	r12-24(SP), R12
	MOVQ	bp-16(SP),  BP
	MOVQ	bx-8(SP),   BX
	RET

TEXT runtime·mmap(SB),NOSPLIT,$0
	MOVQ	addr+0(FP), DI		// arg 1 - addr
	MOVQ	n+8(FP), SI		// arg 2 - len
	MOVL	prot+16(FP), DX		// arg 3 - prot
	MOVL	flags+20(FP), R10		// arg 4 - flags
	MOVL	fd+24(FP), R8		// arg 5 - fd
	MOVL	off+28(FP), R9
	SUBQ	$16, SP
	MOVQ	R9, 8(SP)		// arg 7 - offset (passed on stack)
	MOVQ	$0, R9			// arg 6 - pad
	MOVL	$197, AX
	SYSCALL
	JCC	ok
	ADDQ	$16, SP
	MOVQ	$0, p+32(FP)
	MOVQ	AX, err+40(FP)
	RET
ok:
	ADDQ	$16, SP
	MOVQ	AX, p+32(FP)
	MOVQ	$0, err+40(FP)
	RET

TEXT runtime·munmap(SB),NOSPLIT,$0
	MOVQ	addr+0(FP), DI		// arg 1 - addr
	MOVQ	n+8(FP), SI		// arg 2 - len
	MOVL	$73, AX			// sys_munmap
	SYSCALL
	JCC	2(PC)
	MOVL	$0xf1, 0xf1		// crash
	RET

TEXT runtime·sigaltstack(SB),NOSPLIT,$-8
	MOVQ	new+0(FP), DI		// arg 1 - nss
	MOVQ	old+8(FP), SI		// arg 2 - oss
	MOVQ	$288, AX		// sys_sigaltstack
	SYSCALL
	JCC	2(PC)
	MOVL	$0xf1, 0xf1		// crash
	RET

// needed for asm_amd64.s
TEXT runtime·settls(SB),NOSPLIT,$0
	RET

TEXT runtime·sysctl(SB),NOSPLIT,$0
	MOVQ	mib+0(FP), DI		// arg 1 - name
	MOVL	miblen+8(FP), SI		// arg 2 - namelen
	MOVQ	out+16(FP), DX		// arg 3 - oldp
	MOVQ	size+24(FP), R10		// arg 4 - oldlenp
	MOVQ	dst+32(FP), R8		// arg 5 - newp
	MOVQ	ndst+40(FP), R9		// arg 6 - newlen
	MOVQ	$202, AX		// sys___sysctl
	SYSCALL
	JCC	4(PC)
	NEGQ	AX
	MOVL	AX, ret+48(FP)
	RET
	MOVL	$0, AX
	MOVL	AX, ret+48(FP)
	RET

// int32 runtime·kqueue(void);
TEXT runtime·kqueue(SB),NOSPLIT,$0
	MOVL	$269, AX
	SYSCALL
	JCC	2(PC)
	NEGQ	AX
	MOVL	AX, ret+0(FP)
	RET

// int32 runtime·kevent(int kq, Kevent *changelist, int nchanges, Kevent *eventlist, int nevents, Timespec *timeout);
TEXT runtime·kevent(SB),NOSPLIT,$0
	MOVL	kq+0(FP), DI
	MOVQ	ch+8(FP), SI
	MOVL	nch+16(FP), DX
	MOVQ	ev+24(FP), R10
	MOVL	nev+32(FP), R8
	MOVQ	ts+40(FP), R9
	MOVL	$72, AX
	SYSCALL
	JCC	2(PC)
	NEGQ	AX
	MOVL	AX, ret+48(FP)
	RET

// void runtime·closeonexec(int32 fd);
TEXT runtime·closeonexec(SB),NOSPLIT,$0
	MOVL	fd+0(FP), DI	// fd
	MOVQ	$2, SI		// F_SETFD
	MOVQ	$1, DX		// FD_CLOEXEC
	MOVL	$92, AX		// fcntl
	SYSCALL
	RET

// func runtime·setNonblock(int32 fd)
TEXT runtime·setNonblock(SB),NOSPLIT,$0-4
	MOVL    fd+0(FP), DI  // fd
	MOVQ    $3, SI  // F_GETFL
	MOVQ    $0, DX
	MOVL	$92, AX // fcntl
	SYSCALL
	MOVL	fd+0(FP), DI // fd
	MOVQ	$4, SI // F_SETFL
	MOVQ	$4, DX // O_NONBLOCK
	ORL	AX, DX
	MOVL	$92, AX // fcntl
	SYSCALL
	RET

// Called by runtime·asmcgocall or runtime·cgocall.
// NOT USING GO CALLING CONVENTION.
TEXT runtime·asmsysobsd6(SB),NOSPLIT,$0
  PUSHQ DI
  MOVQ  libcall_fn(DI), AX
  MOVQ  libcall_args(DI), R11
  MOVQ  libcall_args(DI), R10

  get_tls(CX)
  MOVQ  g(CX), BX
  CMPQ  BX, $0
  JEQ skiperrno1
  MOVQ  g_m(BX), BX
  MOVQ  (m_mOS+mOS_perrno)(BX), DX
  CMPQ  DX, $0
  JEQ skiperrno1
  MOVL  $0, 0(DX)

skiperrno1:
  CMPQ  R11, $0
  JEQ skipargs
  // Load 6 args into corresponding registers
  MOVQ  0(R11), DI
  MOVQ  8(R11), SI
  MOVQ  16(R11), DX
  MOVQ  24(R11), CX
  MOVQ  32(R11), R8
  MOVQ  40(R11), R9

skipargs:
  // Call obsd func
  CALL AX

  // Return result
  POPQ  DI
  MOVQ  AX, libcall_r1(DI)
  MOVQ  DX, libcall_r2(DI)

  get_tls(CX)
  // get goroutine context
  MOVQ  g(CX), BX
  CMPQ  BX, $0
  JEQ skiperrno2
  MOVQ  g_m(BX), BX
  // get errno offset and store addr to errno in AX
  MOVQ  (m_mOS+mOS_perrno)(BX), AX
  // if addr is 0, skip loading errno
  CMPQ  AX, $0
  JEQ skiperrno2
  MOVL  0(AX), AX
  MOVQ  AX, libcall_err(DI)

skiperrno2:
  RET
