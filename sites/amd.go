package sites

import (
	"encoding/json"
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
)

type AMDMain struct {
	Jobs       []AMDJob `json:"jobs"`
	TotalCount int64    `json:"totalCount"`
	Count      int64    `json:"count"`
	// Filter         Filter         `json:"filter"`
	// LanguageCounts LanguageCounts `json:"languageCounts"`
	// RequestID      bool           `json:"request_id"`
	// MetaData       MainMetaData   `json:"meta_data"`
	// Locations      []interface{}  `json:"locations"`
}

// type Filter struct {
// 	DisplayLimit     int64     `json:"displayLimit"`
// 	Categories       Brands    `json:"categories"`
// 	Brands           Brands    `json:"brands"`
// 	ExperienceLevels Brands    `json:"experienceLevels"`
// 	Locations        Brands    `json:"locations"`
// 	FacetList        FacetList `json:"facetList"`
// }

// type Brands struct {
// 	All       []All `json:"all"`
// 	Shortlist []All `json:"shortlist"`
// }

// type All struct {
// 	Category *NameEnum `json:"category,omitempty"`
// 	NumJobs  *int64    `json:"numJobs,omitempty"`
// 	City     *string   `json:"city"`
// 	State    *State    `json:"state,omitempty"`
// 	Country  *Country  `json:"country,omitempty"`
// 	Count    *int64    `json:"count,omitempty"`
// }

// type FacetList struct {
// 	City    []City `json:"city"`
// 	Country []City `json:"country"`
// 	Tags1   []City `json:"tags1"`
// }

// type City struct {
// 	Term  string `json:"term"`
// 	Count int64  `json:"count"`
// }

type AMDJob struct {
	Data AMDData `json:"data"`
}

type AMDData struct {
	// Slug                   string             `json:"slug"`
	// Language               Language           `json:"language"`
	// Languages              []Language         `json:"languages"`
	ReqID string `json:"req_id"`
	Title string `json:"title"`
	// Description            string             `json:"description"`
	LocationName string `json:"location_name"`
	// StreetAddress          *string            `json:"street_address,omitempty"`
	// City                   *string            `json:"city,omitempty"`
	// State                  State              `json:"state"`
	// Country                Country            `json:"country"`
	// CountryCode            Code               `json:"country_code"`
	// PostalCode             *string            `json:"postal_code,omitempty"`
	// LocationType           LocationType       `json:"location_type"`
	// Latitude               float64            `json:"latitude"`
	// Longitude              float64            `json:"longitude"`
	// Categories             []CategoryClass    `json:"categories"`
	// Tags1                  []Tags1            `json:"tags1"`
	// Tags2                  []string           `json:"tags2"`
	// Tags3                  []string           `json:"tags3"`
	// Tags4                  []Tags4            `json:"tags4"`
	// Tags5                  []string           `json:"tags5"`
	// Tags6                  []string           `json:"tags6"`
	// Tags7                  []string           `json:"tags7"`
	// Tags8                  []string           `json:"tags8"`
	// PromotionValue         int64              `json:"promotion_value"`
	// EmploymentType         *EmploymentType    `json:"employment_type,omitempty"`
	// Qualifications         string             `json:"qualifications"`
	HiringOrganization string `json:"hiring_organization"`
	// HiringOrganizationLogo string             `json:"hiring_organization_logo"`
	// Responsibilities       string             `json:"responsibilities"`
	PostedDate string `json:"posted_date"`
	ApplyURL   string `json:"apply_url"`
	// Internal               bool               `json:"internal"`
	// Searchable             bool               `json:"searchable"`
	// Applyable              bool               `json:"applyable"`
	// LiEasyApplyable        bool               `json:"li_easy_applyable"`
	// AtsCode                AtsCode            `json:"ats_code"`
	// MetaData               DataMetaData       `json:"meta_data"`
	// UpdateDate             *string            `json:"update_date,omitempty"`
	// CreateDate             string             `json:"create_date"`
	// Category               []DataCategoryEnum `json:"category"`
	// FullLocation           string             `json:"full_location"`
	// ShortLocation          string             `json:"short_location"`
}

// type CategoryClass struct {
// 	Name NameEnum `json:"name"`
// }

