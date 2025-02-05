package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.NorthropGrumman, common.WorkdayPayload{
		Company: "Northrop Grumman",
		CmpCode: common.NorthropGrumman,
		PreURL:  "https://ngc.wd1.myworkdayjobs.com/en-US/Northrop_Grumman_External_Site",
		JobsURL: "https://ngc.wd1.myworkdayjobs.com/wday/cxs/ngc/Northrop_Grumman_External_Site/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "a111b0a898f1018af8043a63e7007178",
      "a111b0a898f10171f1426263e7007578"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
