package sites

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
)

type OracleJob struct {
	Id                 string `json:"Id"`
	Title              string `json:"Title"`
	PostedDate         string `json:"PostedDate"`
	PrimaryLocation    string `json:"PrimaryLocation"`
	ShortDescription   string `json:"ShortDescriptionStr"`
	SecondaryLocations []struct {
		Name string `json:"Name"`
	} `json:"secondaryLocations"`
}

type OracleJobResponse struct {
	Items []struct {
		TotalJobs       int         `json:"TotalJobsCount"`
		RequisitionList []OracleJob `json:"requisitionList"`
	} `json:"items"`
}

func GetOracleJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Oracle")
	count := 0
	jobsOracle, count, err := oracleJobs(count)
	if err != nil {
		fmt.Println("error processing oracle jobs: ", err)
		return jobsOracle, err
	}

	offset := 14
	for i := 2; i <= count; i++ {
		job, _, err := oracleJobs(offset)
		if err != nil {
			fmt.Println("error processing oracle jobs: ", err.Error())
			continue
		}

		jobsOracle = append(jobsOracle, job...)
		offset += 14
	}

	return jobsOracle, nil
}

func oracleJobs(offset int) ([]common.JobPosting, int, error) {
	client := common.GetClient()

	url := formatOracleURL("https://eeho.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions", offset)

	// url := "https://eeho.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions?onlyData=true&expand=requisitionList.secondaryLocations,flexFieldsFacet.values,requisitionList.requisitionFlexFields&finder=findReqs;siteNumber=CX_45001,facetsList=LOCATIONS%3BWORK_LOCATIONS%3BWORKPLACE_TYPES%3BTITLES%3BCATEGORIES%3BORGANIZATIONS%3BPOSTING_DATES%3BFLEX_FIELDS,limit=50,lastSelectedFacet=POSTING_DATES,locationId=300000000149325,selectedCategoriesFacet=300000001917356%3B300000001917346,selectedLocationsFacet=300000000149325,selectedPostingDatesFacet=7,sortBy=POSTING_DATES_DESC"
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jobsResponseOracle OracleJobResponse
	err = json.Unmarshal(resp.Body(), &jobsResponseOracle)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing response: %v", err)
	}

	totalJobs := float64(jobsResponseOracle.Items[0].TotalJobs)
	jobsPerPage := 14.0
	offset = int(math.Ceil(totalJobs / jobsPerPage))

	var jobPostings []common.JobPosting
	for _, job := range jobsResponseOracle.Items[0].RequisitionList {
		jobPosting := common.JobPosting{
			JobId:        common.Oracle + ":" + job.Id,
			JobTitle:     job.Title,
			Location:     formatOracleLocations(job.PrimaryLocation, job.SecondaryLocations),
			PostedOn:     job.PostedDate,
			ExternalPath: "https://careers.oracle.com/jobs/#en/sites/jobsearch/job/" + job.Id,
			Company:      "Oracle",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, offset, nil
}

func formatOracleLocations(primary string, secondary []struct {
	Name string `json:"Name"`
}) string {
	var locations []string
	locations = append(locations, primary)

	for _, loc := range secondary {
		locations = append(locations, loc.Name)
	}

	return strings.Join(locations, "; ")
}

func formatOracleURL(baseURL string, limit int) string {
	queryParams := url.Values{}
	queryParams.Set("onlyData", "true")
	queryParams.Set("expand", "requisitionList.secondaryLocations,flexFieldsFacet.values,requisitionList.requisitionFlexFields")
	queryParams.Set("finder", fmt.Sprintf("findReqs;siteNumber=CX_45001,facetsList=LOCATIONS,WORK_LOCATIONS,WORKPLACE_TYPES,TITLES,CATEGORIES,ORGANIZATIONS,POSTING_DATES,FLEX_FIELDS,limit=14,lastSelectedFacet=POSTING_DATES,locationId=300000000149325,selectedCategoriesFacet=300000001917356,300000001917346,selectedLocationsFacet=300000000149325,selectedPostingDatesFacet=7,sortBy=POSTING_DATES_DESC,offset=%d", limit))

	return baseURL + "?" + queryParams.Encode()
}
