package process

import "github.com/neyaadeez/go-get-jobs/common"

func processDublicateJobs(jobs []common.JobPosting) ([]common.JobPosting, error) {
	jobIDSet, err := loadJobIDs()
	if err != nil {
		return nil, err
	}

	var processedJobs []common.JobPosting
	for _, job := range jobs {
		if _, exists := jobIDSet[job.JobId]; !exists {
			processedJobs = append(processedJobs, job)
			jobIDSet[job.JobId] = struct{}{}
		}
	}

	if err := saveJobIDs(jobIDSet); err != nil {
		return nil, err
	}

	return processedJobs, nil
}
