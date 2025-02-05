package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Boeing, common.WorkdayPayload{
    Company: "Boeing",
		CmpCode: common.Boeing,
		PreURL:  "https://boeing.wd1.myworkdayjobs.com/en-US/EXTERNAL_CAREERS",
		JobsURL: "https://boeing.wd1.myworkdayjobs.com/wday/cxs/boeing/EXTERNAL_CAREERS/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
	"bc33aa3152ec42d4995f4791a106ed09"
	],
    "jobFamilyGroup": [
      "8b618a30e00f0117c89a81e8143fad25",
      "8b618a30e00f01c7277572e8143f8b25",
      "8b618a30e00f0107641f90e8143fcb25",
      "8b618a30e00f01e427a1a8e8143fed25",
      "8b618a30e00f0162e6d16be8143f7925",
      "75851924369a1001a5385a64b7b80000",
      "bb56e5e6561a01481c721954090210f0",
      "8b618a30e00f01e0101a6be8143f7725",
      "8b618a30e00f01820e826ce8143f7b25",
      "8b618a30e00f01585ff870e8143f8725"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
