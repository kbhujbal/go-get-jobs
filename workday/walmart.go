package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Walmart, common.WorkdayPayload{
		Company: "Walmart",
		CmpCode: common.Walmart,
		PreURL:  "https://walmart.wd5.myworkdayjobs.com/en-US/WalmartExternal",
		JobsURL: "https://walmart.wd5.myworkdayjobs.com/wday/cxs/walmart/WalmartExternal/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "e83ebdbd2a0a01e7e1477a8948e904c6",
      "e83ebdbd2a0a0172c2dd788948e900c6",
      "e83ebdbd2a0a01af0185848948e94dc6",
      "e83ebdbd2a0a01050ff47e8948e912c6",
      "e83ebdbd2a0a01ea72c2808948e924c6",
      "e83ebdbd2a0a01cc71e67f8948e91ac6"
    ],
    "jobFamily": [
      "e83ebdbd2a0a013cdffb345b47e99dc4",
      "e83ebdbd2a0a012494773b5b47e9a1c4",
      "e83ebdbd2a0a01da76781e5a47e902c4",
      "24f96102afce1019026e1b0d04d90000",
      "e83ebdbd2a0a01bab19ce25947e9d8c3",
      "fc5cd4c5537f010c15b4226587ea0000",
      "e83ebdbd2a0a01e6af60e95a47e972c4",
      "e83ebdbd2a0a01bffcb4245a47e906c4",
      "e83ebdbd2a0a0170aca2465b47e9a7c4",
      "e83ebdbd2a0a01629c279b5a47e94ac4",
      "e83ebdbd2a0a017aa3a52b5b47e997c4",
      "e83ebdbd2a0a01bbf7f5df5947e9d6c3",
      "8d2227ff9cba01b0d76e16ec171b74b6",
      "e83ebdbd2a0a019d3ca3105b47e989c4",
      "b3d89f6d739710190268bf61b65c0000",
      "e83ebdbd2a0a01b9fbfa4d5a47e920c4",
      "e83ebdbd2a0a010c7bd4715a47e934c4",
      "e83ebdbd2a0a012c49c25f5a47e92ac4",
      "0933d0bbbee0010c15a34a03767d0000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
  }`,
	})
}
