package main

import (
	"aviutl-plugin-manager/pkg/issues"
	"aviutl-plugin-manager/pkg/library"
	"aviutl-plugin-manager/pkg/profile"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	libraryDir  string
	profilesDir string
	issuesDir   string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Setup config directories
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("could not get user home directory: %v", err))
	}
	configDir := filepath.Join(homeDir, ".config", "aviutl-plugin-manager")

	// Setup library directory
	a.libraryDir = filepath.Join(configDir, "library")
	err = os.MkdirAll(a.libraryDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("could not create library directory: %v", err))
	}

	// Setup profiles directory
	a.profilesDir = filepath.Join(configDir, "profiles")
	err = os.MkdirAll(a.profilesDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("could not create profiles directory: %v", err))
	}

	// Setup issues directory
	a.issuesDir = filepath.Join(configDir, "issues")
	err = os.MkdirAll(a.issuesDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("could not create issues directory: %v", err))
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetPlugins returns the list of installed plugins.
func (a *App) GetPlugins() ([]library.Plugin, error) {
	return library.ListPlugins(a.libraryDir)
}

// GetProfiles returns the list of available profiles.
func (a *App) GetProfiles() ([]profile.Profile, error) {
	return profile.ListProfiles(a.profilesDir)
}

// SaveProfile saves a profile.
func (a *App) SaveProfile(p profile.Profile) error {
	return profile.SaveProfile(a.profilesDir, p)
}

// ActivateProfile activates the given profile.
func (a *App) ActivateProfile(profileName string, pluginInstallPath string) error {
	profiles, err := a.GetProfiles()
	if err != nil {
		return fmt.Errorf("could not list profiles to activate: %w", err)
	}
	var p *profile.Profile
	for i := range profiles {
		if profiles[i].Name == profileName {
			p = &profiles[i]
			break
		}
	}
	if p == nil {
		return fmt.Errorf("profile '%s' not found", profileName)
	}
	return profile.ActivateProfile(*p, a.libraryDir, pluginInstallPath)
}

// AddPluginFromZip adds a new plugin from a zip file.
func (a *App) AddPluginFromZip(zipPath string) (*library.Plugin, error) {
	return library.AddPluginFromZip(a.libraryDir, zipPath)
}

// LaunchAviUtl starts the AviUtl executable.
func (a *App) LaunchAviUtl(aviutlDir string) error {
	aviutlExePath := filepath.Join(aviutlDir, "aviutl2.exe")
	if _, err := os.Stat(aviutlExePath); os.IsNotExist(err) {
		return fmt.Errorf("aviutl2.exe not found at %s", aviutlExePath)
	}

	cmd := exec.Command(aviutlExePath)
	cmd.Dir = aviutlDir // Set the working directory to the AviUtl folder
	return cmd.Start()
}

// SelectFile prompts the user to select a file and returns the path.
func (a *App) SelectFile() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Plugin Zip File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Zip Archives (*.zip)",
				Pattern:     "*.zip",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}

// GetIssues returns the list of all issues.
func (a *App) GetIssues() ([]issues.Issue, error) {
	return issues.ListIssues(a.issuesDir)
}

// CreateIssue creates a new issue.
func (a *App) CreateIssue(title, description, category, priority string) (*issues.Issue, error) {
	return issues.CreateIssue(a.issuesDir, title, description, category, priority)
}

// UpdateIssueStatus updates the status of an existing issue.
func (a *App) UpdateIssueStatus(issueID, newStatus string) error {
	return issues.UpdateIssueStatus(a.issuesDir, issueID, newStatus)
}

// GetIssue retrieves a specific issue by ID.
func (a *App) GetIssue(issueID string) (*issues.Issue, error) {
	return issues.GetIssue(a.issuesDir, issueID)
}

// DeleteIssue removes an issue.
func (a *App) DeleteIssue(issueID string) error {
	return issues.DeleteIssue(a.issuesDir, issueID)
}
