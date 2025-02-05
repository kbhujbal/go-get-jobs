package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type IBMMain struct {
	Hits IBMHits `json:"hits"`
}

type IBMHits struct {
	Total    IBMTotal    `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []IBMHit    `json:"hits"`
}

type IBMHit struct {
	Index  string      `json:"_index"`
	ID     string      `json:"_id"`
	Score  interface{} `json:"_score"`
	Source IBMSource   `json:"_source"`
	Sort   []int64     `json:"sort"`
}

type IBMSource struct {
	Title          string `json:"title"`
	URL            string `json:"url"`
	Description    string `json:"description"`
	Entitled       string `json:"entitled"`
	FieldKeyword08 string `json:"field_keyword_08"`
	FieldKeyword17 string `json:"field_keyword_17"`
	FieldKeyword18 string `json:"field_keyword_18"`
	FieldKeyword19 string `json:"field_keyword_19"`
}

type IBMTotal struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

func GetIBMJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "IBM")
	client := common.GetClient()

	url := "https://www-api.ibm.com/search/api/v2"
	payload := `{
  "appId": "careers",
  "scopes": [
    "careers"
  ],
  "query": {
    "bool": {
      "must": []
    }
  },
  "post_filter": {
    "bool": {
      "must": [
        {
          "term": {
            "field_keyword_18": "Entry Level"
          }
        },
        {
          "term": {
            "field_keyword_05": "United States"
          }
        }
      ]
    }
  },
  "aggs": {
    "field_keyword_172": {
      "filter": {
        "bool": {
          "must": [
            {
              "term": {
                "field_keyword_18": "Entry Level"
              }
            },
            {
              "term": {
                "field_keyword_05": "United States"
              }
            }
          ]
        }
      },
      "aggs": {
        "field_keyword_17": {
          "terms": {
            "field": "field_keyword_17",
            "size": 6
          }
        },
        "field_keyword_17_count": {
          "cardinality": {
            "field": "field_keyword_17"
          }
        }
      }
    },
    "field_keyword_083": {
      "filter": {
        "bool": {
          "must": [
            {
              "term": {
                "field_keyword_18": "Entry Level"
              }
            },
            {
              "term": {
                "field_keyword_05": "United States"
              }
            }
          ]
        }
      },
      "aggs": {
        "field_keyword_08": {
          "terms": {
            "field": "field_keyword_08",
            "size": 6
          }
        },
        "field_keyword_08_count": {
          "cardinality": {
            "field": "field_keyword_08"
          }
        }
      }
    },
    "field_keyword_184": {
      "filter": {
        "term": {
          "field_keyword_05": "United States"
        }
      },
      "aggs": {
        "field_keyword_18": {
          "terms": {
            "field": "field_keyword_18",
            "size": 6
          }
        },
        "field_keyword_18_count": {
          "cardinality": {
            "field": "field_keyword_18"
          }
        }
      }
    },
    "field_keyword_055": {
      "filter": {
        "term": {
          "field_keyword_18": "Entry Level"
        }
      },
      "aggs": {
        "field_keyword_05": {
          "terms": {
            "field": "field_keyword_05",
            "size": 1000
          }
        },
        "field_keyword_05_count": {
          "cardinality": {
            "field": "field_keyword_05"
          }
        }
      }
    }
  },
  "size": 100,
  "sort": [
    {
      "dcdate": "desc"
    }
  ],
  "lang": "zz",
  "localeSelector": {},
  "sm": {
    "query": "",
    "lang": "zz"
  },
  "_source": [
    "_id",
    "title",
    "url",
    "description",
    "language",
    "entitled",
    "field_keyword_17",
    "field_keyword_08",
    "field_keyword_18",
    "field_keyword_19"
  ]
}`

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("Accept", "application/json").SetBody(payload).Post(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(ibm jobs): %v", err)
	}

	var ibmJobs IBMMain
	err = json.Unmarshal(resp.Body(), &ibmJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (ibm jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range ibmJobs.Hits.Hits {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "IBM",
			JobId:        common.IBM + ":" + job.ID,
			JobTitle:     job.Source.Title,
			Location:     job.Source.FieldKeyword19,
			ExternalPath: job.Source.URL,
		})
	}

	return jobPostings, nil
}
