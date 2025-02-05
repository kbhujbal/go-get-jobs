package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Tancent, common.WorkdayPayload{
		Company: "Tancent",
		CmpCode: common.Tancent,
		PreURL:  "https://tencent.wd1.myworkdayjobs.com/en-US/Tencent_Careers",
		JobsURL: "https://tencent.wd1.myworkdayjobs.com/wday/cxs/tencent/Tencent_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locations": [
      "b3d4dad114e4100177c032bef7130000",
      "1c8376485a4e100177e622668b570000",
      "1c8376485a4e100177e619f974680000",
      "b3d4dad114e4100177c0535a2e410000",
      "1c8376485a4e100177e6c3ec4b4d0000",
      "b32f1ee18078012f1fd236f105740000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
