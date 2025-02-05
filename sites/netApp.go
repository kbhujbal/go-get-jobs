package sites

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
	"golang.org/x/net/html"
)

type NetAppMain struct {
	Filters    string `json:"filters"`
	Results    string `json:"results"`
	HasJobs    bool   `json:"hasJobs"`
	HasContent bool   `json:"hasContent"`
}

func GetNetAppJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Net App")
	client := common.GetClient()

	url := "https://careers.netapp.com/search-jobs/results?ActiveFacetID=8604064&CurrentPage=1&RecordsPerPage=1000&Distance=500&RadiusUnitType=0&Keywords&Location&ShowRadius=False&IsPagination=False&CustomFacetName&FacetTerm&FacetType=0&FacetFilters%5B0%5D.ID=8607760&FacetFilters%5B0%5D.FacetType=1&FacetFilters%5B0%5D.Display=Cloud&FacetFilters%5B0%5D.IsApplied=true&FacetFilters%5B0%5D.FieldName&FacetFilters%5B1%5D.ID=8604160&FacetFilters%5B1%5D.FacetType=1&FacetFilters%5B1%5D.Display=Engineering&FacetFilters%5B1%5D.IsApplied=true&FacetFilters%5B1%5D.FieldName&FacetFilters%5B2%5D.ID=8604176&FacetFilters%5B2%5D.FacetType=1&FacetFilters%5B2%5D.Display=Information+Technology&FacetFilters%5B2%5D.IsApplied=true&FacetFilters%5B2%5D.FieldName&FacetFilters%5B3%5D.ID=8604192&FacetFilters%5B3%5D.FacetType=1&FacetFilters%5B3%5D.Count=9&FacetFilters%5B3%5D.Display=Software+Engineering&FacetFilters%5B3%5D.IsApplied=true&FacetFilters%5B3%5D.FieldName&FacetFilters%5B4%5D.ID=8604064&FacetFilters%5B4%5D.FacetType=1&FacetFilters%5B4%5D.Display=Systems+Engineering&FacetFilters%5B4%5D.IsApplied=true&FacetFilters%5B4%5D.FieldName&FacetFilters%5B5%5D.ID=8604240&FacetFilters%5B5%5D.FacetType=1&FacetFilters%5B5%5D.Display=University&FacetFilters%5B5%5D.IsApplied=true&FacetFilters%5B5%5D.FieldName&FacetFilters%5B6%5D.ID=6252001&FacetFilters%5B6%5D.FacetType=2&FacetFilters%5B6%5D.Display=United+States&FacetFilters%5B6%5D.IsApplied=true&FacetFilters%5B6%5D.FieldName&SearchResultsModuleName=Search+Results&SearchFiltersModuleName=Search+Filters&SortCriteria=2&SortDirection=0&SearchType=5&PostalCode&ResultsType=0&fc&fl&fcf&afc&afl&afcf"
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(Net App jobs): %v", err)
	}

	var netAppJobs NetAppMain
	err = json.Unmarshal(resp.Body(), &netAppJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (Net App jobs): %v", err)
	}

	htmlReader := strings.NewReader(netAppJobs.Results)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	jobPostings := extractJobDetails(doc)
	return jobPostings, nil
}

func extractJobDetails(n *html.Node) []common.JobPosting {
	var jobs []common.JobPosting

	if n.Type == html.ElementNode && n.Data == "a" {
		var jobID, jobTitle, externalPath, jobLocation string

		for _, attr := range n.Attr {
			if attr.Key == "data-job-id" {
				jobID = attr.Val
			}

			if attr.Key == "href" {
				externalPath = "https://careers.netapp.com" + attr.Val
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.ElementNode && child.Data == "h3" {
				jobTitle = child.FirstChild.Data
			}
		}

		for sibling := n.FirstChild; sibling != nil; sibling = sibling.NextSibling {
			if sibling.Type == html.ElementNode && sibling.Data == "span" {
				for _, attr := range sibling.Attr {
					if attr.Key == "class" && attr.Val == "job-location" {
						jobLocation = sibling.FirstChild.Data
					}
				}
			}
		}

		if jobID != "" && jobTitle != "" {
			job := common.JobPosting{
				Company:      "NetApp",
				JobId:        common.NetApp + ":" + jobID,
				JobTitle:     jobTitle,
				Location:     jobLocation,
				ExternalPath: externalPath,
			}
			jobs = append(jobs, job)
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		childJobs := extractJobDetails(child)
		jobs = append(jobs, childJobs...)
	}

	return jobs
}
