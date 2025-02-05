package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.CCCIntelligentSolutions, common.WorkdayPayload{
		Company: "CCC Intelligent Solutions",
		CmpCode: common.CCCIntelligentSolutions,
		PreURL:  "https://cccis.wd1.myworkdayjobs.com/en-US/broadbean_external",
		JobsURL: "https://cccis.wd1.myworkdayjobs.com/wday/cxs/cccis/broadbean_external/jobs",
		PayLoad: `{
  "appliedFacets": {
    "JobFamilyGroup": [
      "d2044f38ceca01914e44c60cbe017f95",
      "d2044f38ceca0199dec6acfcbd017a95"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
