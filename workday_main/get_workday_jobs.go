package workdaymain

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
)

func GetWorkdayJobs(workdayPayload common.WorkdayPayload) ([]common.JobPosting, error) {
	fmt.Println("Processing: ", workdayPayload.Company)
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, 1)

	var jobPostings []common.JobPosting
	resp, err := workdayJobs(workdayPayload, 0)
	if err != nil {
		return nil, err
	}

	jobPostings = append(jobPostings, resp.JobPostings...)

	loop := int(resp.Total / 20)
	for i := 1; i <= loop; i++ {
		wg.Add(1)

		go func(offset int) {
			defer wg.Done()

			r, err := workdayJobs(workdayPayload, offset)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			jobPostings = append(jobPostings, r.JobPostings...)
			mu.Unlock()
		}(i * 20)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err = range errChan {
		if err != nil {
			return jobPostings, err
		}
	}

	return jobPostings, nil
}

func workdayJobs(workdayPayload common.WorkdayPayload, offset int) (*common.JobsResponse, error) {
	client := common.GetClient()

	workdayPayload.PayLoad = fmt.Sprintf(workdayPayload.PayLoad, offset)

	var jobsResponse common.JobsResponse
	var err error

	for attempts := 0; attempts < 3; attempts++ {
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetHeader("X-Calypso-CSRF-Token", "YOUR_CSRF_TOKEN").
			SetBody(workdayPayload.PayLoad).
			Post(workdayPayload.JobsURL)

		if err != nil {
			if attempts < 2 {
				time.Sleep(time.Second)
			}
			continue
		}

		if err := json.Unmarshal(resp.Body(), &jobsResponse); err != nil {
			return nil, fmt.Errorf("error parsing response: %v", err)
		}

		for i, job := range jobsResponse.JobPostings {
			jobsResponse.JobPostings[i].ExternalPath = workdayPayload.PreURL + job.ExternalPath
			jobid := strings.Split(jobsResponse.JobPostings[i].ExternalPath, "_")
			jobsResponse.JobPostings[i].JobId = workdayPayload.CmpCode + ":" + jobid[len(jobid)-1]
			jobsResponse.JobPostings[i].Company = workdayPayload.Company
		}

		return &jobsResponse, nil
	}

	return nil, fmt.Errorf("error fetching job listings after retries: %v", err)
}
