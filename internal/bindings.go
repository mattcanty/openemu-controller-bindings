package internal

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"howett.net/plist"
)

const (
	bindingsDirMac     = "Library/Application Support/OpenEmu/Bindings"
	bindingsDirLinux   = ".local/share/OpenEmu/Bindings"
	bindingsDirWindows = "OpenEmu/Bindings"
	bindingsFilename   = "Default.oebindings"
)

func WriteBindings(data map[string]SystemBinding) error {
	var buf bytes.Buffer
	encoder := plist.NewEncoderForFormat(&buf, plist.BinaryFormat)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	bindingsDir, err := getBindingsDir()
	if err != nil {
		return err
	}

	bindingsFilePath := filepath.Join(bindingsDir, bindingsFilename)
	return os.WriteFile(bindingsFilePath, buf.Bytes(), 0644)
}

func BackupBindings() error {
	bindingsDir, err := getBindingsDir()
	if err != nil {
		return err
	}

	bindingsFilePath := filepath.Join(bindingsDir, bindingsFilename)

	timestamp := time.Now().Format("20060102-150405")
	backupFilePath := fmt.Sprintf("%s.backup-%s", bindingsFilePath, timestamp)
	source, err := os.Open(bindingsFilePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(backupFilePath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = source.Stat()
	if err != nil {
		return err
	}

	_, err = destination.ReadFrom(source)

	return err
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
