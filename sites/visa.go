package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type VisaMain struct {
	Successful         bool            `json:"successful"`
	ErrorMessage       interface{}     `json:"errorMessage"`
	ResponseTimeMillis int64           `json:"responseTimeMillis"`
	TotalRecords       int64           `json:"totalRecords"`
	RecordsMatched     int64           `json:"recordsMatched"`
	PageSize           int64           `json:"pageSize"`
	Page               int64           `json:"page"`
	From               int64           `json:"from"`
	JobDetails         []VisaJobDetail `json:"jobDetails"`
}

type VisaJobDetail struct {
	RefNumber             string      `json:"refNumber"`
	PostingID             string      `json:"postingId"`
	JobTitle              string      `json:"jobTitle"`
	CompanyDescription    interface{} `json:"companyDescription"`
	JobDescription        string      `json:"jobDescription"`
	Qualifications        string      `json:"qualifications"`
	AdditionalInformation string      `json:"additionalInformation"`
	Function              string      `json:"function"`
	TypeOfEmployment      string      `json:"typeOfEmployment"`
	City                  string      `json:"city"`
	Region                string      `json:"region"`
	Country               string      `json:"country"`
	CountryCode           string      `json:"countryCode"`
	CreatedOn             string      `json:"createdOn"`
	Department            string      `json:"department"`
	SuperDepartment       string      `json:"superDepartment"`
	ApplyURL              string      `json:"applyUrl"`
	Recommendations       interface{} `json:"recommendations"`
}

func GetVisaJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Visa")
	var jobPostings []common.JobPosting

	url := "https://search.visa.com/CAREERS/careers/jobs?q="
	var from int64 = 0

	client := common.GetClient()

	for {
		payload := map[string]interface{}{
			"filters": []map[string]interface{}{
				{
					"department": []string{
						"Cyber Security",
						"Data Architect/Engineering/Science",
						"Data Science/Data Engineering",
						"Information Technology",
						"Intern",
						"Product & Project Management (Technical)",
						"Product Management (Technical)",
						"Risk & Security",
						"Software Development/Engineering",
						"Software Quality Assurance and Testing",
						"Technology and Operations",
					},
				},
			},
			"city": []string{
				"Toronto", "Ashburn", "Atlanta", "Austin",
				"Bellevue", "Denver", "Foster City", "Highlands Ranch",
				"Mentor", "Miami", "New York", "San Francisco",
				"Washington", "Wilmington",
			},
			"from": from,
			"size": 10,
			"sort": map[string]string{
				"createdOn": "DESC",
			},
		}

		body, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON payload: %v", err)
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetBody(body).
			Post(url)

		if err != nil {
			return nil, fmt.Errorf("error fetching Visa jobs: %v", err)
		}

		var visaResponse VisaMain
		if err := json.Unmarshal(resp.Body(), &visaResponse); err != nil {
			return nil, fmt.Errorf("error decoding Visa response: %v", err)
		}

		if !visaResponse.Successful {
			return nil, fmt.Errorf("unsuccessful response from Visa API: %v", visaResponse.ErrorMessage)
		}

		for _, job := range visaResponse.JobDetails {
			jobPosting := common.JobPosting{
				JobId:        common.Visa + ":" + job.PostingID,
				JobTitle:     job.JobTitle,
				Location:     fmt.Sprintf("%s, %s, %s", job.City, job.Region, job.Country),
				ExternalPath: job.ApplyURL,
				Company:      "Visa",
			}
			jobPostings = append(jobPostings, jobPosting)
		}

		from += 10
		if from >= visaResponse.RecordsMatched {
			break
		}
	}

	fmt.Println(jobPostings[0])
	return jobPostings, nil
}
