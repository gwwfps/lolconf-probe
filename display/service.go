package display

type DisplayService struct{}
type EmptyArgs struct{}

func (s *DisplayService) ListAvailableResolutions(args *EmptyArgs, reply *[]ScreenResolution) error {
  resolutions := ListAvailableResolutions()
  reply = &resolutions
  return nil
}
