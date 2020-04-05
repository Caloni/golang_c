package main

import (
	"unsafe"
	"fmt"
)

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L.
////#cgo pkg-config: some-lib-that-uses-pkg-config
//#include "golang_c.h"
//void GoCallback_cgo(int result, void* vpointer, const char* cstring, struct STRUCT* cstruct); // Forward declaration
import "C"

//export GoCallback
func GoCallback(result int, vpointer unsafe.Pointer, cstring *C.char, cstruct *C.struct_STRUCT) {
	vpointer_to_gostring := C.GoString((*C.char)(vpointer))
    gostring := C.GoString(cstring)
    key := "empty"
    value := "empty"
    if cstruct != nil {
        key = C.GoString(cstruct.key)
        value = C.GoString(cstruct.value)
    }
	fmt.Printf("Go: callback called; result:[%d], vpointer:[%s], cstring:[%s], cstruct:[%p], cstruct.key:[%s], cstruct.value:[%s]\n",
        result, vpointer_to_gostring, gostring, cstruct, key, value)
}

func main() {
	string := "some string"
	var vpointer *C.char = C.CString("some string as vpointer")

	fmt.Printf("Go: setting callback at %p\n", unsafe.Pointer(C.GoCallback_cgo));
	C.set_callback((C.Callback)(unsafe.Pointer(C.GoCallback_cgo)))

	fmt.Printf("Go: calling callback at %p passing result:[%d], vpointer:[%p], cstring:[%s]\n", unsafe.Pointer(C.GoCallback_cgo), 5, unsafe.Pointer(vpointer), C.CString(string))
	var result = C.call_callback(5, unsafe.Pointer(vpointer), C.CString(string))
	fmt.Printf("Go: callback returned %d\n", result)
}

