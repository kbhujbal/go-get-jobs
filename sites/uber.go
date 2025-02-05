package sites

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/neyaadeez/go-get-jobs/common"
)

// UberMain represents the main structure of the Uber job API response.
type UberMain struct {
	Status string   `json:"status"`
	Data   UberData `json:"data"`
}

// UberData holds the results and total results of job postings.
type UberData struct {
	Results      []UberResult     `json:"results"`
	TotalResults UberTotalResults `json:"totalResults"`
}

// UberResult represents a job posting from Uber.
type UberResult struct {
	ID                 int64          `json:"id"`
	Title              string         `json:"title"`
	Description        string         `json:"description"`
	Department         string         `json:"department"`
	Type               string         `json:"type"`
	ProgramAndPlatform string         `json:"programAndPlatform"`
	Location           UberLocation   `json:"location"`
	Featured           bool           `json:"featured"`
	Level              string         `json:"level"`
	CreationDate       string         `json:"creationDate"`
	OtherLevels        interface{}    `json:"otherLevels"`
	Team               string         `json:"team"`
	PortalID           string         `json:"portalID"`
	IsPipeline         bool           `json:"isPipeline"`
	StatusID           string         `json:"statusID"`
	StatusName         string         `json:"statusName"`
	UpdatedDate        string         `json:"updatedDate"`
	UniqueSkills       string         `json:"uniqueSkills"`
	TimeType           UberTimeType   `json:"timeType"`
	AllLocations       []UberLocation `json:"allLocations"`
}

// UberLocation represents the location details of a job.
type UberLocation struct {
	Country     string `json:"country"`
	Region      string `json:"region"`
	City        string `json:"city"`
	CountryName string `json:"countryName"`
}

// UberTotalResults holds the total results count of job postings.
type UberTotalResults struct {
	Low      int64 `json:"low"`
	High     int64 `json:"high"`
	Unsigned bool  `json:"unsigned"`
}

const (
	UberCityNewYork  string = "New York"
	UberSANFrancisco string = "San Francisco"
	UberSeattle      string = "Seattle"
	UberSunnyvale    string = "Sunnyvale"
)

const (
	UberUsa string = "USA"
)

const (
	UberUnitedStates string = "United States"
)

const (
	UberCalifornia    string = "California"
	UberRegionNewYork string = "New York"
	UberWashington    string = "Washington"
)

const (
	UberDataScience string = "Data Science"
	UberEngineering string = "Engineering"
	UberUniversity  string = "University"
)

const (
	UberExternal string = "EXTERNAL"
)

const (
	UberD31001 string = "D31001"
)

const (
	UberApproved string = "Approved"
)

// UberTimeType represents the time type of the job.
type UberTimeType string

const (
	UberFullTime UberTimeType = "Full-Time"
	UberIntern   UberTimeType = "Intern"
)

// GetUberJobs fetches job postings from the Uber API.
func GetUberJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Uber")
	var jobPostings []common.JobPosting

	url := "https://www.uber.com/api/loadSearchJobsResults?localeCode=en"
	var page int64
	var totalRecords int64

	client := common.GetClient()

	for {
		payload := map[string]interface{}{
			"params": map[string]interface{}{
				"location": []map[string]string{
					{"country": "USA", "region": "California", "city": "San Francisco"},
					{"country": "USA", "region": "California", "city": "Culver City"},
					{"country": "USA", "region": "California", "city": "Sunnyvale"},
					{"country": "USA", "region": "New York", "city": "New York"},
					{"country": "USA", "region": "Illinois", "city": "Chicago"},
					{"country": "USA", "region": "Washington", "city": "Seattle"},
					{"country": "USA", "region": "Arizona", "city": "Phoenix"},
					{"country": "USA", "region": "Texas", "city": "Dallas"},
					{"country": "USA", "region": "District of Columbia", "city": "Washington"},
					{"country": "USA", "region": "Pennsylvania", "city": "Philadelphia"},
					{"country": "USA", "region": "Pennsylvania", "city": "Pittsburgh"},
					{"country": "USA", "region": "Colorado", "city": "Denver"},
					{"country": "USA", "region": "Florida", "city": "Miami"},
					{"country": "USA", "region": "Georgia", "city": "Atlanta"},
					{"country": "USA", "region": "Massachusetts", "city": "Boston"},
				},
				"department": []string{
					"Data Science",
					"Engineering",
					"University",
					"Product",
				},
				"team":               []string{},
				"programAndPlatform": []string{},
				"lineOfBusinessName": []string{},
			},
			"page":  page,
			"limit": 10,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON payload: %v", err)
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetHeader("x-csrf-token", "x").
			SetBody(body).
			Post(url)

		if err != nil {
			return nil, fmt.Errorf("error fetching Uber jobs: %v", err)
		}

		var uberResponse UberMain
		if err := json.Unmarshal(resp.Body(), &uberResponse); err != nil {
			return nil, fmt.Errorf("error decoding Visa response: %v", err)
		}

		if uberResponse.Status != "success" {
			return nil, fmt.Errorf("unsuccessful response from Uber API: %v", uberResponse.Data.Results)
		}

		totalRecords = uberResponse.Data.TotalResults.Low

		for _, job := range uberResponse.Data.Results {
			jobPosting := common.JobPosting{
				JobId:        common.Uber + ":" + fmt.Sprintf("%d", job.ID),
				JobTitle:     job.Title,
				Location:     job.Location.City + ", " + job.Location.Region + ", " + job.Location.Country,
				ExternalPath: "https://www.uber.com/global/en/careers/list/" + strconv.Itoa(int(job.ID)),
				Company:      "Uber",
			}
			jobPostings = append(jobPostings, jobPosting)
		}

		page++
		if int64(len(jobPostings)) >= totalRecords {
			break
		}
	}

	return jobPostings, nil
}
