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
  user32, _              = syscall.LoadLibrary("user32.dll")
  enumDisplaySettings, _ = syscall.GetProcAddress(user32, "EnumDisplaySettingsW")
)

type DevMode struct {
  dmDeviceName    [32]byte
  dmSpecVersion   uint16
  dmDriverVersion uint16
  dmSize          uint16
  dmDriverExtra   uint16
  dmFields        uint32

  dmOrientation   uint16
  dmPaperSize     uint16
  dmPaperLength   uint16
  dmPaperWidth    uint16
  dmScale         int16
  dmCopies        int16
  dmDefaultSource uint16
  dmPrintQuality  uint16

  dmColor       int16
  dmDuplex      int16
  dmYResolution int16
  dmTTOption    int16
  dmCollate     int16
  dmFormName    [32]byte

  dmLogPixels        uint16
  dmBitsPerPel       uint32
  dmPelsWidth        uint32
  dmPelsHeight       uint32
  dmDisplayFlags     uint32
  dmDisplayFrequency uint32
}

func invokeEnumDisplaySettings(iModeNum int) *DevMode {
  var devMode = new(DevMode)
  devMode.dmSize = uint16(unsafe.Sizeof(DevMode{}))
  result, _, callErr := syscall.Syscall(uintptr(enumDisplaySettings), 3, 0, uintptr(iModeNum), uintptr(unsafe.Pointer(devMode)))
  if callErr != 0 {
    abort("failed call", callErr)
  }
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
    fmt.Println(devMode)
  }
}
