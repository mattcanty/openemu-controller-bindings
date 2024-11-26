package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mattcanty/openemu-controller-bindings/internal"
	"github.com/spf13/cobra"
)

const ()

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a controller mapping",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configFilePath := path.Join("config", fmt.Sprintf("%s.yaml", args[0]))

		config, err := internal.LoadGameMappingConfig(configFilePath)
		if err != nil {
			return err
		}

		system, err := internal.LoadSystemConfig(config)
		if err != nil {
			return err
		}

		err = internal.WriteBindings(config.ToOpenEmuBindings(system))
		if err != nil {
			return err
		}

		fmt.Printf("Loaded '%s'\n", args[0])

		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved binding configurations",
	RunE: func(cmd *cobra.Command, args []string) error {
		valid, err := validateMappingsDir()
		if err != nil {
			return err
		}

		if !valid {
			return nil
		}

		mappingFiles, err := internal.ListAvailableMappingFiles()
		if err != nil {
			return err
		}

		if len(mappingFiles) == 0 {
			fmt.Println("No mapping files found")
			return nil
		}

		for _, item := range mappingFiles {
			fmt.Printf("%s\n", item)
		}

		return nil
	},
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Creates a time-stamped backup of the current bindings file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.BackupBindings()
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the mappings directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Initializing mappings directory")

		return internal.InitMappingsDir()
	},
}

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture a controller mapping",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		system := args[0]
		game := strings.ToLower(strings.ReplaceAll(args[1], " ", "-"))

		if _, exists := internal.SystemConfigs[system]; !exists {
			fmt.Printf("invalid system: %s\n", system)
			fmt.Println("\nAvailable systems:")

			systems := make([]string, 0, len(internal.SystemConfigs))
			for system := range internal.SystemConfigs {
				systems = append(systems, system)
			}
			fmt.Println(strings.Join(systems, ", "))

			return nil
		}

		fmt.Printf("Capturing mapping for %s-%s\n", system, game)

		gameMappingConfig, err := internal.CaptureMapping(system, game)
		if err != nil {
			return err
		}

		mappingsDir, err := internal.MappingsDir()
		if err != nil {
			return err
		}

		configFilePath := strings.ToLower(path.Join(mappingsDir, fmt.Sprintf("%s-%s.yaml", system, game)))

		err = internal.SaveGameMappingConfig(configFilePath, gameMappingConfig)
		if err != nil {
			return err
		}

		fmt.Printf("Saved mapping to %s\n", configFilePath)

		return nil
	},
}

func main() {
	rootCmd := &cobra.Command{Use: "ocb"}
	rootCmd.AddCommand(captureCmd, initCmd, loadCmd, listCmd, backupCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateMappingsDir() (bool, error) {
	exists, err := internal.MappingsDirExists()
	if err != nil {
		return false, err
	}

	mappingsDir, err := internal.MappingsDir()
	if err != nil {
		return false, err
	}

	if !exists {
		fmt.Printf("Ensure mappings directory exists at %s\n", mappingsDir)
		return false, nil
	}

	return true, nil
}
