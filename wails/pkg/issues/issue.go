package issues

import "time"

// Issue represents a problem report or feature request.
type Issue struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "open", "closed", "in_progress"
	Priority    string    `json:"priority"` // "low", "medium", "high", "critical"
	Category    string    `json:"category"` // "bug", "feature", "enhancement", "question"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IssueStatus constants for issue status
const (
	StatusOpen       = "open"
	StatusClosed     = "closed"
	StatusInProgress = "in_progress"
)

// IssuePriority constants for issue priority
const (
	PriorityLow      = "low"
	PriorityMedium   = "medium"
	PriorityHigh     = "high"
	PriorityCritical = "critical"
)

// IssueCategory constants for issue category
const (
	CategoryBug         = "bug"
	CategoryFeature     = "feature"
	CategoryEnhancement = "enhancement"
	CategoryQuestion    = "question"
)