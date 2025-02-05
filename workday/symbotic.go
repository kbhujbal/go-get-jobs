package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Symbotic, common.WorkdayPayload{
		Company: "Symbotic",
		CmpCode: common.Symbotic,
		PreURL:  "https://symbotic.wd1.myworkdayjobs.com/en-US/Symbotic",
		JobsURL: "https://symbotic.wd1.myworkdayjobs.com/wday/cxs/symbotic/Symbotic/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "673bde061f5801325d4cc3656a016e10",
      "673bde061f580116636fc6656a017c10"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
