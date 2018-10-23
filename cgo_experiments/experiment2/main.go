package main

/*
#cgo CFLAGS: -I/Users/brunsgaard/code/pdfiumleg/deps/pdfium/include
#cgo LDFLAGS: -L/Users/brunsgaard/code/pdfiumleg/deps/pdfium/lib -lpdfium
#include "fpdfview.h"
#include "fpdf_annot.h"
#include "fpdf_attachment.h"
#include "fpdf_dataavail.h"
#include "fpdf_edit.h"
#include "fpdf_ext.h"
#include "fpdf_formfill.h"
#include "fpdf_progressive.h"
#include "fpdf_structtree.h"
#include "fpdf_text.h"
#include "fpdfview.h"
*/
import "C"

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"unsafe"
)

// Document is good
type Document struct {
	doc C.FPDF_DOCUMENT
}

// NewDocument shoud have docs
func NewDocument(data *[]byte) (*Document, error) {
	// doc := C.FPDF_LoadDocument(C.CString("in.pdf"), nil)
	doc := C.FPDF_LoadMemDocument(
		unsafe.Pointer(&((*data)[0])),
		C.int(len(*data)),
		nil)
	if doc == nil {
		defer C.FPDF_CloseDocument(doc)

		errorcase := C.FPDF_GetLastError()
		switch errorcase {
		case C.FPDF_ERR_SUCCESS:
			println("FPDF_ERR_SUCCESS")
		case C.FPDF_ERR_UNKNOWN:
			println("FPDF_ERR_UNKNOWN")
		case C.FPDF_ERR_FILE:
			println("FPDF_ERR_FILE")
		case C.FPDF_ERR_FORMAT:
			println("FPDF_ERR_FORMAT")
			println("Unknown error:", errorcase)
		case C.FPDF_ERR_PASSWORD:
			println("FPDF_ERR_PASSWORD")
		case C.FPDF_ERR_SECURITY:
			println("FPDF_ERR_SECURITY")
		case C.FPDF_ERR_PAGE:
			println("FPDF_ERR_PAGE")
		default:
			println("Unknown error:", errorcase)
		}
		return nil, errors.New("Error due to")
	}
	return &Document{doc: doc}, nil
}

// GetPageCount shoud have docs
func (d *Document) GetPageCount() int {
	return int(C.FPDF_GetPageCount(d.doc))
}

// CloseDocument shoud have docs
func (d *Document) CloseDocument() {
	C.FPDF_CloseDocument(d.doc)
}

// RenderPage shoud have docs
func (d *Document) RenderPage(i int) *image.RGBA {
	page := C.FPDF_LoadPage(d.doc, C.int(i))
	scale := 300 / 72
	width := C.int(C.FPDF_GetPageWidth(page) * C.double(scale))
	height := C.int(C.FPDF_GetPageHeight(page) * C.double(scale))
	alpha := C.FPDFPage_HasTransparency(page)
	bitmap := C.FPDFBitmap_Create(width, height, alpha)
	fillColor := 4294967295
	if int(alpha) == 1 {
		fillColor = 0
	}
	C.FPDFBitmap_FillRect(bitmap, 0, 0, width, height, C.ulong(fillColor))
	C.FPDF_RenderPageBitmap(bitmap, page, 0, 0, width, height, 0,
		C.FPDF_ANNOT)
	println(int(C.FPDFBitmap_GetFormat(bitmap)))
	p := C.FPDFBitmap_GetBuffer(bitmap)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	img.Stride = int(C.FPDFBitmap_GetStride(bitmap))
	println(img.Stride)
	println(int(width * 4))
	bgra := make([]byte, 4)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			for i := range bgra {
				bgra[i] = *((*byte)(p))
				p = unsafe.Pointer(uintptr(p) + 1)
			}
			img.SetRGBA(x, y, color.RGBA{B: bgra[0], G: bgra[1], R: bgra[2], A: bgra[3]})

		}
	}
	C.FPDFBitmap_Destroy(bitmap)
	C.FPDF_ClosePage(page)
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	return img
}

func main() {

	data, _ := ioutil.ReadFile("out.png")
	C.FPDF_InitLibrary()
	d, err := NewDocument(&data)
	if err == nil {
		d.GetPageCount()
		d.RenderPage(0)
		d.CloseDocument()
	}
	C.FPDF_DestroyLibrary()
}
