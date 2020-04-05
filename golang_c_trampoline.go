package main

/*
#include "golang_c.h"

// The gateway functions
//

void GoCallback_cgo(int result, void* vpointer, const char* cstring, struct STRUCT* cstruct) {
	void GoCallback(int result, void* vpointer, const char* cstring, struct STRUCT* cstruct);
	GoCallback(result, vpointer, cstring, cstruct);
}
*/
import "C"
