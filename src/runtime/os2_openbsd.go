// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

const (
	_SS_DISABLE              = 4
	_SIG_BLOCK               = 1
	_SIG_UNBLOCK             = 2
	_SIG_SETMASK             = 3
	_NSIG                    = 33
	_SI_USER                 = 0
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

//go:cgo_import_dynamic libpthread_attr_init pthread_attr_init "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_destroy pthread_attr_destroy "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_setstack pthread_attr_setstack "libpthread.so"
//go:cgo_import_dynamic libpthread_attr_setstacksize pthread_attr_setstacksize "libpthread.so:
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

//go:linkname libpthread_attr_init libpthread_attr_init
//go:linkname libpthread_attr_destroy libpthread_attr_destroy
//go:linkname libpthread_attr_setstack libpthread_attr_setstack
//go:linkname libpthread_attr_setstacksize libpthread_attr_setstacksize
//go:linkname libpthread_attr_setdetachstate libpthread_attr_setdetachstate
//go:linkname libpthread_create libpthread_create

var (
	__errno        libcFunc
	libc_read      libcFunc
	libc_close     libcFunc
	libc_exit      libcFunc
	libc_nanosleep libcFunc
	libc_write     libcFunc
	libc_open      libcFunc
	libc_madvise   libcFunc
	libc___threxit libcFunc

	libpthread_attr_init         libcFunc
	libpthread_attr_destroy      libcFunc
	libpthread_attr_setstack     libcFunc
	libpthread_attr_setstacksize libcFunc
	libpthread_create            libcFunc
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

	syscall1(&libc_nanosleep, uintptr(unsafe.Pointer(ts)))
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
	r, _ := syscall3(&libc_madvise, uintptr(unsafe.Pointer(addr)), n, uintptr(flags))
	return int32(r)
}

func exitThread(wait *uint32) {
	syscall1(&libc___threxit, uintptr(unsafe.Pointer(wait)))
}

func pthread_attr_init(attr *pthreadattr) int32 {
	r, _ := syscall1(&libpthread_attr_init, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_destroy(attr *pthreadattr) int32 {
	r, _ := syscall1(&libpthread_attr_destroy, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_setstack(attr *pthreadattr, stackaddr unsafe.Pointer, stacksize int32) int32 {
	r, _ := syscall3(&libpthread_attr_setstack, uintptr(unsafe.Pointer(attr)), uintptr(stackaddr), uintptr(stacksize))
	return int32(r)
}

func pthread_attr_setdetachstate(attr *pthreadattr, detachstate int32) int32 {
	r, _ := syscall3(&libpthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(detachstate))
	return int32(r)
}
