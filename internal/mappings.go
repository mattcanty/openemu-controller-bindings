package internal

import (
	"os"
	"strings"
)

func InitMappingsDir() error {
	mappingsDir, err := MappingsDir()
	if err != nil {
		return err
	}

	return os.MkdirAll(mappingsDir, 0755)
}

func CaptureMapping(system, game string) (GameMappingConfig, error) {
	bindings, err := loadBindings()
	if err != nil {
		return GameMappingConfig{}, err
	}

	ns := SystemConfigs[system].Namespace
	systemPrefix := SystemConfigs[system].Prefix

	systemBinding := bindings[ns]

	gameMappingConfig := GameMappingConfig{
		System: system,
	}
	for name, bindings := range systemBinding.ControllerBindings {
		controller := Controller{
			Name: strings.TrimPrefix(name, "OEController"),
		}

		for _, bindingDict := range bindings {
			for key, value := range bindingDict {
				controller.Mappings = append(controller.Mappings, Mapping{
					System:     strings.TrimPrefix(key, systemPrefix),
					Controller: strings.TrimPrefix(value, name),
				})
			}
		}

		gameMappingConfig.Controllers = append(gameMappingConfig.Controllers, controller)
	}

	return gameMappingConfig, nil
}
