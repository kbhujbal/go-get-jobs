package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.ASML, common.WorkdayPayload{
    Company: "ASML",
		CmpCode: common.ASML,
		PreURL:  "https://asml.wd3.myworkdayjobs.com/en-US/ASMLEXT1",
		JobsURL: "https://asml.wd3.myworkdayjobs.com/wday/cxs/asml/ASMLEXT1/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locations": [
      "4c9a1193e459100107474e5c2b090000",
      "4f25bf82e8451000f11e0b8942900000",
      "9d31b5b978021000f11edea15c370000",
      "4f25bf82e8451000f11cdae88c860000",
      "6009b7ae1ec71000a28b9738ba810000",
      "2b0ca4fa4f631000f11fb827f65e0000",
      "cc051a073cfe1000f11fa18fcc210000",
      "9d31b5b978021000f11ff70211770000",
      "cc051a073cfe1000f120477da6740000",
      "4f25bf82e8451000f11dc0f812980000",
      "9d31b5b978021000f11e9f0081330000"
    ],
    "jobFamilyGroup": [
      "719a7319274f01014848779465390000",
      "719a7319274f0101484876606da70001",
      "719a7319274f0101484875c6316d0002",
      "719a7319274f0101484875c6316d0000",
      "719a7319274f01014848752b77760001",
      "719a7319274f01014848705b12ec0004",
      "719a7319274f01014848735dc84c0002",
      "719a7319274f01014848735dc84c0000",
      "719a7319274f01014848718f5b8e0000",
      "719a7319274f010148486f26e0840000",
      "719a7319274f01014848718f5b8e0002",
      "719a7319274f01014848705b12ec0000",
      "719a7319274f0101484874914fbe0000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
