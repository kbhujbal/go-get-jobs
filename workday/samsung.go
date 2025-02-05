package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Samsung, common.WorkdayPayload{
		Company: "Samsung",
		CmpCode: common.Samsung,
		PreURL:  "https://sec.wd3.myworkdayjobs.com/en-US/Samsung_Careers",
		JobsURL: "https://sec.wd3.myworkdayjobs.com/wday/cxs/sec/Samsung_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "Location_Country": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "189767dd6c9201b4198fe1a6db2997c7",
      "189767dd6c9201e189e3eaa6db299dc7",
      "189767dd6c9201fe2536dea6db2995c7",
      "189767dd6c920111f76cbba6db297fc7"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
