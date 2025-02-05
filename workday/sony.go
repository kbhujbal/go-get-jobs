package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Sony, common.WorkdayPayload{
		Company: "Sony",
		CmpCode: common.Sony,
		PreURL:  "https://sonyglobal.wd1.myworkdayjobs.com/en-US/SonyGlobalCareers",
		JobsURL: "https://sonyglobal.wd1.myworkdayjobs.com/wday/cxs/sonyglobal/SonyGlobalCareers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "7306bd11847f108d56a585fb30065499",
      "7306bd11847f108d56a689b7002554ab",
      "7306bd11847f108d56a4475431ec5473",
      "7306bd11847f108d56a4602612245479",
      "bf0e94cb2dac0126e6d1bfd19e01df13",
      "0ec8556502e701f470f19b989b32708f",
      "7306bd11847f108d56a5e004ccc354a1",
      "7306bd11847f108d56a59b3e274c549b"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
