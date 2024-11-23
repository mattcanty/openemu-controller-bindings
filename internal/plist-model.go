package internal

type ControllerBinding map[string][]map[string]string

type SystemBinding struct {
	ControllerBindings ControllerBinding `plist:"controllerBindings"`
}
