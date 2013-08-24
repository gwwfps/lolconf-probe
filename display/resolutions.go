package display

import (
  "syscall"
  "unsafe"
)

var (
  user32              = syscall.NewLazyDLL("user32.dll")
  enumDisplaySettings = user32.NewProc("EnumDisplaySettingsExW")
)

type ScreenResolution struct {
  Width  uint8
  Height uint8
}

type simpleDevMode struct {
  filler1 [34]uint16

  Size uint16

  filler2 [50]uint16

  Width  uint32
  Height uint32

  filler3 [10]uint32
}

func invokeEnumDisplaySettings(iModeNum int) *simpleDevMode {
  var devMode = new(simpleDevMode)
  devMode.DmSize = uint16(unsafe.Sizeof(simpleDevMode{}))

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

func ListAvailableResolutions() []ScreenResolution {
  resolutions := make([]ScreenResolution)

  for iModeNum := 0; ; iModeNum++ {
    devMode := invokeEnumDisplaySettings(iModeNum)
    if devMode == nil {
      break
    }

    found := false
    for _, res := range resolutions {
      if res.Width == devMode.Width && res.Height == devMode.Height {
        found = true
      }
    }

    if !found {
      res := new(ScreenResolution)
      res.Width = devMode.Width
      res.Height = devMode.Height
      append(resolutions, &res)
    }
  }

  return resolutions
}
