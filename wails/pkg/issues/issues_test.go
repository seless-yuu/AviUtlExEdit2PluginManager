package issues

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndListIssues(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Test listing with no issues
	issues, err := ListIssues(issuesDir)
	require.NoError(t, err)
	assert.Empty(t, issues)

	// Test creating a new issue
	issue, err := CreateIssue(issuesDir, "Test Bug", "This is a test bug description", CategoryBug, PriorityHigh)
	require.NoError(t, err)
	assert.NotEmpty(t, issue.ID)
	assert.Equal(t, "Test Bug", issue.Title)
	assert.Equal(t, "This is a test bug description", issue.Description)
	assert.Equal(t, CategoryBug, issue.Category)
	assert.Equal(t, PriorityHigh, issue.Priority)
	assert.Equal(t, StatusOpen, issue.Status)
	assert.True(t, time.Since(issue.CreatedAt) < time.Minute)
	assert.True(t, time.Since(issue.UpdatedAt) < time.Minute)

	// Test listing with one issue
	issues, err = ListIssues(issuesDir)
	require.NoError(t, err)
	assert.Len(t, issues, 1)
	assert.Equal(t, issue.ID, issues[0].ID)

	// Test creating another issue
	issue2, err := CreateIssue(issuesDir, "Feature Request", "Need new feature", CategoryFeature, PriorityMedium)
	require.NoError(t, err)
	assert.NotEmpty(t, issue2.ID)
	assert.NotEqual(t, issue.ID, issue2.ID)

	// Test listing with two issues
	issues, err = ListIssues(issuesDir)
	require.NoError(t, err)
	assert.Len(t, issues, 2)
}

func TestCreateIssueValidation(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Test empty title
	_, err = CreateIssue(issuesDir, "", "description", CategoryBug, PriorityHigh)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title cannot be empty")

	// Test invalid category defaults to bug
	issue, err := CreateIssue(issuesDir, "Test", "description", "invalid_category", PriorityHigh)
	require.NoError(t, err)
	assert.Equal(t, CategoryBug, issue.Category)

	// Test invalid priority defaults to medium
	issue, err = CreateIssue(issuesDir, "Test2", "description", CategoryFeature, "invalid_priority")
	require.NoError(t, err)
	assert.Equal(t, PriorityMedium, issue.Priority)
}

func TestGetIssue(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Create an issue
	createdIssue, err := CreateIssue(issuesDir, "Test Issue", "Test description", CategoryBug, PriorityHigh)
	require.NoError(t, err)

	// Get the issue
	retrievedIssue, err := GetIssue(issuesDir, createdIssue.ID)
	require.NoError(t, err)
	assert.Equal(t, createdIssue.ID, retrievedIssue.ID)
	assert.Equal(t, createdIssue.Title, retrievedIssue.Title)
	assert.Equal(t, createdIssue.Description, retrievedIssue.Description)

	// Test getting non-existent issue
	_, err = GetIssue(issuesDir, "non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUpdateIssueStatus(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Create an issue
	issue, err := CreateIssue(issuesDir, "Test Issue", "Test description", CategoryBug, PriorityHigh)
	require.NoError(t, err)
	assert.Equal(t, StatusOpen, issue.Status)

	// Update status to in progress
	err = UpdateIssueStatus(issuesDir, issue.ID, StatusInProgress)
	require.NoError(t, err)

	// Verify status was updated
	updatedIssue, err := GetIssue(issuesDir, issue.ID)
	require.NoError(t, err)
	assert.Equal(t, StatusInProgress, updatedIssue.Status)
	assert.True(t, updatedIssue.UpdatedAt.After(issue.UpdatedAt))

	// Update status to closed
	err = UpdateIssueStatus(issuesDir, issue.ID, StatusClosed)
	require.NoError(t, err)

	// Verify status was updated
	updatedIssue, err = GetIssue(issuesDir, issue.ID)
	require.NoError(t, err)
	assert.Equal(t, StatusClosed, updatedIssue.Status)

	// Test invalid status
	err = UpdateIssueStatus(issuesDir, issue.ID, "invalid_status")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid status")

	// Test updating non-existent issue
	err = UpdateIssueStatus(issuesDir, "non-existent-id", StatusClosed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestDeleteIssue(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Create an issue
	issue, err := CreateIssue(issuesDir, "Test Issue", "Test description", CategoryBug, PriorityHigh)
	require.NoError(t, err)

	// Verify issue exists
	_, err = GetIssue(issuesDir, issue.ID)
	require.NoError(t, err)

	// Delete the issue
	err = DeleteIssue(issuesDir, issue.ID)
	require.NoError(t, err)

	// Verify issue no longer exists
	_, err = GetIssue(issuesDir, issue.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// Test deleting non-existent issue
	err = DeleteIssue(issuesDir, "non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestSaveIssue(t *testing.T) {
	issuesDir, err := os.MkdirTemp("", "test-issues")
	require.NoError(t, err)
	defer os.RemoveAll(issuesDir)

	// Test saving issue with empty ID
	issue := Issue{
		Title:       "Test",
		Description: "Test description",
		Status:      StatusOpen,
		Category:    CategoryBug,
		Priority:    PriorityMedium,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = SaveIssue(issuesDir, issue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ID cannot be empty")

	// Test saving valid issue
	issue.ID = "test-id"
	err = SaveIssue(issuesDir, issue)
	require.NoError(t, err)

	// Verify file was created
	filePath := filepath.Join(issuesDir, "test-id.json")
	_, err = os.Stat(filePath)
	assert.NoError(t, err)
}

func TestListIssuesWithNonExistentDirectory(t *testing.T) {
	// Test listing issues from non-existent directory
	issues, err := ListIssues("/path/that/does/not/exist")
	require.NoError(t, err)
	assert.Empty(t, issues)
}

func TestIssueConstants(t *testing.T) {
	// Test status constants
	assert.Equal(t, "open", StatusOpen)
	assert.Equal(t, "closed", StatusClosed)
	assert.Equal(t, "in_progress", StatusInProgress)

	// Test priority constants
	assert.Equal(t, "low", PriorityLow)
	assert.Equal(t, "medium", PriorityMedium)
	assert.Equal(t, "high", PriorityHigh)
	assert.Equal(t, "critical", PriorityCritical)

	// Test category constants
	assert.Equal(t, "bug", CategoryBug)
	assert.Equal(t, "feature", CategoryFeature)
	assert.Equal(t, "enhancement", CategoryEnhancement)
	assert.Equal(t, "question", CategoryQuestion)
}