package sites

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/neyaadeez/go-get-jobs/common"
)

type AkunaMain struct {
	MatchedJobs      []AkunaMatchedJob `json:"matched_jobs"`
	SpecialtyFilters []string          `json:"specialty_filters"`
}

type AkunaMatchedJob struct {
	Location    []string `json:"location"`
	Experience  string   `json:"experience"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	ID          int64    `json:"id"`
	Specialties []string `json:"specialties"`
	Department  []string `json:"department"`
}

func GetAkunaCapitalJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Akuna Capital")
	client := common.GetClient()

	url := "https://akunacapital.com/wp-admin/admin-ajax.php?action=gh_ajax_request&experience=&department=Development&location=Chicago&search_term="

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(akuna capital jobs): %v", err)
	}

	var akunaJobs AkunaMain
	err = json.Unmarshal(resp.Body(), &akunaJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (akuna capital jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range akunaJobs.MatchedJobs {
		location := "USA"
		if len(job.Location) > 1 {
			location = job.Location[0] + " " + location
		}

		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "Akuna Capital",
			JobId:        common.AkunaCapital + ":" + strconv.Itoa(int(job.ID)),
			JobTitle:     job.Name,
			Location:     location,
			ExternalPath: "https://akunacapital.com/job-details?gh_jid=" + strconv.Itoa(int(job.ID)),
		})
	}

	return jobPostings, nil
}
