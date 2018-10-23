# go-pdfium-render

https://github.com/bblanchon/pdfium-binaries needs to be present on your system.

```
prefix=/opt/pdfium
libdir=/opt/pdfium/lib
includedir=/opt/pdfium/include

Name: pdfium
Description: pdfium
Version: 3580
Requires:

Libs: -L${libdir} -lpdfium
Cflags: -I${includedir}
```

I wrote my own .pc file.


### The project

The end goal is to have a golang libary that uses pdfium for rendering pdfs into images ppm(especially pgm) and png, everything should happen in memory.

Why pdfium? Because tt is the best in class and because it has a BSD style license, also there will be continued development on pdfium as long as chrome is around. [0]

### Resources

[0] https://tinyurl.com/yd5fb2rz   
[2] https://pdfium.googlesource.com/pdfium/+/HEAD/docs/getting-started.md   
[3] https://karthikkaranth.me/blog/calling-c-code-from-go/   
[4] https://github.com/arrieta/golang-cpp-basic-example   
[5] https://groups.google.com/forum/#!topic/pdfium/r6KCGo6q7Fo   
[6] https://github.com/cgilling/build-pdfium   
[7] http://cdn01.foxitsoftware.com/pub/foxit/manual/enu/FoxitPDF_SDK20_Guide.pdf   
[8] https://github.com/bblanchon/pdfium-binaries   
[9] https://tinyurl.com/y8s3aen5   (libvips)     