// type DataMetaData struct {
// 	ClientCode      ClientCode   `json:"client_code"`
// 	Googlejobs      Googlejobs   `json:"googlejobs"`
// 	Icims           IcimsClass   `json:"icims"`
// 	ImportID        string       `json:"import_id"`
// 	ImportSource    ImportSource `json:"import_source"`
// 	RedirectOnApply bool         `json:"redirectOnApply"`
// 	CanonicalURL    string       `json:"canonical_url"`
// 	LastMod         string       `json:"last_mod"`
// 	Gdpr            bool         `json:"gdpr"`
// }

// type Googlejobs struct {
// 	JobName           string      `json:"jobName"`
// 	CompanyName       string      `json:"companyName"`
// 	JobHash           string      `json:"jobHash"`
// 	DerivedInfo       DerivedInfo `json:"derivedInfo"`
// 	JobSummary        string      `json:"jobSummary"`
// 	JobTitleSnippet   string      `json:"jobTitleSnippet"`
// 	SearchTextSnippet string      `json:"searchTextSnippet"`
// }

// type DerivedInfo struct {
// 	JobCategories []JobCategory `json:"jobCategories"`
// 	Locations     []Location    `json:"locations"`
// }

// type Location struct {
// 	LatLng        LatLngClass      `json:"latLng"`
// 	LocationType  LocationTypeEnum `json:"locationType"`
// 	PostalAddress PostalAddress    `json:"postalAddress"`
// 	RadiusInMiles float64          `json:"radiusInMiles"`
// }

// type LatLngClass struct {
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

// type PostalAddress struct {
// 	AddressLines       []string               `json:"addressLines"`
// 	AdministrativeArea AdministrativeAreaEnum `json:"administrativeArea"`
// 	LanguageCode       string                 `json:"languageCode"`
// 	Locality           string                 `json:"locality"`
// 	Organization       string                 `json:"organization"`
// 	PostalCode         string                 `json:"postalCode"`
// 	Recipients         []interface{}          `json:"recipients"`
// 	RegionCode         Code                   `json:"regionCode"`
// 	Revision           int64                  `json:"revision"`
// 	SortingCode        string                 `json:"sortingCode"`
// 	Sublocality        string                 `json:"sublocality"`
// }

// type IcimsClass struct {
// 	ConfigKeys              interface{}             `json:"config_keys"`
// 	DateUpdated             string                  `json:"date_updated"`
// 	JpsIsPublic             bool                    `json:"jps_is_public"`
// 	PrimaryPostedSiteObject PrimaryPostedSiteObject `json:"primary_posted_site_object"`
// 	RevisionInt             int64                   `json:"revision_int"`
// 	UUID                    string                  `json:"uuid"`
// }

// type PrimaryPostedSiteObject struct {
// 	DatePosted string `json:"datePosted"`
// 	Site       Site   `json:"site"`
// 	SiteID     string `json:"siteId"`
// }

// type LanguageCounts struct {
// 	EnUs EnUsClass `json:"en-us"`
// }

// type EnUsClass struct {
// 	DisplayName string `json:"displayName"`
// 	Count       int64  `json:"count"`
// }

// type MainMetaData struct {
// 	ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
// }

// type ResponseMetadata struct {
// 	RequestID string `json:"requestId"`
// }

// type NameEnum string
// const (
// 	Engineering NameEnum = "Engineering"
// 	InformationTechnology NameEnum = "Information Technology"
// 	StudentInternTemp NameEnum = "Student / Intern / Temp"
// )

// type Country string
// const (
// 	UnitedStates Country = "United States"
// )

// type State string
// const (
// 	California State = "California"
// 	Colorado State = "Colorado"
// 	Florida State = "Florida"
// 	Massachusetts State = "Massachusetts"
// 	Minnesota State = "Minnesota"
// 	NewYork State = "New York"
// 	Oregon State = "Oregon"
// 	Texas State = "Texas"
// 	Washington State = "Washington"
// )

// type AtsCode string
// const (
// 	Icims AtsCode = "icims"
// )

// type DataCategoryEnum string
// const (
// 	CategoryEngineering DataCategoryEnum = " Engineering"
// 	CategoryInformationTechnology DataCategoryEnum = " Information Technology"
// 	CategoryStudentInternTemp DataCategoryEnum = " Student / Intern / Temp"
// )

// type Code string
// const (
// 	Us Code = "US"
// )

