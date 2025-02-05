package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Twitter, common.WorkdayPayload{
		Company: "Twitter",
		CmpCode: common.Twitter,
		PreURL:  "https://twitter.wd5.myworkdayjobs.com/en-US/X",
		JobsURL: "https://twitter.wd5.myworkdayjobs.com/wday/cxs/twitter/X/jobs",
		PayLoad: `{
  "appliedFacets": {},
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
