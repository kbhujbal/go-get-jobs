package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type FortinetMain struct {
	Items []FortinetItem `json:"items"`
	// Count   int64  `json:"count"`
	// HasMore bool   `json:"hasMore"`
	// Limit   int64  `json:"limit"`
	// Offset  int64  `json:"offset"`
	// Links   []Link `json:"links"`
}

type FortinetItem struct {
	// SearchID                    int64             `json:"SearchId"`
	// Keyword                     interface{}       `json:"Keyword"`
	// CorrectedKeyword            interface{}       `json:"CorrectedKeyword"`
	// UseExactKeywordFlag         bool              `json:"UseExactKeywordFlag"`
	// SuggestedKeyword            interface{}       `json:"SuggestedKeyword"`
	// ExecuteSpellCheckFlag       bool              `json:"ExecuteSpellCheckFlag"`
	// Location                    interface{}       `json:"Location"`
	// LocationID                  interface{}       `json:"LocationId"`
	// Radius                      int64             `json:"Radius"`
	// RadiusUnit                  string            `json:"RadiusUnit"`
	// SelectedTitlesFacet         string            `json:"SelectedTitlesFacet"`
	// SelectedCategoriesFacet     interface{}       `json:"SelectedCategoriesFacet"`
	// SelectedPostingDatesFacet   interface{}       `json:"SelectedPostingDatesFacet"`
	// SelectedLocationsFacet      string            `json:"SelectedLocationsFacet"`
	// LastSelectedFacet           string            `json:"LastSelectedFacet"`
	// Facets                      string            `json:"Facets"`
	// Offset                      int64             `json:"Offset"`
	// Limit                       int64             `json:"Limit"`
	// SortBy                      string            `json:"SortBy"`
	// TotalJobsCount              int64             `json:"TotalJobsCount"`
	// Latitude                    interface{}       `json:"Latitude"`
	// Longitude                   interface{}       `json:"Longitude"`
	// SiteNumber                  string            `json:"SiteNumber"`
	// JobFamilyID                 interface{}       `json:"JobFamilyId"`
	// PostingStartDate            interface{}       `json:"PostingStartDate"`
	// PostingEndDate              interface{}       `json:"PostingEndDate"`
	// SelectedWorkLocationsFacet  interface{}       `json:"SelectedWorkLocationsFacet"`
	// RequisitionID               interface{}       `json:"RequisitionId"`
	// CandidateNumber             interface{}       `json:"CandidateNumber"`
	// WorkLocationZipCode         interface{}       `json:"WorkLocationZipCode"`
	// WorkLocationCountryCode     interface{}       `json:"WorkLocationCountryCode"`
	// SelectedFlexFieldsFacets    interface{}       `json:"SelectedFlexFieldsFacets"`
	// OrganizationID              interface{}       `json:"OrganizationId"`
	// SelectedOrganizationsFacet  interface{}       `json:"SelectedOrganizationsFacet"`
	// UserTargetFacetName         interface{}       `json:"UserTargetFacetName"`
	// UserTargetFacetInputTerm    interface{}       `json:"UserTargetFacetInputTerm"`
	// HotJobFlag                  interface{}       `json:"HotJobFlag"`
	// WorkplaceType               interface{}       `json:"WorkplaceType"`
	// SelectedWorkplaceTypesFacet interface{}       `json:"SelectedWorkplaceTypesFacet"`
	// BotQRShortCode              interface{}       `json:"BotQRShortCode"`
	RequisitionList []FortinetRequisitionList `json:"requisitionList"`
	// CategoriesFacet             []SFacet          `json:"categoriesFacet"`
	// LocationsFacet              []SFacet          `json:"locationsFacet"`
	// PostingDatesFacet           []SFacet          `json:"postingDatesFacet"`
	// TitlesFacet                 []EsFacet         `json:"titlesFacet"`
	// WorkLocationsFacet          []SFacet          `json:"workLocationsFacet"`
	// FlexFieldsFacet             []interface{}     `json:"flexFieldsFacet"`
	// OrganizationsFacet          []SFacet          `json:"organizationsFacet"`
	// WorkplaceTypesFacet         []EsFacet         `json:"workplaceTypesFacet"`
}

// type SFacet struct {
// 	ID         int64  `json:"Id"`
// 	Name       string `json:"Name"`
// 	TotalCount int64  `json:"TotalCount"`
// }

