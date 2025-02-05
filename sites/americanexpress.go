package sites

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/neyaadeez/go-get-jobs/common"
)

type AmexMain struct {
	Positions []AmexPosition `json:"positions"`
	Count     int64          `json:"count"`
}

type AmexPosition struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	// Locations            []string    `json:"locations"`
	// Hot                  int64       `json:"hot"`
	// Department           string      `json:"department"`
	// TUpdate              int64       `json:"t_update"`
	TCreate int64 `json:"t_create"`
	// AtsJobID             string      `json:"ats_job_id"`
	// DisplayJobID         string      `json:"display_job_id"`
	// IDLocale             string      `json:"id_locale"`
	// JobDescription       string      `json:"job_description"`
	// MedallionProgram     interface{} `json:"medallionProgram"`
	// LocationFlexibility  interface{} `json:"location_flexibility"`
	CanonicalPositionURL string `json:"canonicalPositionUrl"`
	// IsPrivate            bool        `json:"isPrivate"`
}

func GetAmexJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "American Express")
	count := 1
	start := 0
	jobsAmex, count, err := amexJobs(start)
	if err != nil {
		fmt.Println("error processing amex jobs: ", err)
		return jobsAmex, err
	}

	for range count - 1 {
		start += 10
		job, _, err := amexJobs(start)
		if err != nil {
			fmt.Println("error processing amex jobs: ", err.Error())
			continue
		}

		jobsAmex = append(jobsAmex, job...)
	}

	return jobsAmex, nil
}

func amexJobs(page int) ([]common.JobPosting, int, error) {
	client := common.GetClient()

	url := fmt.Sprintf("https://aexp.eightfold.ai/api/apply/v2/jobs?domain=aexp.com&start=%d&num=10&exclude_pid=24247225&location=United%%20States&pid=24247225&Select%%20Primary%%20Career%%20Areas=technology&domain=aexp.com&sort_by=newest", page)
	//url := formatAmexURL(strconv.Itoa(page), "https://aexp.eightfold.ai/api/apply/v2/jobs")

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jobsAmexResponse AmexMain
	err = json.Unmarshal(resp.Body(), &jobsAmexResponse)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing response: %v", err)
	}

	totalJobs := float64(jobsAmexResponse.Count)
	jobsPerPage := 10.0

	page = int(math.Ceil(totalJobs / jobsPerPage))

	var jobPostings []common.JobPosting
	for _, job := range jobsAmexResponse.Positions {
		jobPosting := common.JobPosting{
			JobId:        common.Amex + ":" + strconv.Itoa(int(job.ID)),
			JobTitle:     job.Name,
			Location:     job.Location,
			PostedOn:     strconv.Itoa(int(job.TCreate)),
			ExternalPath: job.CanonicalPositionURL,
			Company:      "American Express",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, page, nil
}

// func formatAmexURL(start string, baseURL string) string {
// 	queryParams := url.Values{}
// 	queryParams.Set("lc", "aexp.com")
// 	queryParams.Set("start", start)
// 	queryParams.Set("num", "20")
// 	queryParams.Set("exclude_pid", "24247225")
// 	queryParams.Set("location", "United%20States")
// 	queryParams.Set("pid", "24247225")
// 	queryParams.Set("Select%20Primary%20Career%20Areas", "technology")
// 	queryParams.Set("sort_by", "newest")

// 	return baseURL + "?" + queryParams.Encode()
// }
