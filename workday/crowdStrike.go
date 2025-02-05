package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.CrowdStrike, common.WorkdayPayload{
		Company: "CrowdStrike",
		CmpCode: common.CrowdStrike,
		PreURL:  "https://crowdstrike.wd5.myworkdayjobs.com/en-US/crowdstrikecareers",
		JobsURL: "https://crowdstrike.wd5.myworkdayjobs.com/wday/cxs/crowdstrike/crowdstrikecareers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "Job_Family": [
      "cb19f044639b1001f6a02595bc920000",
      "1408861ee6e201641be2c2f6b000c00b",
      "1408861ee6e20197f95adbf6b000d20b",
      "1408861ee6e201d67af3e0f6b000d60b",
      "1408861ee6e2015adbe3e7f6b000de0b",
      "1408861ee6e201df327f0ff7b000fa0b"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
