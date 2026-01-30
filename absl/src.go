package absl

import "embed"

// Source embeds the Abseil C++ source files for vendoring.
//
// CMake files for building
//go:embed CMakeLists.txt
//go:embed */CMakeLists.txt
//go:embed algorithm/*.h
//go:embed base/*.cc base/*.h
//go:embed base/internal/*.cc base/internal/*.h base/internal/*.inc
//go:embed cleanup/*.h
//go:embed cleanup/internal/*.h
//go:embed container/*.cc container/*.h
//go:embed container/internal/*.cc container/internal/*.h
//go:embed crc/*.cc crc/*.h
//go:embed crc/internal/*.cc crc/internal/*.h
//go:embed debugging/*.cc debugging/*.h debugging/*.inc
//go:embed debugging/internal/*.cc debugging/internal/*.h debugging/internal/*.inc
//go:embed flags/*.cc flags/*.h
//go:embed flags/internal/*.cc flags/internal/*.h
//go:embed functional/*.h
//go:embed functional/internal/*.h
//go:embed hash/*.cc hash/*.h
//go:embed hash/internal/*.cc hash/internal/*.h
//go:embed log/*.cc log/*.h log/*.inc
//go:embed log/internal/*.cc log/internal/*.h
//go:embed memory/*.h
//go:embed meta/*.h
//go:embed meta/internal/*.cc meta/internal/*.h
//go:embed extend/internal/*.cc extend/internal/*.h
//go:embed numeric/*.cc numeric/*.h numeric/*.inc
//go:embed numeric/internal/*.h
//go:embed profiling/*.cc profiling/*.h
//go:embed profiling/internal/*.cc profiling/internal/*.h
//go:embed random/*.cc random/*.h
//go:embed random/internal/*.cc random/internal/*.h
//go:embed status/*.cc status/*.h
//go:embed status/internal/*.cc status/internal/*.h
//go:embed strings/*.cc strings/*.h
//go:embed strings/internal/*.cc strings/internal/*.h
//go:embed strings/internal/str_format/*.cc strings/internal/str_format/*.h
//go:embed synchronization/*.cc synchronization/*.h
//go:embed synchronization/internal/*.cc synchronization/internal/*.h
//go:embed time/*.cc time/*.h
//go:embed time/internal/*.cc time/internal/*.h time/internal/*.inc
//go:embed time/internal/cctz/include/cctz/*.h
//go:embed time/internal/cctz/src/*.cc time/internal/cctz/src/*.h
//go:embed types/*.cc types/*.h
//go:embed types/internal/*.h
//go:embed utility/*.h
var Source embed.FS
