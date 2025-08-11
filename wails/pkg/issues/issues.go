package issues

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ListIssues scans the issues directory and returns all found issues.
func ListIssues(issuesDir string) ([]Issue, error) {
	issues := make([]Issue, 0)
	entries, err := os.ReadDir(issuesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return issues, nil // Return empty list if issues dir doesn't exist yet
		}
		return nil, fmt.Errorf("could not read issues directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			filePath := filepath.Join(issuesDir, entry.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				// Log error but continue
				fmt.Fprintf(os.Stderr, "Warning: could not read issue file %s: %v\n", filePath, err)
				continue
			}

			var issue Issue
			if err := json.Unmarshal(data, &issue); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: could not parse issue file %s: %v\n", filePath, err)
				continue
			}
			issues = append(issues, issue)
		}
	}

	return issues, nil
}

// CreateIssue creates a new issue and saves it to the issues directory.
func CreateIssue(issuesDir string, title, description, category, priority string) (*Issue, error) {
	if title == "" {
		return nil, fmt.Errorf("issue title cannot be empty")
	}

	// Validate category
	if category != CategoryBug && category != CategoryFeature && 
	   category != CategoryEnhancement && category != CategoryQuestion {
		category = CategoryBug // Default to bug
	}

	// Validate priority
	if priority != PriorityLow && priority != PriorityMedium && 
	   priority != PriorityHigh && priority != PriorityCritical {
		priority = PriorityMedium // Default to medium
	}

	now := time.Now()
	issue := &Issue{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Status:      StatusOpen,
		Priority:    priority,
		Category:    category,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := SaveIssue(issuesDir, *issue); err != nil {
		return nil, fmt.Errorf("could not save new issue: %w", err)
	}

	return issue, nil
}

// SaveIssue saves a single issue to a JSON file in the issues directory.
func SaveIssue(issuesDir string, issue Issue) error {
	if issue.ID == "" {
		return fmt.Errorf("issue ID cannot be empty")
	}

	// Ensure the issues directory exists
	if err := os.MkdirAll(issuesDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create issues directory: %w", err)
	}

	fileName := fmt.Sprintf("%s.json", issue.ID)
	filePath := filepath.Join(issuesDir, fileName)

	data, err := json.MarshalIndent(issue, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal issue to json: %w", err)
	}

	return os.WriteFile(filePath, data, 0644)
}

// UpdateIssueStatus updates the status of an existing issue.
func UpdateIssueStatus(issuesDir, issueID, newStatus string) error {
	// Validate status
	if newStatus != StatusOpen && newStatus != StatusClosed && newStatus != StatusInProgress {
		return fmt.Errorf("invalid status: %s", newStatus)
	}

	issues, err := ListIssues(issuesDir)
	if err != nil {
		return fmt.Errorf("could not list issues: %w", err)
	}

	for i, issue := range issues {
		if issue.ID == issueID {
			issues[i].Status = newStatus
			issues[i].UpdatedAt = time.Now()
			return SaveIssue(issuesDir, issues[i])
		}
	}

	return fmt.Errorf("issue with ID %s not found", issueID)
}

// GetIssue retrieves a specific issue by ID.
func GetIssue(issuesDir, issueID string) (*Issue, error) {
	filePath := filepath.Join(issuesDir, fmt.Sprintf("%s.json", issueID))
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("issue with ID %s not found", issueID)
		}
		return nil, fmt.Errorf("could not read issue file: %w", err)
	}

	var issue Issue
	if err := json.Unmarshal(data, &issue); err != nil {
		return nil, fmt.Errorf("could not parse issue file: %w", err)
	}

	return &issue, nil
}

// DeleteIssue removes an issue file from the issues directory.
func DeleteIssue(issuesDir, issueID string) error {
	filePath := filepath.Join(issuesDir, fmt.Sprintf("%s.json", issueID))
	
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("issue with ID %s not found", issueID)
		}
		return fmt.Errorf("could not delete issue file: %w", err)
	}

	return nil
}