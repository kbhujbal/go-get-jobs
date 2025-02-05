package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Medtronic, common.WorkdayPayload{
		Company: "Medtronic",
		CmpCode: common.Medtronic,
		PreURL:  "https://medtronic.wd1.myworkdayjobs.com/en-US/MedtronicCareers",
		JobsURL: "https://medtronic.wd1.myworkdayjobs.com/wday/cxs/medtronic/MedtronicCareers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "jobFamilyGroup": [
      "9f511399cde0412cb986049830df9cbd",
      "be3ab1d7a62801c1c7c82b804a0529d0",
      "46a4fe85ccfe40b3b1aef9d430a132d0",
      "2fe8588f35e84eb98ef535f4d738f243"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
