package sites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/neyaadeez/go-get-jobs/common"
)

func GetAppleJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Apple")
	var jobPostings []common.JobPosting
	totalPages := 15

	for page := 1; page <= totalPages; page++ {
		jobs, err := appleJobs(page)
		if err != nil {
			fmt.Printf("Error fetching jobs on page %d: %v\n", page, err)
			continue
		}

		jobPostings = append(jobPostings, jobs...)
	}

	return jobPostings, nil
}

// The updated appleJobs function to return jobs for each page.
func appleJobs(page int) ([]common.JobPosting, error) {
	csrfUrl := "https://jobs.apple.com/api/csrfToken"
	client := common.GetClient()

	resp, err := client.R().Get(csrfUrl)
	if err != nil {
		return nil, fmt.Errorf("error fetching CSRF token: %v", err)
	}
	csrfToken := resp.Header().Get("x-apple-csrf-token")
	if csrfToken == "" {
		return nil, fmt.Errorf("CSRF token not found (apple jobs)")
	}

	apiUrl := "https://jobs.apple.com/api/role/search"
	payload := fmt.Sprintf(`{
  "query": "",
  "filters": {
    "postingpostLocation": [
      "postLocation-USA"
    ],
    "range": {
      "standardWeeklyHours": {
        "start": null,
        "end": null
      }
    },
    "teams": [
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-CLD"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-COS"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-DSR"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-EPM"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-ISTECH"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-MCHLN"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-SEC"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-SQAT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-WSFT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-ARCH"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-CAM"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-DISP"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-EPM"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-ENVT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-HT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-MCHLN"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-ME"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-PE"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-REL"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-SENT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-SILT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-SDE"
      },
      {
        "teams.teamID": "teamsAndSubTeams-HRDWR",
        "teams.subTeamID": "subTeam-WT"
      },
      {
        "teams.teamID": "teamsAndSubTeams-STDNT",
        "teams.subTeamID": "subTeam-INTRN"
      },
      {
        "teams.teamID": "teamsAndSubTeams-STDNT",
        "teams.subTeamID": "subTeam-CORP"
      },
      {
        "teams.teamID": "teamsAndSubTeams-STDNT",
        "teams.subTeamID": "subTeam-ACR"
      },
      {
        "teams.teamID": "teamsAndSubTeams-MLAI",
        "teams.subTeamID": "subTeam-MLI"
      },
      {
        "teams.teamID": "teamsAndSubTeams-MLAI",
        "teams.subTeamID": "subTeam-DLRL"
      },
      {
        "teams.teamID": "teamsAndSubTeams-MLAI",
        "teams.subTeamID": "subTeam-NLP"
      },
      {
        "teams.teamID": "teamsAndSubTeams-MLAI",
        "teams.subTeamID": "subTeam-CV"
      },
      {
        "teams.teamID": "teamsAndSubTeams-MLAI",
        "teams.subTeamID": "subTeam-AR"
      },
      {
        "teams.teamID": "teamsAndSubTeams-SFTWR",
        "teams.subTeamID": "subTeam-AF"
      }
    ]
  },
  "page": %d,
  "locale": "en-us",
  "sort": "newest"
}`, page)

	apiReq, err := http.NewRequest("POST", apiUrl, strings.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating API request(apple jobs): %v", err)
	}

	apiReq.Header.Add("Content-Type", "application/json")
	apiReq.Header.Add("X-CSRF-Token", csrfToken)

	client1 := &http.Client{}
	apiResp, err := client1.Do(apiReq)
	if err != nil {
		return nil, fmt.Errorf("error making API request(apple jobs): %v", err)
	}
	defer apiResp.Body.Close()

	body, err := io.ReadAll(apiResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading API response(apple jobs): %v", err)
	}

	var jobs AppleMain
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing API response (apple jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range jobs.AppleSearchResults {
		jobPosting := common.JobPosting{
			JobId:        common.Apple + ":" + job.ID,
			JobTitle:     job.PostingTitle,
			Location:     job.Locations[0].City + " USA",
			PostedOn:     string(job.PostingDate),
			ExternalPath: "https://jobs.apple.com/en-us/details/" + job.ID,
			Company:      "Apple",
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return jobPostings, nil
}

type AppleMain struct {
	AppleSearchResults []AppleSearchResult `json:"searchResults"`
	TotalRecords       int64               `json:"totalRecords"`
}

type AppleSearchResult struct {
	ID string `json:"id"`
	//JobSummary   string      `json:"jobSummary"`
	Locations    []AppleLocation `json:"locations"`
	PostingDate  string          `json:"postingDate"`
	PostingTitle string          `json:"postingTitle"`
}

type AppleLocation struct {
	City string `json:"city"`
}
