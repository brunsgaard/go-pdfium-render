cmake_minimum_required(VERSION 3.2)
project(pdfiumtk)

set(SOURCE_FILES main.cpp)

add_executable(pti ${SOURCE_FILES})
target_link_libraries(pti pdfium)

install(TARGETS pti DESTINATION ${PDFIUMTK_INSTALL_BIN_DIR})
