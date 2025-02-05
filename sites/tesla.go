package sites

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
)

type TeslaMain struct {
	Lookup      TeslaLookup         `json:"lookup"`
	Departments map[string][]string `json:"departments"`
	Geo         []TeslaGeo          `json:"geo"`
	Listings    []TeslaListing      `json:"listings"`
}

type TeslaGeo struct {
	ID    string      `json:"id"`
	Sites []TeslaSite `json:"sites"`
}

type TeslaSite struct {
	ID     string              `json:"id"`
	States []TeslaState        `json:"states,omitempty"`
	Cities map[string][]string `json:"cities,omitempty"`
}

type TeslaState struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	Cities map[string][]string `json:"cities"`
}

type TeslaListing struct {
	ID string `json:"id"`
	T  string `json:"t"`
	DP string `json:"dp"`
	F  string `json:"f"`
	L  string `json:"l"`
	Y  int64  `json:"y"`
}

type TeslaLookup struct {
	Regions     map[string]string `json:"regions"`
	Sites       map[string]string `json:"sites"`
	Locations   map[string]string `json:"locations"`
	Departments map[string]string `json:"departments"`
	Types       map[string]string `json:"types"`
}

func GetTeslaJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Tesla")
	var jobPostings []common.JobPosting

	cookies, err := getTeslaCookies()
	if err != nil {
		return nil, fmt.Errorf("error getting Tesla cookies: %v", err)
	}

	teslaJobsData, err := fetchTeslaJobsData(cookies)
	if err != nil {
		return nil, fmt.Errorf("error fetching Tesla jobs: %v", err)
	}

	jobs := filterUSJobs(teslaJobsData)
	for _, job := range jobs {
		jobPosting := common.JobPosting{
			JobId:        common.Tesla + ":" + job.ID,
			JobTitle:     job.T,
			Location:     teslaJobsData.Lookup.Locations[job.L],
			ExternalPath: "https://www.tesla.com/careers/job/" + job.ID,
			Company:      "Tesla",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, nil
}

func getTeslaCookies() (string, error) {
	url := "https://www.tesla.com/careers/search"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	// Set required headers
	req.Header.Set("Host", "www.tesla.com")
	req.Header.Set("Sec-Ch-Ua", `"Not;A=Brand";v="24", "Chromium";v="128"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.120 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	cookies := resp.Header.Get("Set-Cookie")
	if cookies == "" {
		return "", fmt.Errorf("no cookies found in Tesla response")
	}

	return cookies, nil
}

func fetchTeslaJobsData(cookies string) (*TeslaMain, error) {
	url := "https://www.tesla.com/cua-api/apps/careers/state"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set required headers and cookies
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Sec-Ch-Ua", `"Not;A=Brand";v="24", "Chromium";v="128"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.120 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.tesla.com/careers/search")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Priority", "u=1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response body: %s", resp.StatusCode, bodyBytes)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader: %v", err)
		}
		defer reader.Close()
	case "deflate":
		reader, err = zlib.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error creating zlib reader: %v", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var teslaResponse TeslaMain
	if err := json.Unmarshal(bodyBytes, &teslaResponse); err != nil {
		return nil, fmt.Errorf("error parsing API response: %v", err)
	}

	return &teslaResponse, nil
}

var usStates = []string{
	"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware",
	"Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky",
	"Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri",
	"Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York",
	"North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island",
	"South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington",
	"West Virginia", "Wisconsin", "Wyoming", "District Of Columbia", "United States",
}

var departments = []string{
	"Vehicle Service", "Engineering & Information Technology", "AI & Robotics", "Vehicle Software",
}

func isUSLocation(location string) bool {
	for _, state := range usStates {
		if strings.Contains(location, state) {
			return true
		}
	}
	return false
}

func isCSRelatedJob(dp string) bool {
	for _, department := range departments {
		if strings.Contains(dp, department) {
			return true
		}
	}
	return false
}

func filterUSJobs(teslaData *TeslaMain) []TeslaListing {
	usJobs := []TeslaListing{}

	for _, job := range teslaData.Listings {
		location := teslaData.Lookup.Locations[job.L]
		department := teslaData.Lookup.Departments[job.DP]
		if isUSLocation(location) && isCSRelatedJob(department) {
			usJobs = append(usJobs, job)
		}
	}

	return usJobs
}
