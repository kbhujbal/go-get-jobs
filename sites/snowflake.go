package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
	commonconst "github.com/neyaadeez/go-get-jobs/common_const"
)

type SnowflakeMain struct {
	RefineSearch SnowflakeRefineSearch `json:"refineSearch"`
}

type SnowflakeRefineSearch struct {
	Status int64         `json:"status"`
	Data   SnowflakeData `json:"data"`
}

type SnowflakeData struct {
	Jobs []SnowflakeJob `json:"jobs"`
}

type SnowflakeJob struct {
	// MlSkills    []string `json:"ml_skills"`
	CompanyName string `json:"companyName"`
	// Type               Type                 `json:"type"`
	// DescriptionTeaser  string               `json:"descriptionTeaser"`
	// State              State                `json:"state"`
	// SiteType           SiteType             `json:"siteType"`
	// MultiCategory      []Category           `json:"multi_category"`
	// ReqID              string               `json:"reqId"`
	// Grade              string               `json:"grade"`
	// City               City                 `json:"city"`
	// Latitude           string               `json:"latitude"`
	// Industry           Industry             `json:"industry"`
	// MultiLocation      []Location           `json:"multi_location"`
	HiringManagerName string `json:"hiringManagerName"`
	ApplyURL          string `json:"applyUrl"`
	// MlJobParser        MlJobParser          `json:"ml_job_parser"`
	// ExternalApply      bool                 `json:"externalApply"`
	// CityState          CityState            `json:"cityState"`
	// Country            Country              `json:"country"`
	// VisibilityType     VisibilityType       `json:"visibilityType"`
	// Longitude          string               `json:"longitude"`
	JobID string `json:"jobId"`
	// Locale             Locale               `json:"locale"`
	Title string `json:"title"`
	// JobSeqNo           string               `json:"jobSeqNo"`
	PostedDate string `json:"postedDate"`
	// WorkLocation       Location             `json:"workLocation"`
	// DateCreated        string               `json:"dateCreated"`
	// CityStateCountry   CityStateCountry     `json:"cityStateCountry"`
	// Department         string               `json:"department"`
	// JobVisibility      []SiteType           `json:"jobVisibility"`
	Location string `json:"location"`
	// Category           Category             `json:"category"`
	// IsMultiLocation    bool                 `json:"isMultiLocation"`
	// MultiLocationArray []MultiLocationArray `json:"multi_location_array"`
	// IsMultiCategory    bool                 `json:"isMultiCategory"`
	// MultiCategoryArray []MultiCategoryArray `json:"multi_category_array"`
	// Badge              string               `json:"badge"`
}

func GetSnowflakeJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Snowflake")
	var jobPostings []common.JobPosting

	client := common.GetClient()

	payload := `{
  "lang": "en_us",
  "deviceType": "desktop",
  "country": "us",
  "pageName": "search-results",
  "ddoKey": "refineSearch",
  "sortBy": "Most recent",
  "subsearch": "",
  "from": 0,
  "jobs": true,
  "counts": true,
  "all_fields": [
    "category",
    "department",
    "location",
    "region",
    "jobLevel"
  ],
  "size": 1000,
  "clearAll": false,
  "jdsource": "facets",
  "isSliderEnable": false,
  "pageId": "page11",
  "siteType": "external",
  "keywords": "software",
  "global": true,
  "selected_fields": {
    "category": [
      "Engineering",
      "IT",
      "Security",
      "Product"
    ],
    "region": [
      "AMS"
    ],
    "location": [
      "Bellevue, Washington, USA",
      "McLean, Virginia, USA",
      "Dublin, California, USA",
      "San Mateo, California, USA"
    ]
  },
  "sort": {
    "order": "desc",
    "field": "postedDate"
  },
  "locationData": {}
}`
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("https://careers.snowflake.com/widgets")
	if err != nil {
		return nil, fmt.Errorf("error accessing the URL (snowflake jobs): %v", err)
	}

	var jobs SnowflakeMain
	err = json.Unmarshal(resp.Body(), &jobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing response(snowflake jobs): %v", err)
	}

	var snowflakeJobsData []SnowflakeJobData
	for _, job := range jobs.RefineSearch.Data.Jobs {
		jobPosting := common.JobPosting{
			JobId:        common.Snowflake + ":" + job.JobID,
			JobTitle:     job.Title,
			Location:     job.Location,
			ExternalPath: job.ApplyURL,
			Company:      job.CompanyName,
		}
		jobPostings = append(jobPostings, jobPosting)

		snowflakeJobsData = append(snowflakeJobsData, SnowflakeJobData{
			JobId:         job.JobID,
			JobTitle:      job.Title,
			HiringManager: job.HiringManagerName,
			PostingDate:   job.PostedDate,
			Location:      job.Location,
		})
	}

	UpdateSnowflakeJobs(snowflakeJobsData)
	return jobPostings, nil
}

// SnowflakeJobData represents the structure of each job entry in the JSON file
type SnowflakeJobData struct {
	JobId         string `json:"job_id"`
	JobTitle      string `json:"job_title"`
	HiringManager string `json:"hiring_manager"`
	PostingDate   string `json:"posting_date"`
	Location      string `json:"location"`
}

// UpdateSnowflakeJobs updates the local snowflake.json file with job details
func UpdateSnowflakeJobs(newJobs []SnowflakeJobData) error {
	filePath := commonconst.SnowflakeHiringManagersFile

	var existingJobs []SnowflakeJobData
	if _, err := os.Stat(filePath); err == nil {
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}
		if err := json.Unmarshal(fileBytes, &existingJobs); err != nil {
			return fmt.Errorf("error unmarshalling json: %v", err)
		}
	}

	jobMap := make(map[string]SnowflakeJobData)
	for _, job := range existingJobs {
		jobMap[job.JobId] = job
	}

	for _, job := range newJobs {
		if _, exists := jobMap[job.JobId]; !exists {
			jobMap[job.JobId] = job
		}
	}

	var updatedJobs []SnowflakeJobData
	for _, job := range jobMap {
		updatedJobs = append(updatedJobs, job)
	}
	sort.Slice(updatedJobs, func(i, j int) bool {
		t1, _ := time.Parse("2006-01-02T15:04:05.000-0700", updatedJobs[i].PostingDate)
		t2, _ := time.Parse("2006-01-02T15:04:05.000-0700", updatedJobs[j].PostingDate)
		return t1.After(t2)
	})

	fileData, err := json.MarshalIndent(updatedJobs, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling updated json: %v", err)
	}

	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
