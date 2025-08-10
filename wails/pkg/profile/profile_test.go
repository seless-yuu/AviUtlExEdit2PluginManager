package profile

import (
	"aviutl-plugin-manager/pkg/library"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveAndListProfiles(t *testing.T) {
	profilesDir, err := os.MkdirTemp("", "test-profiles")
	require.NoError(t, err)
	defer os.RemoveAll(profilesDir)

	// Test listing with no profiles (should create a default)
	profiles, err := ListProfiles(profilesDir)
	require.NoError(t, err)
	require.Len(t, profiles, 1)
	assert.Equal(t, "Default", profiles[0].Name)

	// Test saving a new profile
	newProfile := Profile{
		Name: "Editing",
		EnabledPlugins: map[string]bool{
			"plugin-id-1": true,
		},
	}
	err = SaveProfile(profilesDir, newProfile)
	require.NoError(t, err)

	// Test listing again
	profiles, err = ListProfiles(profilesDir)
	require.NoError(t, err)
	require.Len(t, profiles, 2)

	// Verify the content of the saved profile file
	filePath := filepath.Join(profilesDir, "Editing.json")
	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	var loadedProfile Profile
	err = json.Unmarshal(data, &loadedProfile)
	require.NoError(t, err)
	assert.Equal(t, "Editing", loadedProfile.Name)
	assert.True(t, loadedProfile.EnabledPlugins["plugin-id-1"])
}

func TestActivateProfile(t *testing.T) {
	// 1. Setup directories
	libraryDir, err := os.MkdirTemp("", "test-lib-for-profile")
	require.NoError(t, err)
	defer os.RemoveAll(libraryDir)

	pluginDestDir, err := os.MkdirTemp("", "test-plugindest")
	require.NoError(t, err)
	defer os.RemoveAll(pluginDestDir)

	// 2. Create a fake plugin in the library
	fakePlugin := library.Plugin{
		ID:    "fake-plugin-123",
		Name:  "Fake Plugin",
		Files: []string{"fake.auf", filepath.Join("sub", "fake.aui")},
	}
	pluginDir := filepath.Join(libraryDir, fakePlugin.ID)
	require.NoError(t, os.MkdirAll(filepath.Join(pluginDir, "sub"), os.ModePerm))
	// Create dummy files
	_, err = os.Create(filepath.Join(pluginDir, "fake.auf"))
	require.NoError(t, err)
	_, err = os.Create(filepath.Join(pluginDir, "sub", "fake.aui"))
	require.NoError(t, err)
	// Create plugin.json so ListPlugins can find it
	data, err := json.Marshal(fakePlugin)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(pluginDir, "plugin.json"), data, 0644))

	// 3. Define a profile that enables this plugin
	p := Profile{
		Name: "TestActivation",
		EnabledPlugins: map[string]bool{
			"fake-plugin-123": true,
			"disabled-plugin": false, // This one should be ignored
		},
	}

	// 4. Activate the profile
	err = ActivateProfile(p, libraryDir, pluginDestDir)
	require.NoError(t, err)

	// 5. Verify the symlinks
	// Check fake.auf
	destAuf := filepath.Join(pluginDestDir, "fake.auf")
	assert.FileExists(t, destAuf)
	link, err := os.Readlink(destAuf)
	require.NoError(t, err)
	assert.Equal(t, filepath.Join(pluginDir, "fake.auf"), link)

	// Check sub/fake.aui
	destAui := filepath.Join(pluginDestDir, "sub", "fake.aui")
	assert.FileExists(t, destAui)
	link, err = os.Readlink(destAui)
	require.NoError(t, err)
	assert.Equal(t, filepath.Join(pluginDir, "sub", "fake.aui"), link)
}
