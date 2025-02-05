package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.SalesForce, common.WorkdayPayload{
		Company: "SalesForce",
		CmpCode: common.SalesForce,
		PreURL:  "https://salesforce.wd12.myworkdayjobs.com/en-US/External_Career_Site",
		JobsURL: "https://salesforce.wd12.myworkdayjobs.com/wday/cxs/salesforce/External_Career_Site/jobs",
		PayLoad: `{
  "appliedFacets": {
    "CF_-_REC_-_LRV_-_Job_Posting_Anchor_-_Country_from_Job_Posting_Location_Extended": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "14fa3452ec7c1011f90d0002a2100000",
      "14fa3452ec7c1011f90cf661a7c80000",
      "14fa3452ec7c1011f90cfe3492140000",
      "14fa3452ec7c1011f90d056a9bc80000",
      "14fa3452ec7c1011f90cf2c552640000",
      "14fa3452ec7c1011f90cf8c9c5960000",
      "14fa3452ec7c1011f90cfc667aed0000",
      "14fa3452ec7c1011f90ce92475890000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
