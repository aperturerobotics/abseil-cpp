// Package abseil_cpp embeds the abseil-cpp C++ library source files.
package abseil_cpp

import "embed"

// CMakeFiles embeds CMakeLists.txt and cmake helper files needed for building.
//
//go:embed CMakeLists.txt
//go:embed CMake/*.cmake
//go:embed absl/CMakeLists.txt
//go:embed absl/*/CMakeLists.txt
//go:embed absl/copts/*.cmake
var CMakeFiles embed.FS
