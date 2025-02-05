package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.CVS, common.WorkdayPayload{
		Company: "CVS",
		CmpCode: common.CVS,
		PreURL:  "https://cvshealth.wd1.myworkdayjobs.com/en-US/CVS_Health_Careers",
		JobsURL: "https://cvshealth.wd1.myworkdayjobs.com/wday/cxs/cvshealth/CVS_Health_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
    "e65dbadf6a50100168ed86fe4cf50001",
    "e65dbadf6a50100168ed7f2a693c0001"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
