package main

import (
  "fmt"
  "syscall"
  "unsafe"
)

func abort(funcname string, err error) {
  panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
  user32              = syscall.NewLazyDLL("user32.dll")
  enumDisplaySettings = user32.NewProc("EnumDisplaySettingsExW")
)

type SimpleDevMode struct {
  filler1 [34]uint16

  Size uint16

  filler2 [50]uint16

  Width  uint32
  Height uint32

  filler3 [10]uint32
}

func invokeEnumDisplaySettings(iModeNum int) *SimpleDevMode {
  var devMode = new(SimpleDevMode)
  devMode.DmSize = uint16(unsafe.Sizeof(SimpleDevMode{}))

  result, _, _ := enumDisplaySettings.Call(
    uintptr(unsafe.Pointer(nil)),
    uintptr(iModeNum),
    uintptr(unsafe.Pointer(devMode)),
    uintptr(0))

  if result == 1 {
    return devMode
  } else {
    return nil
  }
}

func main() {
  for iModeNum := 0; ; iModeNum++ {
    devMode := invokeEnumDisplaySettings(iModeNum)
    if devMode == nil {
      break
    }
    fmt.Printf("%dx%d\n", devMode.DmPelsWidth, devMode.DmPelsHeight)
  }
}
