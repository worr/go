// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// WARNING: This file has automatically been generated on Tue, 03 Dec 2019 22:25:52 CET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package openbsd

/*
#include "pthread"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocPthreadattrMemory allocates memory for type C.pthread_attr_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPthreadattrMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPthreadattrValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPthreadattrValue = unsafe.Sizeof([1]C.pthread_attr_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *pthreadattr) Ref() *C.pthread_attr_t {
	if x == nil {
		return nil
	}
	return x.refabfcad8c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *pthreadattr) Free() {
	if x != nil && x.allocsabfcad8c != nil {
		x.allocsabfcad8c.(*cgoAllocMap).Free()
		x.refabfcad8c = nil
	}
}

// NewpthreadattrRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewpthreadattrRef(ref unsafe.Pointer) *pthreadattr {
	if ref == nil {
		return nil
	}
	obj := new(pthreadattr)
	obj.refabfcad8c = (*C.pthread_attr_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *pthreadattr) PassRef() (*C.pthread_attr_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refabfcad8c != nil {
		return x.refabfcad8c, nil
	}
	memabfcad8c := allocPthreadattrMemory(1)
	refabfcad8c := (*C.pthread_attr_t)(memabfcad8c)
	allocsabfcad8c := new(cgoAllocMap)
	allocsabfcad8c.Add(memabfcad8c)

	var cstack_addr_allocs *cgoAllocMap
	refabfcad8c.stack_addr, cstack_addr_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.stack_addr)), cgoAllocsUnknown
	allocsabfcad8c.Borrow(cstack_addr_allocs)

	var cstack_size_allocs *cgoAllocMap
	refabfcad8c.stack_size, cstack_size_allocs = (C.size_t)(x.stack_size), cgoAllocsUnknown
	allocsabfcad8c.Borrow(cstack_size_allocs)

	var cguard_size_allocs *cgoAllocMap
	refabfcad8c.guard_size, cguard_size_allocs = (C.size_t)(x.guard_size), cgoAllocsUnknown
	allocsabfcad8c.Borrow(cguard_size_allocs)

	var cdetach_state_allocs *cgoAllocMap
	refabfcad8c.detach_state, cdetach_state_allocs = (C.int)(x.detach_state), cgoAllocsUnknown
	allocsabfcad8c.Borrow(cdetach_state_allocs)

	var ccontention_scope_allocs *cgoAllocMap
	refabfcad8c.contention_scope, ccontention_scope_allocs = (C.int)(x.contention_scope), cgoAllocsUnknown
	allocsabfcad8c.Borrow(ccontention_scope_allocs)

	var csched_policy_allocs *cgoAllocMap
	refabfcad8c.sched_policy, csched_policy_allocs = (C.int)(x.sched_policy), cgoAllocsUnknown
	allocsabfcad8c.Borrow(csched_policy_allocs)

	var csched_param_allocs *cgoAllocMap
	refabfcad8c.sched_param, csched_param_allocs = x.sched_param.PassValue()
	allocsabfcad8c.Borrow(csched_param_allocs)

	var csched_inherit_allocs *cgoAllocMap
	refabfcad8c.sched_inherit, csched_inherit_allocs = (C.int)(x.sched_inherit), cgoAllocsUnknown
	allocsabfcad8c.Borrow(csched_inherit_allocs)

	x.refabfcad8c = refabfcad8c
	x.allocsabfcad8c = allocsabfcad8c
	return refabfcad8c, allocsabfcad8c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x pthreadattr) PassValue() (C.pthread_attr_t, *cgoAllocMap) {
	if x.refabfcad8c != nil {
		return *x.refabfcad8c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *pthreadattr) Deref() {
	if x.refabfcad8c == nil {
		return
	}
	x.stack_addr = (unsafe.Pointer)(unsafe.Pointer(x.refabfcad8c.stack_addr))
	x.stack_size = (uint)(x.refabfcad8c.stack_size)
	x.guard_size = (uint)(x.refabfcad8c.guard_size)
	x.detach_state = (int32)(x.refabfcad8c.detach_state)
	x.contention_scope = (int32)(x.refabfcad8c.contention_scope)
	x.sched_policy = (int32)(x.refabfcad8c.sched_policy)
	x.sched_param = *Newsched_paramRef(unsafe.Pointer(&x.refabfcad8c.sched_param))
	x.sched_inherit = (int32)(x.refabfcad8c.sched_inherit)
}

// allocStruct_sched_paramMemory allocates memory for type C.struct_sched_param in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_sched_paramMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_sched_paramValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStruct_sched_paramValue = unsafe.Sizeof([1]C.struct_sched_param{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *sched_param) Ref() *C.struct_sched_param {
	if x == nil {
		return nil
	}
	return x.ref6ecda53c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *sched_param) Free() {
	if x != nil && x.allocs6ecda53c != nil {
		x.allocs6ecda53c.(*cgoAllocMap).Free()
		x.ref6ecda53c = nil
	}
}

// Newsched_paramRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsched_paramRef(ref unsafe.Pointer) *sched_param {
	if ref == nil {
		return nil
	}
	obj := new(sched_param)
	obj.ref6ecda53c = (*C.struct_sched_param)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *sched_param) PassRef() (*C.struct_sched_param, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6ecda53c != nil {
		return x.ref6ecda53c, nil
	}
	mem6ecda53c := allocStruct_sched_paramMemory(1)
	ref6ecda53c := (*C.struct_sched_param)(mem6ecda53c)
	allocs6ecda53c := new(cgoAllocMap)
	allocs6ecda53c.Add(mem6ecda53c)

	var csched_priority_allocs *cgoAllocMap
	ref6ecda53c.sched_priority, csched_priority_allocs = (C.int)(x.sched_priority), cgoAllocsUnknown
	allocs6ecda53c.Borrow(csched_priority_allocs)

	x.ref6ecda53c = ref6ecda53c
	x.allocs6ecda53c = allocs6ecda53c
	return ref6ecda53c, allocs6ecda53c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x sched_param) PassValue() (C.struct_sched_param, *cgoAllocMap) {
	if x.ref6ecda53c != nil {
		return *x.ref6ecda53c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *sched_param) Deref() {
	if x.ref6ecda53c == nil {
		return
	}
	x.sched_priority = (int32)(x.ref6ecda53c.sched_priority)
}
