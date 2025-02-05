package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Nvidia, common.WorkdayPayload{
		Company: "Nvidia",
		CmpCode: common.Nvidia,
		PreURL:  "https://nvidia.wd5.myworkdayjobs.com/en-US/NVIDIAExternalCareerSite",
		JobsURL: "https://nvidia.wd5.myworkdayjobs.com/wday/cxs/nvidia/NVIDIAExternalCareerSite/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationHierarchy1": [
      "2fcb99c455831013ea52fb338f2932d8"
    ],
    "jobFamilyGroup": [
      "0c40f6bd1d8f10ae43ffaefd46dc7e78",
      "0c40f6bd1d8f10ae43ffc3fc7d8c7e8a",
      "0c40f6bd1d8f10ae43ffc668c6847e8c",
      "0c40f6bd1d8f10ae43ffbd1459047e84"
    ],
    "workerSubType": [
      "ab40a98049581037a3ada55b087049b7"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
