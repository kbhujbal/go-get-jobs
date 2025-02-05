package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Mantech, common.WorkdayPayload{
		Company: "Mantech",
		CmpCode: common.Mantech,
		PreURL:  "https://mantech.wd1.myworkdayjobs.com/en-US/External",
		JobsURL: "https://mantech.wd1.myworkdayjobs.com/wday/cxs/mantech/External/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "fa130d59872001b048b9459f4300a850",
      "fa130d59872001cbb1a4479f4300ac50"
    ],
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
