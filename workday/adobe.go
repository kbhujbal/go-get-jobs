package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Adobe, common.WorkdayPayload{
		Company: "Adobe",
		CmpCode: common.Adobe,
		PreURL:  "https://adobe.wd5.myworkdayjobs.com/en-US/external_experienced",
		JobsURL: "https://adobe.wd5.myworkdayjobs.com/wday/cxs/adobe/external_experienced/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "591af8b812fa10737b43a1662896f01c",
      "591af8b812fa10737af39db3d96eed9f",
      "591af8b812fa10737b0e880e0e3eeee9"
    ],
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
