package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type AmazonMain struct {
	SearchHits []AmazonSearchHit `json:"searchHits"`
}

type AmazonSearchHit struct {
	Fields map[string][]string `json:"fields"`
}

func GetAmazonJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Amazon")
	client := common.GetClient()

	url := "https://www.amazon.jobs/api/jobs/search"
	payload := `{
  "accessLevel": "EXTERNAL",
  "contentFilterFacets": [
    {
      "name": "primarySearchLabel",
      "requestedFacetCount": 9999
    }
  ],
  "excludeFacets": [
    {
      "name": "isConfidential",
      "values": [
        {
          "name": "1"
        }
      ]
    },
    {
      "name": "businessCategory",
      "values": [
        {
          "name": "a-confidential-job"
        }
      ]
    }
  ],
  "filterFacets": [
    {
      "name": "category",
      "requestedFacetCount": 9999,
      "values": [
        {
          "name": "Software Development"
        }
      ]
    }
  ],
  "includeFacets": [],
  "jobTypeFacets": [],
  "locationFacets": [
    [
      {
        "name": "country",
        "requestedFacetCount": 9999,
        "values": [
          {
            "name": "US"
          }
        ]
      },
      {
        "name": "normalizedStateName",
        "requestedFacetCount": 9999
      },
      {
        "name": "normalizedCityName",
        "requestedFacetCount": 9999
      }
    ]
  ],
  "query": "",
  "size": 1000,
  "start": 0,
  "treatment": "OM",
  "cookieInfo": "__Host-mons-sid=356-9610498-7812526; preferred_locale=en-US; __Host-mons-ubid=355-3579436-6210606; cwr_u=eba9be99-339f-45a4-aaf3-7f5ad79b3015; _ga=GA1.2.485601763.1724572502; _gcl_au=1.1.1637541891.1724572502; check_for_eu_countries=false; uid=2987a145-3397-4d55-a7c7-8c27cf2abf3a; peoplesoft_id=; passport_id=d2be9358-7b07-42e2-92b4-d2cd5c0de24e; cookie_preferences=%7B%22advertising%22%3Atrue%2C%22analytics%22%3Atrue%2C%22version%22%3A2%7D; advertising_id=a2fc98d6-ff36-4e71-b50f-1638836e894c; AMCV_4EE1BB6555F9369A7F000101%40AdobeOrg=-1124106680%7CMCIDTS%7C19983%7CMCMID%7C47412749704312236853749202927108312944%7CMCAAMLH-1727132315%7C9%7CMCAAMB-1727132315%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1726534715s%7CNONE%7CvVersion%7C5.2.0; analytics_id=5c235370-24ca-477d-9956-bb4c7109910e; tracking_id=f8a9337f4cc7b08259578365590ed3a3; __Host-mons-st=5z+CNejUH9luVPiO7/YoIpWZgr06oDBbk7KmpjJ+AmjtpY8FCBvP5pZZMIcKvXbirTX/E0Zi0mWe16ShswW2VN/kM0Rk/A0zLlMYaIAPRG/cisF/BfN0VXfmJTVxpQG6wZ8dvGLfCxT/AbcWn8g4SoOMFjHK3hl5ZEUYetVZbfOAkLHC1eVaiAMkE/OMeNNUnYOlGOzQ3bgzbQsCEFxtmZe53Jy4tjH3DZ8VLbN8oCEGLZHwk+bK4hamqEapBa9Qtdc7xAZDNsLZQ4CY9RikoWZFWYVVNZqt9eksDMAVwrtNc4Yc4XrmjO1yaCsdQrkbmA51+AJy4TZfTq8KsVkRI/358HbIAQgQyL/fH5PJXjw=; source=%7B%22azref%22%3A%22https%3A%2F%2Fwww.google.com%2F%22%7D; _gat=1; AMCVS_CCBC879D5572070E7F000101%40AdobeOrg=1; AMCV_CCBC879D5572070E7F000101%40AdobeOrg=-1124106680%7CMCIDTS%7C19997%7CMCMID%7C47416669608961527963748812107031669801%7CMCAAMLH-1728288936%7C9%7CMCAAMB-1728288936%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1727691336s%7CNONE%7CMCAID%7CNONE%7CvVersion%7C5.2.0; csm-sid=165-3855017-2546422; s_lv_s=More%20than%207%20days; s_cc=true; cwr_s_f8eef880-f8cb-4128-9b7d-a097c3324176=eyJzZXNzaW9uSWQiOiJhOGYzZTJhZS1mYzkwLTRlYWItYjZmYy03MWIzYmNlM2RhN2YiLCJyZWNvcmQiOnRydWUsImV2ZW50Q291bnQiOjMwLCJwYWdlIjp7InBhZ2VJZCI6Ii9lbi8iLCJpbnRlcmFjdGlvbiI6MCwicmVmZXJyZXIiOiJodHRwczovL3d3dy5nb29nbGUuY29tLyIsInJlZmVycmVyRG9tYWluIjoid3d3Lmdvb2dsZS5jb20iLCJzdGFydCI6MTcyNzY4NDEzNjQxN319; gpv=Amazon.jobs%20%7C%20Job%20Categories%20%7C%20Software%20Development; s_sq=%5B%5BB%5D%5D; s_ppvl=Amazon.jobs%2520%257C%2520Job%2520Categories%2520%257C%2520Software%2520Development%2C55%2C55%2C2185%2C669%2C812%2C1440%2C900%2C2%2CL; s_ppv=Amazon.jobs%2520%257C%2520Job%2520Categories%2520%257C%2520Software%2520Development%2C60%2C29%2C4670%2C669%2C812%2C1440%2C900%2C2%2CL; s_lv=1727684538902; s_nr30=1727684538903-Repeat; cwr_s_2cf45a3b-4b97-4b16-86e2-fa97eb8dc7f9=eyJzZXNzaW9uSWQiOiIxMWIwYTIwMi0wMWU4LTQyNzItYjgxNC1kN2RmMTRiZTAxMmUiLCJyZWNvcmQiOnRydWUsImV2ZW50Q291bnQiOjI4MSwicGFnZSI6eyJwYWdlSWQiOiIvY29udGVudC9lbi9qb2ItY2F0ZWdvcmllcy9zb2Z0d2FyZS1kZXZlbG9wbWVudCIsImludGVyYWN0aW9uIjowLCJyZWZlcnJlciI6Imh0dHBzOi8vd3d3LmFtYXpvbi5qb2JzL2VuLyIsInJlZmVycmVyRG9tYWluIjoid3d3LmFtYXpvbi5qb2JzIiwic3RhcnQiOjE3Mjc2ODQxNDc1Nzh9fQ==",
  "sort": {
    "sortOrder": "DESCENDING",
    "sortType": "CREATED_DATE"
  }
}`

	resp, err := client.R().SetBody(payload).Post(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(amazon jobs): %v", err)
	}

	var amazonJobs AmazonMain
	err = json.Unmarshal(resp.Body(), &amazonJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (amazon jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range amazonJobs.SearchHits {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      job.Fields["companyName"][0],
			JobId:        common.Amazon + ":" + job.Fields["icimsJobId"][0],
			JobTitle:     job.Fields["title"][0],
			Location:     job.Fields["location"][0],
			PostedOn:     job.Fields["createdDate"][0],
			ExternalPath: "https://www.amazon.jobs/en/jobs/" + job.Fields["icimsJobId"][0],
		})
	}

	return jobPostings, nil
}
