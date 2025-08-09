package library

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// ListPlugins scans the library directory and returns a list of all plugins found.
func ListPlugins(libraryDir string) ([]Plugin, error) {
	plugins := make([]Plugin, 0)
	entries, err := os.ReadDir(libraryDir)
	if err != nil {
		if os.IsNotExist(err) {
			return plugins, nil // Return empty list if library dir doesn't exist yet
		}
		return nil, fmt.Errorf("could not read library directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			pluginID := entry.Name()
			pluginJSONPath := filepath.Join(libraryDir, pluginID, "plugin.json")

			if _, err := os.Stat(pluginJSONPath); os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Warning: plugin.json not found for plugin %s, skipping\n", pluginID)
				continue
			}

			data, err := os.ReadFile(pluginJSONPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not read plugin.json for plugin %s: %v\n", pluginID, err)
				continue
			}

			var plugin Plugin
			if err := json.Unmarshal(data, &plugin); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not parse plugin.json for plugin %s: %v\n", pluginID, err)
				continue
			}
			plugins = append(plugins, plugin)
		}
	}
	return plugins, nil
}

// AddPluginFromZip unpacks a plugin from a zip file into the library.
func AddPluginFromZip(libraryDir, zipFilePath string) (*Plugin, error) {
	// 1. Generate a unique ID for the new plugin
	pluginID := uuid.New().String()
	pluginDir := filepath.Join(libraryDir, pluginID)

	if err := os.MkdirAll(pluginDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("could not create directory for new plugin: %w", err)
	}

	// 2. Unzip the file
	archive, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open zip file: %w", err)
	}
	defer archive.Close()

	var pluginFiles []string
	for _, f := range archive.File {
		// Basic protection against zip slip
		if strings.Contains(f.Name, "..") {
			continue
		}

		filePath := filepath.Join(pluginDir, f.Name)
		pluginFiles = append(pluginFiles, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return nil, err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return nil, err
		}

		srcFile, err := f.Open()
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(dstFile, srcFile)

		dstFile.Close()
		srcFile.Close()

		if err != nil {
			return nil, err
		}
	}

	// 3. Create metadata
	plugin := &Plugin{
		ID:          pluginID,
		Name:        strings.TrimSuffix(filepath.Base(zipFilePath), filepath.Ext(zipFilePath)),
		Version:     "0.0.0", // Placeholder
		Author:      "Unknown", // Placeholder
		Description: "Added from " + filepath.Base(zipFilePath),
		Files:       pluginFiles,
	}

	// 4. Save metadata to plugin.json
	data, err := json.MarshalIndent(plugin, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("could not marshal plugin metadata: %w", err)
	}

	if err := os.WriteFile(filepath.Join(pluginDir, "plugin.json"), data, 0644); err != nil {
		return nil, fmt.Errorf("could not write plugin.json: %w", err)
	}

	return plugin, nil
}
