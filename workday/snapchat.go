package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Snapchat, common.WorkdayPayload{
		Company: "Snapchat",
		CmpCode: common.Snapchat,
		PreURL:  "https://wd1.myworkdaysite.com/en-US/recruiting/snapchat/snap",
		JobsURL: "https://wd1.myworkdaysite.com/wday/cxs/snapchat/snap/jobs",
		PayLoad: `{
  "appliedFacets": {
    "locations": [
      "efe1a86507310144a123773b020a0e37",
      "efe1a86507310187e01ef207030a7937",
      "2b0a835c9646011d58da08236e4f6726",
      "055b898d3e7c01a726fb4052c348eeaf",
      "08836f686d73101d8025e3c730de4ba5",
      "8bf70c1877bb01f58a864a033aab9149",
      "efe1a865073101e5380680f9020a7437",
      "a66859edee6201380c5d86b798075428",
      "a66859edee62013b741f38a9ea06dd1a",
      "b9cf6982655e1001a9ff7ae350d10000",
      "7a68e5b6d6b51001a9de39167c5f0000",
      "efe1a865073101b9db6c8da7020a6037",
      "256f279d5e741082c567c24fca236272",
      "efe1a8650731016c130aaddd010aed36",
      "ed80bd24de91105cf3b1aec9a82eb5a0",
      "256f279d5e741082c567c8a528f4627c",
      "efe1a86507310105e56ad10d020af736",
      "137dd6cbab601000bf830bdee83d0000",
      "efe1a865073101ddec60ef19020afc36",
      "c52c83bb81a21000cf303bd607c00000"
    ],
    "jobFamily": [
      "8d73f0a7971d102b9d459841e16ae3a5",
      "426a37839b0f0144214a29a56501496a",
      "80a3a1160116015e0d6b64caaa14b598",
      "8d73f0a7971d102b9fa6f985fc48edc6",
      "8d73f0a7971d102b9db74b4c3651e667",
      "80a3a116011601e60da38bf1aa14ba98",
      "426a37839b0f01b83be0bb8f6501ae69",
      "426a37839b0f01aad90037a865014b6a"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
