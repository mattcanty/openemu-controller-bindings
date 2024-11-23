package internal

func (g GameMappingConfig) ToOpenEmuBindings(s SystemConfig) map[string]SystemBinding {
	data := map[string]SystemBinding{
		s.Namespace: {
			ControllerBindings: ControllerBinding{},
		},
	}

	for _, controller := range g.Controllers {
		mappingData := make(map[string]string)
		for _, mapping := range controller.Mappings {
			mappingData[mapping.MappingKey(s)] = mapping.MappingValue(controller)
		}

		data[s.Namespace].ControllerBindings[controller.ControllerKey()] = []map[string]string{
			mappingData,
		}
	}

	return data
}
