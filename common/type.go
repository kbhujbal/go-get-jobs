package common

// Job represents the structure to hold job listing data
type JobPosting struct {
	Company      string `json:"company,omitempty"`
	JobId        string `json:"jobId"`
	JobTitle     string `json:"title"`
	Location     string `json:"locationsText,omitempty"`
	PostedOn     string `json:"postedOn,omitempty"`
	ExternalPath string `json:"externalPath"`
}

// JobsResponse represents the structure of the full response
type JobsResponse struct {
	JobPostings []JobPosting `json:"jobPostings"`
	Total       int          `json:"total"`
}

// WorkdayPayload
type WorkdayPayload struct {
	Company string
	CmpCode string
	PreURL  string
	JobsURL string
	PayLoad string
}
