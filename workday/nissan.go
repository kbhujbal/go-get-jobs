package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Nissan, common.WorkdayPayload{
		Company: "Nissan",
		CmpCode: common.Nissan,
		PreURL:  "https://alliance.wd3.myworkdayjobs.com/en-US/nissanjobs",
		JobsURL: "https://alliance.wd3.myworkdayjobs.com/wday/cxs/alliance/nissanjobs/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "cf37143cc8d10124391f46d52ab0c118",
      "cf37143cc8d10131dc5bfca92ab04c18"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
