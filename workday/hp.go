package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.HP, common.WorkdayPayload{
		Company: "HP",
		CmpCode: common.HP,
		PreURL:  "https://hp.wd5.myworkdayjobs.com/en-US/ExternalCareerSite",
		JobsURL: "https://hp.wd5.myworkdayjobs.com/wday/cxs/hp/ExternalCareerSite/jobs",
		PayLoad: `{
  "appliedFacets": {
    "Location_Country": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