// type EmploymentType string
// const (
// 	FullTime EmploymentType = "FULL_TIME"
// )

// type HiringOrganization string
// const (
// 	AdvancedMicroDevicesInc HiringOrganization = "Advanced Micro Devices, Inc"
// )

// type Language string
// const (
// 	EnUs Language = "en-us"
// )

// type LocationType string
// const (
// 	LatLng LocationType = "LAT_LNG"
// )

// type ClientCode string
// const (
// 	AMD ClientCode = "amd"
// )

// type JobCategory string
// const (
// 	AccountingAndFinance JobCategory = "ACCOUNTING_AND_FINANCE"
// 	AdvertisingAndMarketing JobCategory = "ADVERTISING_AND_MARKETING"
// 	ArtFashionAndDesign JobCategory = "ART_FASHION_AND_DESIGN"
// 	BusinessOperations JobCategory = "BUSINESS_OPERATIONS"
// 	ComputerAndIt JobCategory = "COMPUTER_AND_IT"
// 	Management JobCategory = "MANAGEMENT"
// 	MediaCommunicationsAndWriting JobCategory = "MEDIA_COMMUNICATIONS_AND_WRITING"
// 	ProtectiveServices JobCategory = "PROTECTIVE_SERVICES"
// 	ScienceAndEngineering JobCategory = "SCIENCE_AND_ENGINEERING"
// )

// type LocationTypeEnum string
// const (
// 	AdministrativeArea LocationTypeEnum = "ADMINISTRATIVE_AREA"
// 	StreetAddress LocationTypeEnum = "STREET_ADDRESS"
// )

// type AdministrativeAreaEnum string
// const (
// 	CA AdministrativeAreaEnum = "CA"
// 	Co AdministrativeAreaEnum = "CO"
// 	FL AdministrativeAreaEnum = "FL"
// 	Ma AdministrativeAreaEnum = "MA"
// 	Mn AdministrativeAreaEnum = "MN"
// 	Ny AdministrativeAreaEnum = "NY"
// 	Or AdministrativeAreaEnum = "OR"
// 	Tx AdministrativeAreaEnum = "TX"
// 	Wa AdministrativeAreaEnum = "WA"
// )

// type Site string
// const (
// 	CampusAMD Site = "campus-amd"
// 	CareersAMD Site = "careers-amd"
// )

// type ImportSource string
// const (
// 	ImporterService ImportSource = "ImporterService"
// )

// type Tags1 string
// const (
// 	No Tags1 = "No"
// 	Yes Tags1 = "Yes"
// )

// type Tags4 string
// const (
// 	CampusUS Tags4 = "Campus US"
// 	USCareersExternal Tags4 = "US Careers (External)"
// )

func GetAMDJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "AMD")
	offset := 0
	allamdJobs, count, err := amdJobs(offset)
	if err != nil {
		fmt.Println("error processing amd jobs: ", err)
		return allamdJobs, err
	}

	for offset+100 < count {
		offset += 100
		job, _, err := amdJobs(offset)
		if err != nil {
			fmt.Println("error processing amd jobs: ", err.Error())
			continue
		}

		allamdJobs = append(allamdJobs, job...)
	}

	return allamdJobs, nil
}

func amdJobs(offset int) ([]common.JobPosting, int, error) {
	client := common.GetClient()

	url := fmt.Sprintf("https://careers.amd.com/api/jobs?country=United%%20States&categories=Engineering%%7CInformation%%20Technology%%7CStudent%%20/%%20Intern%%20/%%20Temp&sortBy=posted_date&descending=true&internal=false&limit=100&offset=%d", offset)

	resp, err := client.R().Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating API request(amd jobs): %v", err)
	}

	var amdJobs AMDMain
	err = json.Unmarshal(resp.Body(), &amdJobs)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing json (amd jobs): %v", err)
	}

	var jobPostings []common.JobPosting
	for _, job := range amdJobs.Jobs {
		jobPostings = append(jobPostings, common.JobPosting{
			Company:      "AMD",
			JobId:        common.AMD + ":" + job.Data.ReqID,
			JobTitle:     job.Data.Title,
			Location:     job.Data.LocationName,
			ExternalPath: job.Data.ApplyURL,
		})
	}

	return jobPostings, int(amdJobs.TotalCount), nil
}
