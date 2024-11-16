package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const (
	bindingsDirMac         = "Library/Application Support/OpenEmu/Bindings/"
	bindingsDirLinux       = ".local/share/OpenEmu/Bindings/"
	bindingsDirWindows     = "OpenEmu/Bindings/"
	bindingsFilenameFormat = "%s.oebindings"
	defaultFilename        = "Default.oebindings"
	backupFilename         = "Default.BACKUP.oebindings"
	suffix                 = ".oebindings"
)

var saveCmd = &cobra.Command{
	Use:   "save <NAME>",
	Short: "Save the current bindings to a new file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		bindingsDir, err := getBindingsDir()
		if err != nil {
			return err
		}

		defaultFile := filepath.Join(bindingsDir, defaultFilename)
		newFilename := fmt.Sprintf(bindingsFilenameFormat, name)
		newFile := filepath.Join(bindingsDir, newFilename)

		if err := copyFile(defaultFile, newFile); err != nil {
			return err
		}

		fmt.Printf("Controller bindings '%s' saved.\n", name)

		return nil
	},
}

var loadCmd = &cobra.Command{
	Use:   "load <NAME>",
	Short: "Load a saved binding configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		bindingsDir, err := getBindingsDir()
		if err != nil {
			return err
		}

		if err := doBackup(); err != nil {
			return err
		}

		defaultFile := filepath.Join(bindingsDir, defaultFilename)
		savedFilename := fmt.Sprintf(bindingsFilenameFormat, name)
		savedFile := filepath.Join(bindingsDir, savedFilename)

		if err := copyFile(savedFile, defaultFile); err != nil {
			return err
		}

		fmt.Printf("Controller bindings '%s' loaded. Restart OpenEmu.\n", name)

		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved binding configurations",
	RunE: func(cmd *cobra.Command, args []string) error {
		bindingsDir, err := getBindingsDir()
		if err != nil {
			return err
		}

		files, err := os.ReadDir(bindingsDir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if file.Name() != defaultFilename && file.Name() != backupFilename {
				fmt.Println(strings.TrimSuffix(file.Name(), suffix))
			}
		}
		return nil
	},
}

func main() {
	rootCmd := &cobra.Command{Use: "ocb"}
	rootCmd.AddCommand(saveCmd, loadCmd, listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func doBackup() error {
	bindingsDir, err := getBindingsDir()
	if err != nil {
		return err
	}

	defaultFile := filepath.Join(bindingsDir, defaultFilename)
	backupFile := filepath.Join(bindingsDir, backupFilename)

	if err := copyFile(defaultFile, backupFile); err != nil {
		return err
	}

	return nil
}

func getBindingsDir() (string, error) {
	var bindingsDir string

	switch runtime.GOOS {
	case "darwin": // macOS
		bindingsDir = filepath.Join(os.Getenv("HOME"), bindingsDirMac)
	case "linux": // Linux
		bindingsDir = filepath.Join(os.Getenv("HOME"), bindingsDirLinux)
	case "windows": // Windows
		bindingsDir = filepath.Join(os.Getenv("APPDATA"), bindingsDirWindows)
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return bindingsDir, nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
