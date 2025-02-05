package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Target, common.WorkdayPayload{
		Company: "Target",
		CmpCode: common.Target,
		PreURL:  "https://target.wd5.myworkdayjobs.com/en-US/targetcareers",
		JobsURL: "https://target.wd5.myworkdayjobs.com/wday/cxs/target/targetcareers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "Location_Country": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "daccab9f1d25019c1fe2b73634578e0d",
      "daccab9f1d25018677ebcc363457460e",
      "daccab9f1d2501682445c5363457210e",
      "daccab9f1d250100d8b2c1363457f90d"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
