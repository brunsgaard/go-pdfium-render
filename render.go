package render

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
// #include "fpdf_annot.h"
// #include "fpdf_edit.h"
// #include "fpdf_structtree.h"
import "C"

import (
	"errors"
	"image"
	"image/color"
	"sync"
	"unsafe"
)

// Document is good
type Document struct {
	doc  C.FPDF_DOCUMENT
	data *[]byte // Keep a refrence to the data otherwise wierd stuff happens
}

var mutex = &sync.Mutex{}

// NewDocument shoud have docs
func NewDocument(data *[]byte) (*Document, error) {
	mutex.Lock()
	defer mutex.Unlock()
	// doc := C.FPDF_LoadDocument(C.CString("in.pdf"), nil)
	doc := C.FPDF_LoadMemDocument(
		unsafe.Pointer(&((*data)[0])),
		C.int(len(*data)),
		nil)

	if doc == nil {
		var errMsg string

		//defer C.FPDF_CloseDocument(doc)
		errorcase := C.FPDF_GetLastError()
		switch errorcase {
		case C.FPDF_ERR_SUCCESS:
			errMsg = "Success"
		case C.FPDF_ERR_UNKNOWN:
			errMsg = "Unknown error"
		case C.FPDF_ERR_FILE:
			errMsg = "Unable to read file"
		case C.FPDF_ERR_FORMAT:
			errMsg = "Incorrect format"
		case C.FPDF_ERR_PASSWORD:
			errMsg = "Invalid password"
		case C.FPDF_ERR_SECURITY:
			errMsg = "Invalid encryption"
		case C.FPDF_ERR_PAGE:
			errMsg = "Incorrect page"
		default:
			errMsg = "Unexpected error"
		}
		return nil, errors.New(errMsg)
	}
	return &Document{doc: doc, data: data}, nil
}

// GetPageCount shoud have docs
func (d *Document) GetPageCount() int {
	mutex.Lock()
	defer mutex.Unlock()
	return int(C.FPDF_GetPageCount(d.doc))
}

// CloseDocument shoud have docs
func (d *Document) Close() {
	mutex.Lock()
	C.FPDF_CloseDocument(d.doc)
	mutex.Unlock()
}

// RenderPage should have docs
func (d *Document) RenderPage(i int, dpi int) *image.RGBA {
	mutex.Lock()

	page := C.FPDF_LoadPage(d.doc, C.int(i))
	scale := float64(dpi) / 72.0
	imgWidth := C.FPDF_GetPageWidth(page) * C.double(scale)
	imgHeight := C.FPDF_GetPageHeight(page) * C.double(scale)

	// pixelBound := int(dpi * (3508 / 300))
	// imgWidthRatio := float64(pixelBound) / float64(imgWidth)
	// imgHeightRatio := float64(pixelBound) / float64(imgHeight)
	// scaleFactor := math.Min(imgWidthRatio, imgHeightRatio)
	scaleFactor := 1.0

	width := C.int(imgWidth * C.double(scaleFactor))
	height := C.int(imgHeight * C.double(scaleFactor))

	alpha := C.FPDFPage_HasTransparency(page)

	bitmap := C.FPDFBitmap_Create(width, height, alpha)

	fillColor := 4294967295
	if int(alpha) == 1 {
		fillColor = 0
	}
	C.FPDFBitmap_FillRect(bitmap, 0, 0, width, height, C.ulong(fillColor))
	C.FPDF_RenderPageBitmap(bitmap, page, 0, 0, width, height, 0, C.FPDF_ANNOT|C.FPDF_GRAYSCALE)

	p := C.FPDFBitmap_GetBuffer(bitmap)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	img.Stride = int(C.FPDFBitmap_GetStride(bitmap))
	mutex.Unlock()

	// This takes a bit of time and I *think* we can do this without the lock
	bgra := make([]byte, 4)
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			for i := range bgra {
				bgra[i] = *((*byte)(p))
				p = unsafe.Pointer(uintptr(p) + 1)
			}
			color := color.RGBA{B: bgra[0], G: bgra[1], R: bgra[2], A: bgra[3]}
			img.SetRGBA(x, y, color)
		}
	}
	mutex.Lock()
	C.FPDFBitmap_Destroy(bitmap)
	C.FPDF_ClosePage(page)
	mutex.Unlock()

	// should maybe return err
	//println(C.FPDF_GetLastError())

	return img
}

func InitLibrary() {
	mutex.Lock()
	C.FPDF_InitLibrary()
	mutex.Unlock()
}

func DestroyLibrary() {
	mutex.Lock()
	C.FPDF_DestroyLibrary()
	mutex.Unlock()
}
