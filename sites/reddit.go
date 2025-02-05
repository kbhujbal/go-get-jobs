package sites

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/neyaadeez/go-get-jobs/common"
)

type RedditMain struct {
	JobPosts RedditJobPosts `json:"jobPosts"`
	Board    RedditBoard    `json:"board"`
	// Departments             []DepartmentElement `json:"departments"`
	// RecentlyLiveDepartments []DepartmentElement `json:"recentlyLiveDepartments"`
	// Offices                 []DepartmentElement `json:"offices"`
	// DepartmentIDS           []interface{}       `json:"departmentIds"`
	// OfficeIDS               []string            `json:"officeIds"`
	// CustomFieldFilters      CustomFieldFilters  `json:"customFieldFilters"`
	// URLToken                string              `json:"urlToken"`
	// CustomFields            []interface{}       `json:"customFields"`
}

type RedditBoard struct {
	Name       string      `json:"name"`
	PublicURL  string      `json:"public_url"`
	Content    string      `json:"content"`
	RedirectTo interface{} `json:"redirect_to"`
}

// type CustomFieldFilters struct {
// }

// type DepartmentElement struct {
// 	ID       int64               `json:"id"`
// 	Value    int64               `json:"value"`
// 	Name     string              `json:"name"`
// 	Label    string              `json:"label"`
// 	Children []DepartmentElement `json:"children,omitempty"`
// }

type RedditJobPosts struct {
	Count      int64         `json:"count"`
	Page       int64         `json:"page"`
	Total      int64         `json:"total"`
	TotalPages int64         `json:"total_pages"`
	Data       []RedditDatum `json:"data"`
}

type RedditDatum struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	InternalJobID int64   `json:"internal_job_id"`
	UpdatedAt     string  `json:"updated_at"`
	RequisitionID *string `json:"requisition_id"`
	Location      string  `json:"location"`
	AbsoluteURL   string  `json:"absolute_url"`
	PublishedAt   string  `json:"published_at"`
	// Content       string          `json:"content"`
	// Department    DatumDepartment `json:"department"`
}

// type DatumDepartment struct {
// 	Name string `json:"name"`
// 	ID   int64  `json:"id"`
// 	Path []Path `json:"path"`
// }

// type Path string
// const (
// 	Engineering Path = "Engineering"
// 	FinanceAccounting Path = "Finance & Accounting"
// 	Product Path = "Product"
// )

func GetRedditJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Reddit")
	page := 0
	allRedditJobs, count, err := redditJobs(page)
	if err != nil {
		fmt.Println("error processing Reddit jobs: ", err)
		return allRedditJobs, err
	}

	for page != count {
		page += 1
		job, _, err := redditJobs(page)
		if err != nil {
			fmt.Println("error processing Reddit jobs: ", err.Error())
			continue
		}

		allRedditJobs = append(allRedditJobs, job...)
	}

	return allRedditJobs, nil
}

func redditJobs(page int) ([]common.JobPosting, int, error) {
	client := common.GetClient()

	url := fmt.Sprintf("https://job-boards.greenhouse.io/reddit?offices%%5B%%5D=10769&offices%%5B%%5D=10168&offices%%5B%%5D=88237&offices%%5B%%5D=10167&offices%%5B%%5D=48028&page=%d&_data=routes%%2F%%24url_token", page)

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating API request(Reddit jobs): %v", err)
	}

	var redditJobs RedditMain
	err = json.Unmarshal(resp.Body(), &redditJobs)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing json (Reddit jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range redditJobs.JobPosts.Data {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "Reddit",
			JobId:        common.Reddit + ":" + strconv.Itoa(int(job.ID)),
			JobTitle:     job.Title,
			Location:     job.Location,
			PostedOn:     job.PublishedAt,
			ExternalPath: job.AbsoluteURL,
		})
	}

	return jobPostings, int(redditJobs.JobPosts.TotalPages), nil
}
