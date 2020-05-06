cmake_minimum_required(VERSION 2.6)
project(hello)

set (HELLO_VERSION_MAJOR 1)
set (HELLO_VERSION_MINOR 0)

# message("PROJECT_BINARY_DIR = ${PROJECT_BINARY_DIR}")
# message("PROJECT_SOURCE_DIR = ${PROJECT_SOURCE_DIR}")

option (USE_MATH
    "Use tutorial provided math implementation" ON)
if (USE_MATH)
    include_directories("${PROJECT_SOURCE_DIR}/math")
    add_subdirectory ("${PROJECT_SOURCE_DIR}/math")
    SET (EXTRA_LIBS ${EXTRA_LIBS} math)
endif (USE_MATH)

configure_file (
  "${PROJECT_SOURCE_DIR}/src/config.h.in"
  "${PROJECT_BINARY_DIR}/include/config.h"
)
include_directories("${PROJECT_BINARY_DIR}/include")



add_executable(hello src/hello.c)
target_link_libraries (hello math)