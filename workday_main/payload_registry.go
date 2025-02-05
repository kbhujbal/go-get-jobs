package workdaymain

import "github.com/neyaadeez/go-get-jobs/common"

var WorkdayPayloads = map[string]common.WorkdayPayload{}

func RegisterPayload(companyCode string, payload common.WorkdayPayload) {
	WorkdayPayloads[companyCode] = payload
}
