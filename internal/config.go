package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type SystemsConfig map[string]SystemConfig

type SystemConfig struct {
	Namespace string `yaml:"Namespace"`
	Prefix    string `yaml:"Prefix"`
}

type GameMappingConfig struct {
	System      string       `yaml:"System"`
	Controllers []Controller `yaml:"Controllers"`
}

type Controller struct {
	Name     string    `yaml:"Name"`
	Mappings []Mapping `yaml:"Mappings"`
}

type Mapping struct {
	System     string `yaml:"System"`
	Controller string `yaml:"Controller"`
}

func (c Controller) ControllerKey() string {
	return fmt.Sprintf("OEController%s", c.Name)
}

func (m Mapping) MappingKey(c SystemConfig) string {
	return fmt.Sprintf("%s%s", c.Prefix, m.System)
}

func (m Mapping) MappingValue(c Controller) string {
	return fmt.Sprintf("OEController%s%s", c.Name, m.Controller)
}

func loadYamlConfig[T any](filename string) (T, error) {
	var target T
	file, err := os.Open(filename)
	if err != nil {
		return target, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&target); err != nil {
		return target, err
	}

	return target, nil
}

func LoadGameMappingConfig(filename string) (GameMappingConfig, error) {
	return loadYamlConfig[GameMappingConfig](filename)
}

func LoadSystemsConfig() (SystemsConfig, error) {
	return loadYamlConfig[SystemsConfig]("config/systems.yaml")
}

func LoadSystemConfig(g GameMappingConfig) (SystemConfig, error) {
	systems, err := LoadSystemsConfig()
	if err != nil {
		return SystemConfig{}, err
	}

	return systems[g.System], nil
}

func ListAvailableMappingFiles(configDir string) ([]string, error) {
	var result []string

	err := filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			base := info.Name()

			if base == "systems.yaml" {
				return nil
			}

			if strings.HasSuffix(base, ".yaml") {
				trimmedName := strings.TrimSuffix(base, ".yaml")
				result = append(result, trimmedName)
			}
		}
		return nil
	})

	return result, err
}
