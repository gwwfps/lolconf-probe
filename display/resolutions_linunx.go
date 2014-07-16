package display

func ListAvailableResolutions() (interface{}, error) {
  resolutions := make([]ScreenResolution, 0)
  resolutions = append(resolutions, ScreenResolution{1920, 1080})
  resolutions = append(resolutions, ScreenResolution{1280, 720})
  return AvailableResolutions{resolutions}, nil
}
