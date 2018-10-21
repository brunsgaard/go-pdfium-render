# go-pdfium-render

# The project

The end goal is to have a golang libary that uses pdfium for rendering pdfs into images ppm(especially pgm) and png, everything shuld happen in memory.

Why pdfium? Because tt is the best in class and because it has a BSD style license, also there will be continued development on pdfium as long as chrome is around. [0]

The steps here are

1. Checkout and compile pdfium [2]
2. Import pdfium into a small c++ application. (main.cc) [6]
3. Create a class in c++ that performs the operation we need [1][6]
4. Read up on cgo and how to glue c++ and go together. [3][4]
5. Create a go libary that binds to the c++ class [5][6]

Resources

[0] https://tinyurl.com/yd5fb2rz   
[1] https://developers.google.com/edu/c++/   
[2] https://pdfium.googlesource.com/pdfium/+/HEAD/docs/getting-started.md   
[3] https://karthikkaranth.me/blog/calling-c-code-from-go/   
[4] https://github.com/arrieta/golang-cpp-basic-example   
[5] https://groups.google.com/forum/#!topic/pdfium/r6KCGo6q7Fo   
[6] https://github.com/cgilling/build-pdfium   


Misc
Just found up to date binaries for pdfium.
https://github.com/bblanchon/pdfium-binaries
