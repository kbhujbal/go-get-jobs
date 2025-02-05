package sites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/neyaadeez/go-get-jobs/common"
)

// IntuitMain struct for unmarshalling JSON response
type IntuitMain struct {
	Filters    string `json:"filters"`
	Results    string `json:"results"`
	HasJobs    bool   `json:"hasJobs"`
	HasContent bool   `json:"hasContent"`
}

func GetIntuitJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: Intuit")
	client := common.GetClient()

	url := "https://jobs.intuit.com/search-jobs/results?ActiveFacetID=8698560&CurrentPage=1&RecordsPerPage=1000&Distance=50&RadiusUnitType=0&Keywords=&Location=&ShowRadius=False&IsPagination=False&CustomFacetName=&FacetTerm=&FacetType=0&FacetFilters%5B0%5D.ID=68338&FacetFilters%5B0%5D.FacetType=1&FacetFilters%5B0%5D.Count=55&FacetFilters%5B0%5D.Display=Data&FacetFilters%5B0%5D.IsApplied=true&FacetFilters%5B0%5D.FieldName=&FacetFilters%5B1%5D.ID=68340&FacetFilters%5B1%5D.FacetType=1&FacetFilters%5B1%5D.Count=35&FacetFilters%5B1%5D.Display=Design+%26+User+Experience&FacetFilters%5B1%5D.IsApplied=true&FacetFilters%5B1%5D.FieldName=&FacetFilters%5B2%5D.ID=68347&FacetFilters%5B2%5D.FacetType=1&FacetFilters%5B2%5D.Count=10&FacetFilters%5B2%5D.Display=Information+Technology&FacetFilters%5B2%5D.IsApplied=true&FacetFilters%5B2%5D.FieldName=&FacetFilters%5B3%5D.ID=8698560&FacetFilters%5B3%5D.FacetType=1&FacetFilters%5B3%5D.Count=26&FacetFilters%5B3%5D.Display=Security%2C+Risk+%26+Fraud&FacetFilters%5B3%5D.IsApplied=true&FacetFilters%5B3%5D.FieldName=&FacetFilters%5B4%5D.ID=68357&FacetFilters%5B4%5D.FacetType=1&FacetFilters%5B4%5D.Count=134&FacetFilters%5B4%5D.Display=Software+Engineering&FacetFilters%5B4%5D.IsApplied=true&FacetFilters%5B4%5D.FieldName=&FacetFilters%5B5%5D.ID=6252001&FacetFilters%5B5%5D.FacetType=2&FacetFilters%5B5%5D.Count=260&FacetFilters%5B5%5D.Display=United+States&FacetFilters%5B5%5D.IsApplied=true&FacetFilters%5B5%5D.FieldName=&SearchResultsModuleName=Search+Results&SearchFiltersModuleName=Search+Filters&SortCriteria=1&SortDirection=0&SearchType=5&PostalCode=&ResultsType=0&fc=&fl=&fcf=&afc=&afl=&afcf=" // replace with your actual API endpoint
	var jobs []common.JobPosting

	resp, err := client.R().Get(url)
	if err != nil {
		return jobs, fmt.Errorf("error accessing the URL: %v", err)
	}

	// Unmarshal JSON response into Main struct
	var mainData IntuitMain
	if err := json.Unmarshal(resp.Body(), &mainData); err != nil {
		return jobs, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Check if jobs are present
	if !mainData.HasJobs {
		return jobs, nil
	}

	// Parse the HTML in Results field for job details
	var jb []common.JobPosting
	jb, err = parseJobHTML(mainData.Results)
	if err != nil {
		return jobs, err
	}

	jobs = append(jobs, jb...)

	return jobs, nil
}

// Function to parse job listings from HTML content
func parseJobHTML(htmlContent string) ([]common.JobPosting, error) {
	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewReader([]byte(htmlContent))))
	if err != nil {
		return nil, fmt.Errorf("error parsing the HTML: %v", err)
	}

	var jobs []common.JobPosting

	doc.Find("ul.search-list li").Each(func(i int, s *goquery.Selection) {
		jobId, exists := s.Attr("data-intuit-jobid")
		if !exists {
			return
		}
		jobTitle := s.Find("a.sr-item").Text()
		jobTitleField := strings.Split(jobTitle, "\n")
		if len(jobTitleField) > 1 {
			jobTitle = strings.TrimSpace(jobTitleField[1])
		}
		jobLocation, _ := s.Attr("data-orig-location")

		jobs = append(jobs, common.JobPosting{
			JobId:        common.Intuit + ":" + jobId,
			JobTitle:     strings.TrimSpace(jobTitle),
			ExternalPath: fmt.Sprintf("https://jobs.intuit.com%s", s.Find("a.sr-item").AttrOr("href", "")),
			Location:     jobLocation,
			Company:      "Intuit",
		})
	})

	return jobs, nil
}
