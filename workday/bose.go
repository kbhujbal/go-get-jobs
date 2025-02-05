package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Bose, common.WorkdayPayload{
		Company: "Bose",
		CmpCode: common.Bose,
		PreURL:  "https://boseallaboutme.wd1.myworkdayjobs.com/en-US/Bose_Careers",
		JobsURL: "https://boseallaboutme.wd1.myworkdayjobs.com/wday/cxs/boseallaboutme/Bose_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locations": [
      "c286d09839da010d3342d8a79a599c38",
      "07c71a0b84a610019dad97de85ac0000",
      "31b7f9e93b0f10995df085940a703858",
      "31b7f9e93b0f10995df07763ee20383f",
      "31b7f9e93b0f10995df07be31d403844",
      "61fc0dd0d435018a4474dc689a11aa2a",
      "525231f0186e0117821dc4641b2459de"
    ],
    "jobFamilyGroup": [
      "50f881f2a6861063f9a960a0b0b6670f",
      "50f881f2a6861063f9f901550b256764",
      "50f881f2a6861063f9ba6a019059671d",
      "50f881f2a6861063f990cb0c2cda66c5"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
