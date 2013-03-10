package gnureadline

/*
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/readline.h>
#include <readline/history.h>
#include <stdlib.h> // for free()
*/
import "C"
import "unsafe"

func Readline(prompt string) string {
	c_prompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(c_prompt))
	return C.GoString(C.readline(c_prompt))
}

