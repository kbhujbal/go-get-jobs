package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type ABBMain struct {
	RefineSearch ABBRefineSearch `json:"refineSearch"`
}

type ABBRefineSearch struct {
	Status    int64   `json:"status"`
	Hits      int64   `json:"hits"`
	TotalHits int64   `json:"totalHits"`
	Data      ABBData `json:"data"`
}

type ABBData struct {
	Jobs []Job `json:"jobs"`
	// Aggregations      []Aggregation `json:"aggregations"`
	// SearchConfig      SearchConfig  `json:"SEARCH_CONFIG"`
	// Suggestions       LocationData  `json:"suggestions"`
	// UISkillsSelection interface{}   `json:"ui_skills_selection"`
	// UISelections      UISelections  `json:"ui_selections"`
	// LocationData      LocationData  `json:"locationData"`
}

// type Aggregation struct {
// 	Field string           `json:"field"`
// 	Value map[string]int64 `json:"value"`
// }

type Job struct {
	// SubCategory        SubCategory          `json:"subCategory"`
	// ContractType       ContractType         `json:"contractType"`
	// MlSkills           []string             `json:"ml_skills"`
	// Type               Type                 `json:"type"`
	// DescriptionTeaser  string               `json:"descriptionTeaser"`
	// State              string               `json:"state"`
	// SiteType           SiteType             `json:"siteType"`
	// MultiCategory      []Category           `json:"multi_category"`
	// ProductAlignment   string               `json:"productAlignment"`
	ReqID string `json:"reqId"`
	// Grade              string               `json:"grade"`
	// City               string               `json:"city"`
	// BusinessSegment    BusinessSegment      `json:"businessSegment"`
	// Latitude           string               `json:"latitude"`
	// Industry           Industry             `json:"industry"`
	// MultiLocation      []string             `json:"multi_location"`
	Address string `json:"address"`
	// MlJobParser        MlJobParser          `json:"ml_job_parser"`
	// ExternalApply      bool                 `json:"externalApply"`
	// CityState          string               `json:"cityState"`
	// Country            Query                `json:"country"`
	// VisibilityType     VisibilityType       `json:"visibilityType"`
	// Longitude          string               `json:"longitude"`
	JobID string `json:"jobId"`
	// Locale             Locale               `json:"locale"`
	Title string `json:"title"`
	// JobSeqNo           string               `json:"jobSeqNo"`
	PostedDate string `json:"postedDate"`
	// DateCreated        string               `json:"dateCreated"`
	// CityStateCountry   string               `json:"cityStateCountry"`
	// JobType            Type                 `json:"jobType"`
	// JobVisibility      []SiteType           `json:"jobVisibility"`
	Location string `json:"location"`
	// CityCountry        string               `json:"cityCountry"`
	// Category           Category             `json:"category"`
	// IsMultiLocation    bool                 `json:"isMultiLocation"`
	// MultiLocationArray []MultiLocationArray `json:"multi_location_array"`
	// IsMultiCategory    bool                 `json:"isMultiCategory"`
	// MultiCategoryArray []MultiCategoryArray `json:"multi_category_array"`
	// Badge              string               `json:"badge"`
}

// type MlJobParser struct {
// 	DescriptionTeaser         string `json:"descriptionTeaser"`
// 	DescriptionTeaserAts      string `json:"descriptionTeaser_ats"`
// 	DescriptionTeaserKeyword  string `json:"descriptionTeaser_keyword"`
// 	DescriptionTeaserFirst200 string `json:"descriptionTeaser_first200"`
// }

// type MultiCategoryArray struct {
// 	Category Category `json:"category"`
// }

// type MultiLocationArray struct {
// 	Location string  `json:"location"`
// 	Latlong  Latlong `json:"latlong"`
// }

// type Latlong struct {
// 	Lon float64 `json:"lon"`
// 	Lat float64 `json:"lat"`
// }

// type LocationData struct {
// }

