package sitesmain

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/sites"
)

// New function to fetch jobs based on company name
func FetchJobsByCompany(company string) ([]common.JobPosting, error) {
	switch company {
	case common.Google:
		return sites.GetGoogleJobs()
	case common.Microsoft:
		return sites.GetMicrosoftJobs()
	case common.Oracle:
		return sites.GetOracleJobs()
	case common.Apple:
		return sites.GetAppleJobs()
	case common.Meta:
		return sites.GetMetaJobs()
	case common.Tesla:
		return sites.GetTeslaJobs()
	case common.Chime:
		return sites.GetChimeJobs()
	case common.Splunk:
		return sites.GetSplunkJobs()
	case common.Visa:
		return sites.GetVisaJobs()
	case common.Uber:
		return sites.GetUberJobs()
	case common.Databricks:
		return sites.GetDatabricksJobs()
	case common.Amazon:
		return sites.GetAmazonJobs()
	case common.Amex:
		return sites.GetAmexJobs()
	case common.Snowflake:
		return sites.GetSnowflakeJobs()
	case common.Intuit:
		return sites.GetIntuitJobs()
	case common.IBM:
		return sites.GetIBMJobs()
	case common.ABB:
		return sites.GetABBJobs()
	case common.AMD:
		return sites.GetAMDJobs()
	case common.AkunaCapital:
		return sites.GetAkunaCapitalJobs()
	case common.Fortinet:
		return sites.GetFortinetJobs()
	case common.Reddit:
		return sites.GetRedditJobs()
	case common.NetApp:
		return sites.GetNetAppJobs()
	case common.Nokia:
		return sites.GetNokiaJobs()
	default:
		return nil, fmt.Errorf("unknown company: %s", company)
	} //////////////////////// Edit here
}
