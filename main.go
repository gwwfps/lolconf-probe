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
  user32              = syscall.NewLazyDLL("User32.dll")
  enumDisplaySettings = user32.NewProc("EnumDisplaySettingsW")
)

// type PDevMode struct {
//   dmOrientation int16
//   dmPaperSize int16
//   dmPaperLength int16
//   dmPaperWidth int16
//   dmScale int16
//   dmCopies int16
//   dmDefaultSource int16
//   dmPrintQuality int16
// }

// type PointL struct {
//   x int32
//   y int32
// }

// type LPDevMode struct {
//   dmPosition PointL
//   dmDisplayOrientation uint32
//   dmDisplayFixedOutput uint32
// }

type DevMode struct {
  dmDeviceName    string
  dmSpecVersion   uint16
  dmDriverVersion uint16
  dmSize          uint16
  dmDriverExtra   uint16
  dmFields        uint32

  submode uintptr

  dmColor       int16
  dmDuplex      int16
  dmYResolution int16
  dmTTOption    int16
  dmCollate     int16
  dmFormName    string

  dmLogPixels        uint16
  dmBitsPerPel       uint32
  dmPelsWidth        uint32
  dmPelsHeight       uint32
  flagsOrNup         uintptr
  dmDisplayFrequency uint32
  dmICMMethod        uint32
  dmICMIntent        uint32
  dmMediaType        uint32
  dmDitherType       uint32
  dmReserved1        uint32
  dmReserved2        uint32
  dmPanningWidth     uint32
  dmPanningHeight    uint32
}

func main() {
  var iModeNum = 0
  var devMode = new(DevMode)
  _, _, callErr := syscall.Syscall(enumDisplaySettings.Addr(),
    3,
    0,
    uintptr(iModeNum),
    uintptr(unsafe.Pointer(&devMode)))
  if callErr != 0 {
    abort("failed call", callErr)
  }
  fmt.Printf("%+v", devMode)
}
