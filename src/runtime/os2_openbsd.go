// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

const (
	_SS_DISABLE  = 4
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 3
	_NSIG        = 33
	_SI_USER     = 0
)

//go:cgo_import_dynamic __errno __errno "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_close close "libc.so"
//go:cgo_import_dynamic libc_exit _exit "libc.so"
//go:cgo_import_dynamic libc_nanosleep nanosleep "libc.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libc_open open "libc.so"
//go:cgo_import_dynamic libc_madvise madvise "libc.so"
//go:cgo_import_dynamic libc___threxit __threxit "libc.so"
//go:cgo_import_dynamic libc_sched_yield sched_yield "libc.so"
//go:cgo_import_dynamic libc___thrsleep __thrsleep "libc.so"
//go:cgo_import_dynamic libc___thrwakeup __thrwakeup "libc.so"
//go:cgo_import_dynamic libc_pipe pipe "libc.so"
//go:cgo_import_dynamic libc_pipe2 pipe2 "libc.so"
//go:cgo_import_dynamic libc_getthrid getthrid "libc.so"
//go:cgo_import_dynamic libc_thrkill thrkill "libc.so"
//go:cgo_import_dynamic libc_getpid getpid "libc.so"
//go:cgo_import_dynamic libc_kill kill "libc.so"

//go:cgo_import_dynamic libpthread_attr_init pthread_attr_init "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_destroy pthread_attr_destroy "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_setstack pthread_attr_setstack "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_getstack pthread_attr_getstack "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_setstacksize pthread_attr_setstacksize "libpthread.so:
//go:cgo_import_dynamic libpthread_attr_getstacksize pthread_attr_getstacksize "libpthread.so:
//go:cgo_import_dynamic libpthread_attr_setdetachstate pthread_attr_setdetachstate "libpthread.so"
//go:cgo_import_dynamic libpthread_create pthread_create "libpthread.so"

//go:linkname __errno __errno
//go:linkname libc_read libc_read
//go:linkname libc_close libc_close
//go:linkname libc_exit libc_exit
//go:linkname libc_nanosleep libc_nanosleep
//go:linkname libc_write libc_write
//go:linkname libc_open libc_open
//go:linkname libc_madvise libc_madvise
//go:linkname libc___threxit libc___threxit
//go:linkname libc_sched_yield libc_sched_yield
//go:linkname libc___thrsleep libc___thrsleep
//go:linkname libc___thrwakeup libc___thrwakeup
//go:linkname libc_pipe libc_pipe
//go:linkname libc_pipe2 libc_pipe2
//go:linkname libc_getthrid libc_getthrid
//go:linkname libc_thrkill libc_thrkill
//go:linkname libc_getpid libc_getpid
//go:linkname libc_kill libc_kill

//go:linkname libpthread_attr_init libpthread_attr_init
//go:linkname libpthread_attr_destroy libpthread_attr_destroy
//go:linkname libpthread_attr_setstack libpthread_attr_setstack
//go:linkname libpthread_attr_getstack libpthread_attr_getstack
//go:linkname libpthread_attr_setstacksize libpthread_attr_setstacksize
//go:linkname libpthread_attr_getstacksize libpthread_attr_getstacksize
//go:linkname libpthread_attr_setdetachstate libpthread_attr_setdetachstate
//go:linkname libpthread_create libpthread_create

var (
	__errno          libcFunc
	libc_read        libcFunc
	libc_close       libcFunc
	libc_exit        libcFunc
	libc_nanosleep   libcFunc
	libc_write       libcFunc
	libc_open        libcFunc
	libc_madvise     libcFunc
	libc___threxit   libcFunc
	libc_sched_yield libcFunc
	libc___thrsleep  libcFunc
	libc___thrwakeup libcFunc
	libc_pipe        libcFunc
	libc_pipe2       libcFunc
	libc_getthrid    libcFunc
	libc_thrkill     libcFunc
	libc_getpid      libcFunc
	libc_kill        libcFunc

	libpthread_attr_init           libcFunc
	libpthread_attr_destroy        libcFunc
	libpthread_attr_setstack       libcFunc
	libpthread_attr_getstack       libcFunc
	libpthread_attr_setstacksize   libcFunc
	libpthread_attr_getstacksize   libcFunc
	libpthread_attr_setdetachstate libcFunc
	libpthread_create              libcFunc
)

//go:nosplit
func read(d int32, buf unsafe.Pointer, nbytes int32) int32 {
	r, errno := syscall3(&libc_read, uintptr(d), uintptr(buf), uintptr(nbytes))

	if int32(r) < 0 {
		return -int32(errno)
	}
	return int32(r)
}

//go:nosplit
func closefd(d int32) int32 {
	r, errno := syscall1(&libc_close, uintptr(d))

	if int32(r) == -1 {
		return int32(errno)
	}
	return int32(3)
}

