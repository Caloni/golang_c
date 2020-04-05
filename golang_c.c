#include "golang_c.h"
#include <stdio.h>

Callback g_callback = NULL;

void set_callback(Callback callback) {
    printf("C: callback set to %p\n", callback);
    g_callback = callback;
}

int call_callback(int result, void* vpointer, const char* cstring) {
    struct STRUCT s = { "key", "value" };
    if( g_callback ) {
        printf("C: calling callback at %p passing result:[%d], vpointer:[%p], cstring:[%s], cstruct:[%p]\n", g_callback, result, vpointer, cstring, &s);
        result = g_callback(result, vpointer, cstring, &s);
        printf("C: callback at %p called; result:[%d]\n", g_callback, result);
    } else {
        printf("C: no callback to call\n");
    }
    return result;
}

