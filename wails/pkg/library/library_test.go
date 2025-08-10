package library

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createDummyZip is a helper function to create a zip file for testing.
func createDummyZip(t *testing.T, dir, zipName string, files map[string]string) string {
	zipPath := filepath.Join(dir, zipName)
	zipFile, err := os.Create(zipPath)
	require.NoError(t, err)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for name, content := range files {
		f, err := zipWriter.Create(name)
		require.NoError(t, err)
		_, err = f.Write([]byte(content))
		require.NoError(t, err)
	}
	return zipPath
}

func TestAddAndListPlugins(t *testing.T) {
	// 1. Setup a temporary directory for the library
	libraryDir, err := os.MkdirTemp("", "test-library")
	require.NoError(t, err)
	defer os.RemoveAll(libraryDir)

	// Setup a temporary directory for the zip file
	zipDir, err := os.MkdirTemp("", "test-zip")
	require.NoError(t, err)
	defer os.RemoveAll(zipDir)

	// 2. Create a dummy zip file
	dummyFiles := map[string]string{
		"test_plugin.auf": "this is a fake auf file",
		"readme.txt":      "hello world",
	}
	zipPath := createDummyZip(t, zipDir, "my_plugin.zip", dummyFiles)

	// 3. Test AddPluginFromZip
	addedPlugin, err := AddPluginFromZip(libraryDir, zipPath)
	require.NoError(t, err)
	require.NotNil(t, addedPlugin)

	// 4. Assert properties of the added plugin
	assert.Equal(t, "my_plugin", addedPlugin.Name) // Name is derived from zip filename
	assert.NotEmpty(t, addedPlugin.ID)
	assert.ElementsMatch(t, []string{"test_plugin.auf", "readme.txt"}, addedPlugin.Files)

	// Verify that the plugin directory and files were created
	pluginDir := filepath.Join(libraryDir, addedPlugin.ID)
	assert.DirExists(t, pluginDir)

	// Verify plugin.json was created
	pluginJSONPath := filepath.Join(pluginDir, "plugin.json")
	assert.FileExists(t, pluginJSONPath)

	// Verify files were extracted
	extractedAufPath := filepath.Join(pluginDir, "test_plugin.auf")
	assert.FileExists(t, extractedAufPath)
	content, err := os.ReadFile(extractedAufPath)
	require.NoError(t, err)
	assert.Equal(t, "this is a fake auf file", string(content))

	// 5. Test ListPlugins
	listedPlugins, err := ListPlugins(libraryDir)
	require.NoError(t, err)
	require.Len(t, listedPlugins, 1)

	// 6. Assert properties of the listed plugin
	listedPlugin := listedPlugins[0]
	assert.Equal(t, addedPlugin.ID, listedPlugin.ID)
	assert.Equal(t, "my_plugin", listedPlugin.Name)
	assert.ElementsMatch(t, []string{"test_plugin.auf", "readme.txt"}, listedPlugin.Files)
}

func TestListPlugins_Empty(t *testing.T) {
	libraryDir, err := os.MkdirTemp("", "test-library-empty")
	require.NoError(t, err)
	defer os.RemoveAll(libraryDir)

	plugins, err := ListPlugins(libraryDir)
	require.NoError(t, err)
	assert.Len(t, plugins, 0)
}
