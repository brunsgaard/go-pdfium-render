package main

// http://www.mischiefblog.com/2014/06/26/example-cgo-golang-app-that-calls-a-native-library-with-a-c-structure/

/*
#cgo CFLAGS: -I/Users/brunsgaard/code/go-pdfium-render/cgotesting
#cgo LDFLAGS: -L/Users/brunsgaard/code/go-pdfium-render/cgotesting -lbridge
#include "bridge.h"
*/
import "C"

type Message C.Message

func CreateMessage(msg string) *C.Message {
	cMsg := C.CString(msg)
	return C.create_message(cMsg)
}

func DisplayMessage(msg *C.Message) {
	C.display_message(msg)
}

func FreeMessage(msg *C.Message) {
	C.free_message(msg)
}

func main() {
	msg := CreateMessage("Hello, world!")
	DisplayMessage(msg)
	DisplayMessage(msg)
	FreeMessage(msg)
}
