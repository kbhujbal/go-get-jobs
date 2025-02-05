package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Comcast, common.WorkdayPayload{
		Company: "Comcast",
		CmpCode: common.Comcast,
		PreURL:  "https://comcast.wd5.myworkdayjobs.com/en-US/Comcast_Careers",
		JobsURL: "https://comcast.wd5.myworkdayjobs.com/wday/cxs/comcast/Comcast_Careers/jobs",
		PayLoad: `{
  "appliedFacets": {
    "redirect": [
      "/Comcast_Careers/job/PA---Philadelphia-1701-John-F-Kennedy-Blvd/XMLNAME-1165-Software-Engineer_R396051/apply"
    ],
    "jobFamilyGroup": [
      "285386867dd9010665c0d2c57d0b4c15",
      "285386867dd901a0ab5627c67d0b5815"
    ],
    "locations": [
      "38d640cf23a80123d0470a517e27cc0d",
      "247bad0b5244013e714617da39329502",
      "38d640cf23a8014c35694aaf7f27223e",
      "3181e8987c761000fffdfa7d8ff20000",
      "41f07387769810015811bdd4deed0000",
      "38d640cf23a8010f4fd1947e7f27773a",
      "38d640cf23a801a1f2ccdf6b7f270039",
      "38d640cf23a8018e1b28496e7f273239",
      "38d640cf23a801137ec9cf5c7f27c037",
      "e11faec32347010cd533e61f2b032a5d",
      "38d640cf23a801449bd6486b7e273018",
      "38d640cf23a8012968217b5d7e270214",
      "38d640cf23a8014edd0b42517e27ef0d",
      "38d640cf23a801eb9c2d57707e278419",
      "38d640cf23a801cea29b6c697e27b317",
      "38d640cf23a801f336c2f7697e27db17",
      "38d640cf23a801dee8e69eb87e271427"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
