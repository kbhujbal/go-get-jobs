package sites

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
)

type MicrosoftJob struct {
	JobID       string `json:"jobId"`
	Title       string `json:"title"`
	PostingDate string `json:"postingDate"`
	Properties  struct {
		Locations       []string `json:"locations"`
		PrimaryLocation string   `json:"primaryLocation"`
		EmploymentType  string   `json:"employmentType"`
	} `json:"properties"`
}

type MicrosoftJobResponse struct {
	OperationResult struct {
		Result struct {
			TotalJobs int            `json:"totalJobs"`
			Jobs      []MicrosoftJob `json:"jobs"`
		} `json:"result"`
	} `json:"operationResult"`
}

func GetMicrosoftJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Microsoft")
	count := 1
	jobsMicrosoft, count, err := microsoftJobs(count)
	if err != nil {
		fmt.Println("error processing microsoft jobs: ", err)
		return jobsMicrosoft, err
	}

	for i := 2; i <= count; i++ {
		job, _, err := microsoftJobs(i)
		if err != nil {
			fmt.Println("error processing microsoft jobs: ", err.Error())
			continue
		}

		jobsMicrosoft = append(jobsMicrosoft, job...)
	}

	return jobsMicrosoft, nil
}

func microsoftJobs(page int) ([]common.JobPosting, int, error) {
	client := common.GetClient()

	url := formatMicrosoftURL(strconv.Itoa(page), "https://gcsservices.careers.microsoft.com/search/api/v1/search")

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jobsResponseMicrosoft MicrosoftJobResponse
	err = json.Unmarshal(resp.Body(), &jobsResponseMicrosoft)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing response: %v", err)
	}

	totalJobs := float64(jobsResponseMicrosoft.OperationResult.Result.TotalJobs)
	jobsPerPage := 20.0

	page = int(math.Ceil(totalJobs / jobsPerPage))

	var jobPostings []common.JobPosting
	for _, job := range jobsResponseMicrosoft.OperationResult.Result.Jobs {
		jobPosting := common.JobPosting{
			JobId:        common.Microsoft + ":" + job.JobID,
			JobTitle:     job.Title,
			Location:     formatLocations(job.Properties.Locations),
			PostedOn:     job.PostingDate,
			ExternalPath: generateMicrosoftJobLink(job.JobID, job.Title),
			Company:      "Microsoft",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, page, nil
}

func formatLocations(locations []string) string {
	if len(locations) == 0 {
		return "Unknown"
	}

	location := strings.Join(locations, "; ")
	return location
}

// generateMicrosoftJobLink dynamically creates the job link using job ID and title
func generateMicrosoftJobLink(jobID, jobTitle string) string {
	baseURL := "https://jobs.careers.microsoft.com/global/en/job"
	encodedTitle := url.PathEscape(strings.ReplaceAll(jobTitle, " ", "-"))
	return fmt.Sprintf("%s/%s/%s", baseURL, jobID, encodedTitle)
}

func formatMicrosoftURL(page string, baseURL string) string {
	queryParams := url.Values{}
	queryParams.Set("domain", "United States")
	queryParams.Set("exp", "Students and graduates")
	queryParams.Set("l", "en_us")
	queryParams.Set("pg", page)
	queryParams.Set("pgSz", "20")
	queryParams.Set("o", "Recent")
	queryParams.Set("flt", "true")

	return baseURL + "?" + queryParams.Encode()
}
