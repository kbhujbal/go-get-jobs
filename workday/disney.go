package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Disney, common.WorkdayPayload{
		Company: "Disney",
		CmpCode: common.Disney,
		PreURL:  "https://disney.wd5.myworkdayjobs.com/en-US/disneycareer",
		JobsURL: "https://disney.wd5.myworkdayjobs.com/wday/cxs/disney/disneycareer/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locations": [
      "4f84d9e8a09701011a6aef3d93fc0000",
      "4f84d9e8a09701011a58839baa360000",
      "4f84d9e8a09701011a6afaad40e30000",
      "4f84d9e8a09701011a7497a5fab30000",
      "4f84d9e8a09701011a5965f4a0e60000",
      "4f84d9e8a09701011a5b2c8df7e60000",
      "4f84d9e8a09701011a6f6d34a6070000",
      "4f84d9e8a09701011a595216d7b20000",
      "4f84d9e8a09701011a7140a91d9e0000",
      "4f84d9e8a09701011a69fd38359c0000",
      "4f84d9e8a09701011a762a24b0900000",
      "4f84d9e8a09701011a568ed1a62e0000",
      "4f84d9e8a09701011a59d726dbab0000",
      "4f84d9e8a09701011a6ff52600b00000",
      "4f84d9e8a09701011a69c40e96710000",
      "4f84d9e8a09701011a71d2e0e0e50000",
      "4f84d9e8a09701011a72ab74b16d0000",
      "4f84d9e8a09701011a5948763c4b0000",
      "4f84d9e8a09701011a5a3cea04650000",
      "4f84d9e8a09701011a6ded3d6b520000",
      "4f84d9e8a09701011a75aec3e82c0000",
      "4f84d9e8a09701011a5a6a0fddb20000",
      "4f84d9e8a09701011a5e034bd6650000",
      "4f84d9e8a09701011a73b611c43c0000",
      "4f84d9e8a09701011a577e62d20b0000",
      "4f84d9e8a09701011a740910665e0000",
      "4f84d9e8a09701011a568bcec9810000",
      "4f84d9e8a09701011a69f86739820000",
      "4f84d9e8a09701011a66448e51f30000",
      "4f84d9e8a09701011a75abc108b40000",
      "4f84d9e8a09701011a572b5070c80000",
      "4f84d9e8a09701011a5fa89789280000"
    ],
    "jobFamilyGroup": [
      "4f84d9e8a097010115f146ed994a0000",
      "4f84d9e8a097010115f0cde63e9e0000",
      "4f84d9e8a097010115f0c30fe69e0000",
      "4f84d9e8a097010115f1495649a00000",
      "4f84d9e8a097010115f0b237c9d40000",
      "4f84d9e8a097010115f0fb1133370000",
      "4f84d9e8a097010115f104197eec0000",
      "4f84d9e8a097010115f0bed97c010000"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
