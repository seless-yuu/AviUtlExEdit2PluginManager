package main

import (
	"aviutl-plugin-manager/pkg/issues"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppIssueManagement(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-app")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create an App instance
	app := NewApp()
	app.ctx = context.Background()
	app.issuesDir = filepath.Join(tempDir, "issues")

	// Test creating an issue
	issue, err := app.CreateIssue("Test Bug Report", "This is a test bug", issues.CategoryBug, issues.PriorityHigh)
	require.NoError(t, err)
	assert.NotEmpty(t, issue.ID)
	assert.Equal(t, "Test Bug Report", issue.Title)
	assert.Equal(t, issues.StatusOpen, issue.Status)

	// Test getting all issues
	allIssues, err := app.GetIssues()
	require.NoError(t, err)
	assert.Len(t, allIssues, 1)
	assert.Equal(t, issue.ID, allIssues[0].ID)

	// Test getting a specific issue
	retrievedIssue, err := app.GetIssue(issue.ID)
	require.NoError(t, err)
	assert.Equal(t, issue.ID, retrievedIssue.ID)
	assert.Equal(t, issue.Title, retrievedIssue.Title)

	// Test updating issue status
	err = app.UpdateIssueStatus(issue.ID, issues.StatusInProgress)
	require.NoError(t, err)

	// Verify status was updated
	updatedIssue, err := app.GetIssue(issue.ID)
	require.NoError(t, err)
	assert.Equal(t, issues.StatusInProgress, updatedIssue.Status)

	// Test creating another issue
	featureIssue, err := app.CreateIssue("Feature Request", "New feature needed", issues.CategoryFeature, issues.PriorityMedium)
	require.NoError(t, err)

	// Test getting all issues again
	allIssues, err = app.GetIssues()
	require.NoError(t, err)
	assert.Len(t, allIssues, 2)

	// Test deleting an issue
	err = app.DeleteIssue(featureIssue.ID)
	require.NoError(t, err)

	// Verify issue was deleted
	allIssues, err = app.GetIssues()
	require.NoError(t, err)
	assert.Len(t, allIssues, 1)
	assert.Equal(t, issue.ID, allIssues[0].ID)
}

func TestAppIssueErrorHandling(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-app-errors")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	app := NewApp()
	app.ctx = context.Background()
	app.issuesDir = filepath.Join(tempDir, "issues")

	// Test creating issue with empty title
	_, err = app.CreateIssue("", "description", issues.CategoryBug, issues.PriorityHigh)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title cannot be empty")

	// Test getting non-existent issue
	_, err = app.GetIssue("non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// Test updating status of non-existent issue
	err = app.UpdateIssueStatus("non-existent-id", issues.StatusClosed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// Test deleting non-existent issue
	err = app.DeleteIssue("non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}