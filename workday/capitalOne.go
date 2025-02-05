package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.CapitalOne, common.WorkdayPayload{
		Company: "Capital One",
		CmpCode: common.CapitalOne,
		PreURL:  "https://capitalone.wd12.myworkdayjobs.com/en-US/Capital_One",
		JobsURL: "https://capitalone.wd12.myworkdayjobs.com/wday/cxs/capitalone/Capital_One/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "a12c70bf789e105802e9caf800542991",
      "a12c70bf789e105802e9e79458dc29ab"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