// type SearchConfig struct {
// 	ContextualSearch     bool `json:"contextualSearch"`
// 	IsSuggestionsEnabled bool `json:"isSuggestionsEnabled"`
// }

// type UISelections struct {
// 	Country  []Query    `json:"country"`
// 	Category []Category `json:"category"`
// }

// type Eid struct {
// 	TrialIndex int64  `json:"trialIndex"`
// 	Eid        string `json:"eid"`
// 	SearchType string `json:"searchType"`
// 	Query      Query  `json:"query"`
// 	Variant    int64  `json:"variant"`
// 	BanditID   string `json:"banditId"`
// }

// type BusinessSegment string
// const (
// 	Electrification BusinessSegment = "Electrification"
// 	Motion BusinessSegment = "Motion"
// 	ProcessAutomation BusinessSegment = "Process Automation"
// 	RoboticsDiscreteAutomation BusinessSegment = "Robotics & Discrete Automation"
// )

// type Category string
// const (
// 	CategoryDigital Category = "Digital"
// 	Engineering Category = "Engineering"
// 	InformationSystems Category = "Information Systems"
// 	ResearchDevelopment Category = "Research & Development"
// )

// type ContractType string
// const (
// 	Consultant ContractType = "Consultant"
// 	RegularPermanent ContractType = "Regular/Permanent"
// 	Temporary ContractType = "Temporary"
// 	Traineeship ContractType = "Traineeship"
// )

// type Query string
// const (
// 	Global Query = "Global"
// 	Usa Query = "USA"
// )

// type Industry string
// const (
// 	AllOthers Industry = "All Others"
// )

// type Type string
// const (
// 	FullTime Type = "Full-Time"
// 	PartTime Type = "Part-time"
// )

// type SiteType string
// const (
// 	External SiteType = "external"
// )

// type Locale string
// const (
// 	EnGLOBAL Locale = "en_GLOBAL"
// )

// type SubCategory string
// const (
// 	DesignAndEngineering SubCategory = "Design and Engineering"
// 	InformationTechnology SubCategory = "Information Technology"
// 	ResearchAndDevelopment SubCategory = "Research and Development"
// 	SubCategoryDigital SubCategory = "Digital"
// )

// type VisibilityType string
// const (
// 	VisibilityTypeExternal VisibilityType = "External"
// )

func GetABBJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "ABB")
	client := common.GetClient()

	url := "https://careers.abb/widgets"
	payload := `{
  "lang": "en_global",
  "deviceType": "desktop",
  "country": "global",
  "pageName": "digital",
  "ddoKey": "refineSearch",
  "sortBy": "Most recent",
  "subsearch": "",
  "from": 0,
  "jobs": true,
  "counts": true,
  "all_fields": [
    "category",
    "subCategory",
    "businessSegment",
    "businessSegmentDescr",
    "continent",
    "country",
    "city",
    "contractType",
    "jobLevel",
    "jobType"
  ],
  "pageType": "category",
  "size": 1000,
  "clearAll": false,
  "jdsource": "facets",
  "isSliderEnable": false,
  "pageId": "page1053",
  "siteType": "external",
  "keywords": "",
  "global": true,
  "selected_fields": {
    "country": [
      "USA"
    ],
    "category": [
      "Digital",
      "Engineering",
      "Information Systems",
      "Research & Development"
    ]
  },
  "sort": {
    "order": "desc",
    "field": "postedDate"
  },
  "locationData": {}
}`

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("Accept", "application/json").SetBody(payload).Post(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(ABB jobs): %v", err)
	}

	var abbJobs ABBMain
	err = json.Unmarshal(resp.Body(), &abbJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (ABB jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range abbJobs.RefineSearch.Data.Jobs {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "ABB",
			JobId:        common.ABB + ":" + job.ReqID,
			JobTitle:     job.Title,
			Location:     job.Location,
			PostedOn:     job.PostedDate,
			ExternalPath: "https://careers.abb/global/en/job/" + job.ReqID,
		})
	}

	return jobPostings, nil
}
