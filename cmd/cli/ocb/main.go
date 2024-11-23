package main

import (
	"fmt"
	"os"
	"path"

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
		mappingFiles, err := internal.ListAvailableMappingFiles("config")
		if err != nil {
			return err
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

func main() {
	rootCmd := &cobra.Command{Use: "ocb"}
	rootCmd.AddCommand(loadCmd, listCmd, backupCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
