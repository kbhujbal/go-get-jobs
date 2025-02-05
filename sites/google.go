package sites

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/neyaadeez/go-get-jobs/common"
)

func GetGoogleJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Google")
	client := common.GetClient()

	url := "https://careers.google.com/jobs/results/?location=United%20States&target_level=EARLY&target_level=INTERN_AND_APPRENTICE&sort_by=date"

	var jobs []common.JobPosting
	for url != "" {
		resp, err := client.R().Get(url)
		if err != nil {
			return jobs, fmt.Errorf("error accessing the URL: %v", err)
		}

		var jb []common.JobPosting
		jb, url, err = parseJobPage(resp.Body())
		if err != nil {
			return jobs, err
		}

		jobs = append(jobs, jb...)
	}

	return jobs, nil
}

// Function to parse the HTML content using goquery
func parseJobPage(body []byte) ([]common.JobPosting, string, error) {
	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, "", fmt.Errorf("error parsing the HTML: %v", err)
	}

	var jobs []common.JobPosting
	nextPageLnk := ""

	doc.Find(".WpHeLc").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		jobId := href
		jbid := strings.Split(jobId, "/")
		if len(jbid) > 2 {
			jbid = strings.Split(jbid[2], "-")
			if len(jbid) > 1 {
				jobId = jbid[0]
			}
		}
		ariaLabel, _ := s.Attr("aria-label")
		if strings.Contains(ariaLabel, "next page") {
			nextPageLnk = href
			return
		}

		if strings.Contains(ariaLabel, "previous page") {
			return
		}

		jobs = append(jobs, common.JobPosting{
			JobId:        common.Google + ":" + jobId,
			JobTitle:     strings.ReplaceAll(ariaLabel, "Learn more about", ""),
			ExternalPath: "https://www.google.com/about/careers/applications/" + href,
			Company:      "Google",
		})
	})

	return jobs, nextPageLnk, nil
}
