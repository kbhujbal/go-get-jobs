package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.SPGlobal, common.WorkdayPayload{
		Company: "SPGLOBAL",
		CmpCode: common.SPGlobal,
		PreURL:  "https://spgi.wd5.myworkdayjobs.com/en-US/SPGI_Careers",
		JobsURL: "https://spgi.wd5.myworkdayjobs.com/wday/cxs/spgi/SPGI_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "jobFamilyGroup": [
      "13f9469a2d3510519478eef523d49d51",
      "13f9469a2d3510519478d9678fd49d47",
      "13f9469a2d3510519478fa95ea849d57"
    ],
    "Location_Region_State_Province": [
      "a9336ce88419106a5a379016393d6c88",
      "038c1c3a6e91106d28193c29c3b91720",
      "a9336ce88419106a5a3846968b956cce",
      "a9336ce88419106a5a3943e151956d2d",
      "a9336ce88419106a5a38a569014d6cf1",
      "a9336ce88419106a5a381f6e3ced6cbf",
      "a9336ce88419106a5a37a9fa4a656c92",
      "cc3052ec2e910100afd00c905efd0000",
      "0565124f9443100204efcd7c6ebd0000",
      "15e67fc7471910005ed114a38b9a0000",
      "15e67fc7471910005ed14b8216980000",
      "4c151630730401f3945472175b6320bd",
      "87455d27fca90194c8baddf35101edd3",
      "87455d27fca901cb4780b3d44d01d4b6",
      "4c151630730401603108f7855c6310bf"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
