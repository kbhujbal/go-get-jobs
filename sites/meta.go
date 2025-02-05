package sites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
)

type MetaMain struct {
	Data MetaData `json:"data"`
}

type MetaData struct {
	JobSearch []MetaJobSearch `json:"job_search"`
}

type MetaJobSearch struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Locations []string `json:"locations"`
	Teams     []string `json:"teams"`
	SubTeams  []string `json:"sub_teams"`
}

func GetMetaJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Meta")
	var jobPostings []common.JobPosting

	jobs, err := metaJobs()
	if err != nil {
		return nil, fmt.Errorf("error fetching Meta jobs: %v", err)
	}

	jobPostings = append(jobPostings, jobs...)

	return jobPostings, nil
}

func metaJobs() ([]common.JobPosting, error) {
	lsdTkn, datrTkn, err := getLsdAndDatr()
	if err != nil {
		return nil, fmt.Errorf("error while processing meta jobs: %v", err)
	}

	apiUrl := "https://www.metacareers.com/graphql"

	formData := url.Values{}
	formData.Set("lsd", lsdTkn)
	formData.Set("variables", `{"search_input":{"q":"","divisions":[],"offices":["North America"],"roles":[],"leadership_levels":["Individual Contributor"],"saved_jobs":[],"saved_searches":[],"sub_teams":[],"teams":[],"is_leadership":false,"is_remote_only":false,"sort_by_new":true,"page":1,"results_per_page":null}}`)
	formData.Set("doc_id", "9114524511922157")

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Cookie", fmt.Sprintf("datr=%s;", datrTkn))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request (Meta jobs): %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response body: %s", resp.StatusCode, bodyBytes)
	}

	var MetaMainResponse MetaMain
	if err := json.NewDecoder(resp.Body).Decode(&MetaMainResponse); err != nil {
		return nil, fmt.Errorf("error parsing API response (Meta jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range MetaMainResponse.Data.JobSearch {
		jobPosting := common.JobPosting{
			JobId:        common.Meta + ":" + job.ID,
			JobTitle:     job.Title,
			Location:     strings.Join(job.Locations, ", "),
			ExternalPath: "https://www.metacareers.com/jobs/" + job.ID,
			Company:      "Meta",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, nil
}

func getLsdAndDatr() (string, string, error) {
	url := "https://www.metacareers.com/jobs"
	datrTkn := ""
	lsdTkn := ""

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading response body: %v", err)
	}

	bodyStr := string(body)

	if strings.Contains(bodyStr, "_js_datr") {
		startIndex := strings.Index(bodyStr, "_js_datr")
		lines := strings.Split(bodyStr[startIndex:], ",")
		if len(lines) > 1 {
			datrTkn = strings.ReplaceAll(lines[1], "\"", "")
		}
	} else {
		return "", "", fmt.Errorf("unable to retrive datr token(meta jobs)")
	}

	if strings.Contains(bodyStr, "LSD") {
		startIndex := strings.Index(bodyStr, "LSD")
		lines := strings.Split(bodyStr[startIndex:], ",")
		if len(lines) > 2 {
			tokenLine := strings.TrimSpace(lines[2])

			re := regexp.MustCompile(`"token"\s*:\s*"([^"]+)"`)
			matches := re.FindStringSubmatch(tokenLine)

			if len(matches) > 1 {
				lsdTkn = matches[1]
			} else {
				return "", "", fmt.Errorf("no token found in the line")
			}
		}
	} else {
		return "", "", fmt.Errorf("unable to retrive lsd token(meta jobs)")
	}

	return lsdTkn, datrTkn, nil
}
