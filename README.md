# ssn-pdf-renderkit

# The project

The end goal is to have a golang libary uses pdfium for rendering images, everything shuld happen in memory.

The steps here are

1. Learn yourself a minimal amount of c++
2. Checkout and compile pdfium
3. Import pdfium into a smat c++ application. (main.cc)
4. Create a class in c++ that performs the operation we need
5. Read up on cgo and how to glue c++ and go together.
6. Create a go libary that binds to the c++ class

Resources 

https://developers.google.com/edu/c++/
https://community.alfresco.com/community/ecm/blog/2018/05/16/pdf-rendering-engine-performance-and-fidelity-comparison

https://karthikkaranth.me/blog/calling-c-code-from-go/
https://github.com/arrieta/golang-cpp-basic-example

https://groups.google.com/forum/#!searchin/pdfium/FPDF_InitLibrary$20golang%7Csort:date/pdfium/r6KCGo6q7Fo/VNOneweGAwAJ
https://github.com/cgilling/build-pdfium
