package sites

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/neyaadeez/go-get-jobs/common"
)

func GetChimeJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Chime")
	client := common.GetClient()
	client.SetTimeout(5 * time.Second)

	url := "https://careers.chime.com/en/jobs"

	var jobs []common.JobPosting

	resp, err := client.R().Get(url)
	time.Sleep(2 * time.Second)
	if err != nil {
		return jobs, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jb []common.JobPosting
	jb, _, err = parseChimeJobPage(resp.Body())
	if err != nil {
		return jobs, err
	}

	jobs = append(jobs, jb...)

	return jobs, nil
}

// Function to parse the HTML content using goquery
func parseChimeJobPage(body []byte) ([]common.JobPosting, string, error) {
	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, "", fmt.Errorf("error parsing the HTML: %v", err)
	}

	var jobs []common.JobPosting
	nextPageLnk := ""

	doc.Find(".card-job-actions.js-job").Each(func(i int, s *goquery.Selection) {
		jobTitle, exists := s.Attr("data-jobtitle")
		if !exists {
			return
		}

		jobID, exists := s.Attr("data-id")
		if !exists {
			return
		}

		jobs = append(jobs, common.JobPosting{
			JobId:        common.Chime + ":" + jobID,
			JobTitle:     strings.TrimSpace(jobTitle),
			Location:     "USA",
			ExternalPath: "https://careers.chime.com/en/jobs/" + jobID,
			Company:      "Chime",
		})
	})

	nextPage := doc.Find(".pagination .page-item.next a")
	if nextPage.Length() > 0 {
		href, exists := nextPage.Attr("href")
		if exists {
			nextPageLnk = href
		}
	}

	return jobs, nextPageLnk, nil
}
