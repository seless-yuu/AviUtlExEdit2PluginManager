package profile

import (
	"aviutl-plugin-manager/pkg/library"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Profile represents a collection of enabled plugins.
type Profile struct {
	Name           string          `json:"name"`
	EnabledPlugins map[string]bool `json:"enabled_plugins"` // Key is Plugin.ID
}

// ListProfiles scans the profiles directory and returns all found profiles.
func ListProfiles(profilesDir string) ([]Profile, error) {
	profiles := make([]Profile, 0)
	entries, err := os.ReadDir(profilesDir)
	if err != nil {
		return nil, fmt.Errorf("could not read profiles directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			filePath := filepath.Join(profilesDir, entry.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				// Log error but continue
				fmt.Fprintf(os.Stderr, "Warning: could not read profile file %s: %v\n", filePath, err)
				continue
			}

			var profile Profile
			if err := json.Unmarshal(data, &profile); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not parse profile file %s: %v\n", filePath, err)
				continue
			}
			profiles = append(profiles, profile)
		}
	}

	// If no profiles are found, create and save a default one.
	if len(profiles) == 0 {
		defaultProfile := Profile{
			Name:           "Default",
			EnabledPlugins: make(map[string]bool),
		}
		if err := SaveProfile(profilesDir, defaultProfile); err != nil {
			return nil, fmt.Errorf("could not save default profile: %w", err)
		}
		profiles = append(profiles, defaultProfile)
	}

	return profiles, nil
}

// SaveProfile saves a single profile to a JSON file in the profiles directory.
func SaveProfile(profilesDir string, profile Profile) error {
	if profile.Name == "" {
		return fmt.Errorf("profile name cannot be empty")
	}

	fileName := fmt.Sprintf("%s.json", profile.Name)
	filePath := filepath.Join(profilesDir, fileName)

	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal profile to json: %w", err)
	}

	return os.WriteFile(filePath, data, 0644)
}

// ActivateProfile creates symlinks for all enabled plugins in a profile.
// It first removes any old symlinks created by this manager.
func ActivateProfile(p Profile, libraryDir, pluginDestDir string) error {
	// This is a simplistic cleanup. A more robust solution would track created symlinks.
	// For now, we assume any .auf, .aui, .auo, etc. in pluginDestDir could be a symlink to clean.
	// This is dangerous and needs a better implementation, maybe by checking if a file is a symlink
	// and if it points to our library.

	// TODO: Implement a safe cleanup mechanism. For now, we'll just create links.

	// Get all plugins from the library to map IDs to file paths
	allPlugins, err := library.ListPlugins(libraryDir)
	if err != nil {
		return fmt.Errorf("could not list library plugins to activate profile: %w", err)
	}

	pluginMap := make(map[string]library.Plugin)
	for _, plugin := range allPlugins {
		pluginMap[plugin.ID] = plugin
	}

	for pluginID, enabled := range p.EnabledPlugins {
		if !enabled {
			continue
		}

		plugin, ok := pluginMap[pluginID]
		if !ok {
			fmt.Fprintf(os.Stderr, "Warning: plugin with ID %s not found in library, skipping\n", pluginID)
			continue
		}

		pluginDir := filepath.Join(libraryDir, plugin.ID)
		for _, file := range plugin.Files {
			sourcePath := filepath.Join(pluginDir, file)
			destPath := filepath.Join(pluginDestDir, file)

			// Ensure the destination directory exists (for plugins in subfolders like /plugins)
			if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
				return fmt.Errorf("could not create destination directory %s: %w", filepath.Dir(destPath), err)
			}

			// Remove existing file/symlink at destination to avoid errors
			if _, err := os.Lstat(destPath); err == nil {
				if err := os.Remove(destPath); err != nil {
					return fmt.Errorf("could not remove existing file at %s: %w", destPath, err)
				}
			}

			// Create the symlink
			if err := os.Symlink(sourcePath, destPath); err != nil {
				return fmt.Errorf("could not create symlink for %s: %w", file, err)
			}
			fmt.Printf("Created symlink: %s -> %s\n", destPath, sourcePath)
		}
	}

	return nil
}
