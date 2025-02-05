package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Redhat, common.WorkdayPayload{
		Company: "Redhat",
		CmpCode: common.Redhat,
		PreURL:  "https://redhat.wd5.myworkdayjobs.com/en-US/jobs",
		JobsURL: "https://redhat.wd5.myworkdayjobs.com/wday/cxs/redhat/jobs/jobs",
		PayLoad: `{
  "appliedFacets": {
    "d": [
      "c18026e7757601cf6eb0136f4e43f04a",
      "c18026e77576010f6ef6126f4e43ec4a"
    ],
    "a": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