type FortinetRequisitionList struct {
	ID         string `json:"Id"`
	Title      string `json:"Title"`
	PostedDate string `json:"PostedDate"`
	// PostingEndDate              interface{}         `json:"PostingEndDate"`
	// Language                    Language            `json:"Language"`
	// PrimaryLocationCountry      Language            `json:"PrimaryLocationCountry"`
	// GeographyID                 int64               `json:"GeographyId"`
	// HotJobFlag                  bool                `json:"HotJobFlag"`
	// WorkplaceTypeCode           interface{}         `json:"WorkplaceTypeCode"`
	// JobFamily                   interface{}         `json:"JobFamily"`
	// JobFunction                 interface{}         `json:"JobFunction"`
	// WorkerType                  interface{}         `json:"WorkerType"`
	// ContractType                interface{}         `json:"ContractType"`
	// ManagerLevel                interface{}         `json:"ManagerLevel"`
	// JobSchedule                 interface{}         `json:"JobSchedule"`
	// JobShift                    interface{}         `json:"JobShift"`
	// JobType                     interface{}         `json:"JobType"`
	// StudyLevel                  interface{}         `json:"StudyLevel"`
	// DomesticTravelRequired      interface{}         `json:"DomesticTravelRequired"`
	// InternationalTravelRequired interface{}         `json:"InternationalTravelRequired"`
	// WorkDurationYears           interface{}         `json:"WorkDurationYears"`
	// WorkDurationMonths          interface{}         `json:"WorkDurationMonths"`
	// WorkHours                   interface{}         `json:"WorkHours"`
	// WorkDays                    interface{}         `json:"WorkDays"`
	// LegalEmployer               interface{}         `json:"LegalEmployer"`
	// BusinessUnit                interface{}         `json:"BusinessUnit"`
	// Department                  interface{}         `json:"Department"`
	// Organization                interface{}         `json:"Organization"`
	// MediaThumbURL               interface{}         `json:"MediaThumbURL"`
	// ShortDescriptionStr         string              `json:"ShortDescriptionStr"`
	PrimaryLocation string `json:"PrimaryLocation"`
	// Distance                    int64               `json:"Distance"`
	// TrendingFlag                bool                `json:"TrendingFlag"`
	// BeFirstToApplyFlag          bool                `json:"BeFirstToApplyFlag"`
	// Relevancy                   float64             `json:"Relevancy"`
	// WorkplaceType               string              `json:"WorkplaceType"`
	// ExternalQualificationsStr   interface{}         `json:"ExternalQualificationsStr"`
	// ExternalResponsibilitiesStr interface{}         `json:"ExternalResponsibilitiesStr"`
	// SecondaryLocations          []SecondaryLocation `json:"secondaryLocations"`
	// RequisitionFlexFields       []interface{}       `json:"requisitionFlexFields"`
}

// type SecondaryLocation struct {
// 	RequisitionLocationID int64       `json:"RequisitionLocationId"`
// 	GeographyNodeID       int64       `json:"GeographyNodeId"`
// 	GeographyID           int64       `json:"GeographyId"`
// 	Name                  string      `json:"Name"`
// 	CountryCode           Language    `json:"CountryCode"`
// 	Latitude              interface{} `json:"Latitude"`
// 	Longitude             interface{} `json:"Longitude"`
// }

// type EsFacet struct {
// 	ID         string `json:"Id"`
// 	Name       string `json:"Name"`
// 	TotalCount int64  `json:"TotalCount"`
// }

// type Link struct {
// 	Rel  string `json:"rel"`
// 	Href string `json:"href"`
// 	Name string `json:"name"`
// 	Kind string `json:"kind"`
// }

// type Language string
// const (
// 	Us Language = "US"
// )

// type PrimaryLocation string
// const (
// 	ChicagoILUnitedStates PrimaryLocation = "Chicago, IL, United States"
// 	SunnyvaleCAUnitedStates PrimaryLocation = "Sunnyvale, CA, United States"
// )

func GetFortinetJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Fortinet")
	client := common.GetClient()

	url := "https://edel.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions?onlyData=true&expand=requisitionList.secondaryLocations,flexFieldsFacet.values,requisitionList.requisitionFlexFields&finder=findReqs;siteNumber=CX_2001,facetsList=LOCATIONS%3BWORK_LOCATIONS%3BWORKPLACE_TYPES%3BTITLES%3BCATEGORIES%3BORGANIZATIONS%3BPOSTING_DATES%3BFLEX_FIELDS,limit=12,lastSelectedFacet=TITLES,selectedLocationsFacet=300000000361862,selectedTitlesFacet=RD%3BIT%3BIS,sortBy=POSTING_DATES_DESC,limit=1000"

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating API request(Fortinet jobs): %v", err)
	}

	var fortinetJobs FortinetMain
	err = json.Unmarshal(resp.Body(), &fortinetJobs)
	if err != nil {
		return nil, fmt.Errorf("error parsing json (Fortinet jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range fortinetJobs.Items[0].RequisitionList {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "Fortinet",
			JobId:        common.Fortinet + ":" + job.ID,
			JobTitle:     job.Title,
			Location:     job.PrimaryLocation,
			PostedOn:     job.PostedDate,
			ExternalPath: "https://edel.fa.us2.oraclecloud.com/hcmUI/CandidateExperience/en/sites/CX_2001/job/" + job.ID,
		})
	}

	return jobPostings, nil
}
