package display

type AvailableResolutions struct {
  Resolutions []ScreenResolution `json:"resolutions"`
}

type ScreenResolution struct {
  Width  uint32 `json:"width"`
  Height uint32 `json:"height"`
}
