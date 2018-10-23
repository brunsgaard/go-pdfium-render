package main

/*
#cgo CFLAGS: -I/Users/brunsgaard/code/pdfiumleg/deps/pdfium/include
#cgo LDFLAGS: -L/Users/brunsgaard/code/pdfiumleg/deps/pdfium/lib -lpdfium
#include "fpdfview.h"
*/
import "C"

func main() {
	C.FPDF_InitLibrary()
	C.FPDF_DestroyLibrary()
}
