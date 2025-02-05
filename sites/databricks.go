package sites

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/neyaadeez/go-get-jobs/common"
)

type DatabricksMain struct {
	ComponentChunkName string           `json:"componentChunkName"`
	Path               string           `json:"path"`
	Result             DatabricksResult `json:"result"`
}

type DatabricksResult struct {
	PageContext DatabricksPageContext `json:"pageContext"`
}

type DatabricksPageContext struct {
	Data DatabricksPageContextData `json:"data"`
}

type DatabricksPageContextData struct {
	AllGreenhouseDepartment DatabricksAllGreenhouseDepartment `json:"allGreenhouseDepartment"`
}

type DatabricksAllGreenhouseDepartment struct {
	Nodes []DatabricksAllGreenhouseDepartmentNode `json:"nodes"`
}

type DatabricksAllGreenhouseDepartmentNode struct {
	Jobs []DatabricksParentJob `json:"jobs"`
}

type DatabricksParentJob struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	AbsoluteURL   string `json:"absolute_url"`
	GhID          int64  `json:"gh_Id"`
	InternalJobID int64  `json:"internal_job_id"`
	// UpdatedAt     string      `json:"updated_at"`
	Location    DatabricksLocation   `json:"location"`
	Departments []DatabricksLocation `json:"departments"`
}

type DatabricksLocation struct {
	Name string `json:"name"`
}

type DatabricksParent struct {
	ID   string                `json:"id"`
	Name string                `json:"name"`
	Jobs []DatabricksParentJob `json:"jobs"`
}

type DatabricksAllGreenhouseOffice struct {
	Nodes []DatabricksAllGreenhouseOfficeNode `json:"nodes"`
}

type DatabricksAllGreenhouseOfficeNode struct {
	GhID     int64                                   `json:"gh_Id"`
	ID       string                                  `json:"id"`
	Name     string                                  `json:"name"`
	ParentID *int64                                  `json:"parent_id"`
	Jobs     []DatabricksChildrenGreenhouseOfficeJob `json:"jobs"`
	Children []DatabricksChild                       `json:"children"`
}

type DatabricksChild struct {
	GhID                     int64                                   `json:"gh_Id"`
	ParentID                 int64                                   `json:"parent_id"`
	Name                     string                                  `json:"name"`
	Jobs                     []DatabricksChildrenGreenhouseOfficeJob `json:"jobs"`
	ChildrenGreenhouseOffice []DatabricksChildrenGreenhouseOffice    `json:"childrenGreenhouseOffice"`
}

type DatabricksChildrenGreenhouseOffice struct {
	GhID                     int64                                   `json:"gh_Id"`
	ParentID                 int64                                   `json:"parent_id"`
	Name                     string                                  `json:"name"`
	Jobs                     []DatabricksChildrenGreenhouseOfficeJob `json:"jobs"`
	ChildrenGreenhouseOffice []DatabricksChildrenGreenhouseOffice    `json:"childrenGreenhouseOffice,omitempty"`
}

type DatabricksChildrenGreenhouseOfficeJob struct {
	GhID int64 `json:"gh_Id"`
}

// type FluffyDrupal struct {
// 	AvailableLanguages []AvailableLanguage  `json:"availableLanguages"`
// 	GeneralPages       []GeneralPageElement `json:"generalPages"`
// }

// type AvailableLanguage struct {
// 	Name   string `json:"name"`
// 	Prefix string `json:"prefix"`
// 	ID     string `json:"id"`
// }

// type GeneralPageElement struct {
// 	ID                 string         `json:"id"`
// 	EntityURL          EntityURLClass `json:"entityUrl"`
// 	EntityTranslations []Translation  `json:"entityTranslations"`
// 	Nid                int64          `json:"nid"`
// 	Vid                int64          `json:"vid"`
// }

// type Translation struct {
// 	EntityLanguage EntityLanguage `json:"entityLanguage"`
// 	EntityURL      EntityURLClass `json:"entityUrl"`
// }

// type EntityLanguage struct {
// 	ID ID `json:"id"`
// }

// type SlicesMap struct {
// }

// type Type string

// const (
// 	DrupalMetaLink     Type = "Drupal_MetaLink"
// 	DrupalMetaProperty Type = "Drupal_MetaProperty"
// 	DrupalMetaValue    Type = "Drupal_MetaValue"
// )

// const (
// 	Administration       Name = "Administration"
// 	BusinessDevelopment  Name = "Business Development"
// 	CustomerSuccess      Name = "Customer Success"
// 	Engineering          Name = "Engineering"
// 	FieldEngineering     Name = "Field Engineering"
// 	Finance              Name = "Finance"
// 	GA                   Name = "G&A"
// 	It                   Name = "IT"
// 	Legal                Name = "Legal"
// 	Marketing            Name = "Marketing"
// 	MosaicAI             Name = "Mosaic AI"
// 	Operations           Name = "Operations"
// 	People               Name = "People"
// 	PeopleAndHR          Name = "People and HR"
// 	Product              Name = "Product"
// 	ProfessionalServices Name = "Professional Services"
// 	Research             Name = "Research"
// 	Security             Name = "Security"
// 	UniversityRecruiting Name = "University Recruiting"
// )

// type ID string

// const (
// 	Br   ID = "br"
// 	De   ID = "de"
// 	En   ID = "en"
// 	Fr   ID = "fr"
// 	IDIt ID = "it"
// 	Ja   ID = "ja"
// 	Ko   ID = "ko"
// )

func GetDatabricksJobs() ([]common.JobPosting, error) {
	fmt.Println("Processing: ", "Databricks")
	var jobPostings []common.JobPosting

	url := "https://www.databricks.com/careers-assets/page-data/company/careers/open-positions/page-data.json?department=Engineering,University%20Recruiting,IT,Security&location=United%20States"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching databricks jobs: %v", err)
	}
	defer resp.Body.Close()

	var databricksJobs DatabricksMain
	if err := json.NewDecoder(resp.Body).Decode(&databricksJobs); err != nil {
		return nil, fmt.Errorf("error decoding databricks response: %v", err)
	}

	for _, node := range databricksJobs.Result.PageContext.Data.AllGreenhouseDepartment.Nodes {
		for _, job := range node.Jobs {
			if isUSLocation(job.Location.Name) && isCSRelatedJobDatabricks(job.Departments) {
				jobPosting := common.JobPosting{
					JobId:        common.Databricks + ":" + strconv.Itoa(int(job.InternalJobID)),
					JobTitle:     job.Title,
					Location:     job.Location.Name,
					ExternalPath: job.AbsoluteURL,
					Company:      "Databricks",
				}
				jobPostings = append(jobPostings, jobPosting)
			}
		}
	}

	return jobPostings, nil
}

func isCSRelatedJobDatabricks(dp []DatabricksLocation) bool {
	for _, l := range dp {
		if databricksCheckDepartment(l.Name) {
			return true
		}
	}
	return false
}

func databricksCheckDepartment(d string) bool {
	switch d {
	case "University Recruiting":
		return true

	case "Security":
		return true

	case "AI & Robotics":
		return true

	case "Engineering":
		return true

	case "IT":
		return true

	case "Mosaic AI":
		return true

	case "Product":
		return true

	default:
		return false
	}
}
