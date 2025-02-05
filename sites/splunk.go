package sites

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neyaadeez/go-get-jobs/common"
)

type SplunkMain struct {
	Length      int64               `json:"length"`
	CareersList []SplunkCareersList `json:"careersList"`
}

type SplunkCareersList struct {
	AllLocations  string            `json:"allLocations"`
	JobTitle      string            `json:"jobTitle"`
	Locations     []string          `json:"locations"`
	Category      SplunkCategory    `json:"category"`
	Region        SplunkRegion      `json:"region"`
	URL           string            `json:"url"`
	JobType       SplunkJobType     `json:"jobType"`
	RemoteType    SplunkRemoteType  `json:"remoteType"`
	RequisitionID string            `json:"requisitionId"`
	JobDivision   SplunkJobDivision `json:"jobDivision"`
}

type SplunkCategory string

const (
	EarlyTalent SplunkCategory = "Early Talent"
)

type SplunkJobDivision string

const (
	EngProdDesign SplunkJobDivision = "Eng, Prod & Design"
	Sales         SplunkJobDivision = "Sales"
)

type SplunkJobType string

const (
	EmergingTalent SplunkJobType = "Emerging Talent"
	FullTime       SplunkJobType = "Full-Time"
	Intern         SplunkJobType = "Intern"
)

type SplunkRegion string

const (
	Americas SplunkRegion = "Americas"
	Emea     SplunkRegion = "EMEA"
)

type SplunkRemoteType string

const (
	HybridRemote SplunkRemoteType = "Hybrid Remote"
	Remote       SplunkRemoteType = "Remote"
)

func GetSplunkJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Splunk")
	var jobPostings []common.JobPosting

	url := "https://www.splunk.com/api/bin/careers/joblist"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching Splunk jobs: %v", err)
	}
	defer resp.Body.Close()

	var splunkJobs SplunkMain
	if err := json.NewDecoder(resp.Body).Decode(&splunkJobs); err != nil {
		return nil, fmt.Errorf("error decoding Splunk response: %v", err)
	}

	for _, job := range splunkJobs.CareersList {
		if job.Region == Americas {
			jobPosting := common.JobPosting{
				JobId:        common.Splunk + ":" + job.RequisitionID,
				JobTitle:     job.JobTitle,
				Location:     job.AllLocations,
				ExternalPath: "https://www.splunk.com" + job.URL,
				Company:      "Splunk",
			}
			jobPostings = append(jobPostings, jobPosting)
		}
	}

	return jobPostings, nil
}
