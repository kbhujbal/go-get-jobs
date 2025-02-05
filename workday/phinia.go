package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Phinia, common.WorkdayPayload{
		Company: "Phinia",
		CmpCode: common.Phinia,
		PreURL:  "https://phinia.wd5.myworkdayjobs.com/en-US/PHINIA_Careers",
		JobsURL: "https://phinia.wd5.myworkdayjobs.com/wday/cxs/phinia/PHINIA_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "Location_Country": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "f6fd9224b4eb100207a15ea926580000",
      "f6fd9224b4eb100207a1803aa5170000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
