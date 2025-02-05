package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Blueorigin, common.WorkdayPayload{
		Company: "Blueorigin",
		CmpCode: common.Blueorigin,
		PreURL:  "https://blueorigin.wd5.myworkdayjobs.com/en-US/BlueOrigin",
		JobsURL: "https://blueorigin.wd5.myworkdayjobs.com/wday/cxs/blueorigin/BlueOrigin/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "5f32d2b8465201de6df43ce63e17db4f",
      "5f32d2b8465201563136de3c38176145",
      "5a2167ca5bdb1001228cb3bd374c0000",
      "5f32d2b8465201b51255d2713817d845"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