//go:nosplit
func exit(status int32) {
	syscall1(&libc_exit, uintptr(status))
}

//go:nosplit
func usleep(usec uint32) {
	ts := &timespec{
		tv_sec:  int64(usec / 1000000),
		tv_nsec: int64((usec % 1000000) * 1000),
	}

	syscall2(&libc_nanosleep, uintptr(unsafe.Pointer(ts)), 0)
}

//go:nosplit
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	r, errno := syscall3(&libc_write, fd, uintptr(p), uintptr(n))

	if n == -1 {
		return -int32(errno)
	}
	return int32(r)
}

//go:nosplit
func open(name *byte, mode, perm int32) int32 {
	r, _ := syscall3(&libc_open, uintptr(unsafe.Pointer(name)), uintptr(mode), uintptr(perm))

	return int32(r)
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) int32 {
	r, _ := syscall3(&libc_madvise, uintptr(addr), n, uintptr(flags))
	return int32(r)
}

func exitThread(wait *uint32) {
	syscall1(&libc___threxit, uintptr(unsafe.Pointer(wait)))
}

func osyield() int32 {
	r, errno := syscall0(&libc_sched_yield)
	if int32(r) == -1 {
		return -int32(errno)
	}

	return int32(r)
}

func thrsleep(id unsafe.Pointer, clock_id int32, abstime *timespec, lock unsafe.Pointer, abort *int32) int32 {
	r, _ := syscall5(&libc___thrsleep, uintptr(id), uintptr(clock_id), uintptr(unsafe.Pointer(abstime)), uintptr(lock), uintptr(unsafe.Pointer(abort)))

	return int32(r)
}

func thrwakeup(id unsafe.Pointer, n int32) int32 {
	r, _ := syscall2(&libc___thrwakeup, uintptr(id), uintptr(n))
	return int32(r)
}

func pipe() (r, w int32, errno int32) {
	arr := []int32{r, w}
	ret, rerrno := syscall1(&libc_pipe, uintptr(unsafe.Pointer(&arr)))
	if int32(ret) == -1 {
		errno = int32(rerrno)
		return
	}

	return r, w, 0
}

func pipe2(flags int32) (r, w int32, errno int32) {
	arr := []int32{r, w}
	ret, rerrno := syscall1(&libc_pipe2, uintptr(unsafe.Pointer(&arr)))
	if int32(ret) == -1 {
		errno = int32(rerrno)
		return
	}

	return r, w, 0
}

func getthrid() int32 {
	r, _ := syscall0(&libc_getthrid)
	return int32(r)
}

func thrkill(tid int32, sig int) {
	syscall3(&libc_thrkill, uintptr(tid), uintptr(sig), 0)
}

func raiseproc(sig uint32) {
	pid, _ := syscall0(&libc_getpid)
	syscall2(&libc_kill, uintptr(pid), uintptr(sig))
}

// pthread funcs

func pthread_attr_init(attr *pthreadattr) int32 {
	r, _ := syscall1(&libpthread_attr_init, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_destroy(attr *pthreadattr) int32 {
	r, _ := syscall1(&libpthread_attr_destroy, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_setstack(attr *pthreadattr, stackaddr unsafe.Pointer, stacksize uintptr) int32 {
	r, _ := syscall3(&libpthread_attr_setstack, uintptr(unsafe.Pointer(attr)), uintptr(stackaddr), uintptr(stacksize))
	return int32(r)
}

func pthread_attr_getstack(attr *pthreadattr, stackaddr unsafe.Pointer, stacksize *uintptr) int32 {
	r, _ := syscall3(&libpthread_attr_getstack, uintptr(unsafe.Pointer(attr)), uintptr(stackaddr), uintptr(unsafe.Pointer(stacksize)))
	return int32(r)
}

func pthread_attr_setstacksize(attr *pthreadattr, stacksize uintptr) int32 {
	r, _ := syscall2(&libpthread_attr_setstacksize, uintptr(unsafe.Pointer(attr)), uintptr(stacksize))
	return int32(r)
}

func pthread_attr_getstacksize(attr *pthreadattr, stacksize *uintptr) int32 {
	r, _ := syscall2(&libpthread_attr_getstacksize, uintptr(unsafe.Pointer(attr)), uintptr(unsafe.Pointer(stacksize)))
	return int32(r)
}

func pthread_attr_setdetachstate(attr *pthreadattr, detachstate int32) int32 {
	r, _ := syscall2(&libpthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(detachstate))
	return int32(r)
}

func pthread_create(thread *pthread, attr *pthreadattr, start_routine uintptr, arg unsafe.Pointer) int32 {
	r, _ := syscall4(&libpthread_create, uintptr(unsafe.Pointer(thread)), uintptr(unsafe.Pointer(attr)), uintptr(unsafe.Pointer(start_routine)), uintptr(arg))
	return int32(r)
}
