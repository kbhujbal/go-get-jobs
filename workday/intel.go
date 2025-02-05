package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Intel, common.WorkdayPayload{
		Company: "Intel",
		CmpCode: common.Intel,
		PreURL:  "https://intel.wd1.myworkdayjobs.com/en-US/External",
		JobsURL: "https://intel.wd1.myworkdayjobs.com/wday/cxs/intel/External/jobs",
		PayLoad: `{
		"appliedFacets": {
			"locations": [
				"1e4a4eb3adf101b8aec18a77bf810dd0",
				"1e4a4eb3adf1018c4bf78f77bf8112d0",
				"1e4a4eb3adf1013ddb7bd877bf8153d0",
				"1e4a4eb3adf10129d05fe377bf815dd0",
				"1e4a4eb3adf10118b1dfe877bf8162d0",
				"1e4a4eb3adf10155d1cc0778bf8180d0",
				"1e4a4eb3adf101d4e5a61779bf8159d1",
				"1e4a4eb3adf10146fd5c5276bf81eece",
				"1e4a4eb3adf1011246675c76bf81f8ce",
				"1e4a4eb3adf1016541777876bf8111cf",
				"1e4a4eb3adf101fa2a777d76bf8116cf",
				"1e4a4eb3adf101770f350977bf8193cf",
				"1e4a4eb3adf10174f0548376bf811bcf",
				"1e4a4eb3adf101cc4e292078bf8199d0"
			]
		},
		"limit": 20,
		"offset": %d,
		"searchText": ""
	}`,
	})
}
