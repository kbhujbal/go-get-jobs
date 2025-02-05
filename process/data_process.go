package process

import (
	"fmt"
	"sync"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/database"
	sitesmain "github.com/neyaadeez/go-get-jobs/sites_main"
	"github.com/neyaadeez/go-get-jobs/workday"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

var (
	cachedJobs  []common.JobPosting
	cachedError error
	onceGetJobs sync.Once
)

func ProcessJobsWithDB() error {
	jobs, err := GetProcessedNewJobs()
	if err != nil {
		fmt.Println("error while processing new jobs: ", err.Error())
		return err
	}
	fmt.Println("Processed Jobs (New Jobs): ", len(jobs))

	// err = database.InsertIntoDB(jobs)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	return nil
}

func GetProcessedNewJobs() ([]common.JobPosting, error) {
	onceGetJobs.Do(func() {
		var allJobs []common.JobPosting

		jobs, err := GetAllWorkdayJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
		}
		fmt.Println("All Workday Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		for company := range common.SitesCompanies {
			jobs, err := sitesmain.FetchJobsByCompany(company)
			if err != nil {
				fmt.Println(err.Error())
				cachedError = err
			}
			fmt.Printf("All %s Jobs: %d\n", company, len(jobs))
			allJobs = append(allJobs, jobs...)
		}

		cachedJobs, cachedError = processDublicateJobs(allJobs)
		if cachedError != nil {
			fmt.Println(cachedError.Error())
		}
	})

	return cachedJobs, cachedError
}

func ProcessJobsWithDBForNewlyAddedJobPortal(company string, w bool) error {
	workday.Init()
	jobs, err := getProcessedNewJobsNewlyAddedJobPortal(company, w)
	if err != nil {
		fmt.Println("error while processing new jobs: ", err.Error())
		return err
	}
	fmt.Println("Processed Jobs (New Jobs): ", len(jobs))

	err = database.InsertIntoDB(jobs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func getProcessedNewJobsNewlyAddedJobPortal(company string, w bool) ([]common.JobPosting, error) {
	onceGetJobs.Do(func() {
		var allJobs []common.JobPosting

		if w {
			jobs, err := workdaymain.GetWorkdayJobs(workdaymain.WorkdayPayloads[company])
			if err != nil {
				fmt.Println(err.Error())
				cachedError = err
				return
			}
			fmt.Println(jobs[0])
			fmt.Println("All "+company+" Jobs: ", len(jobs))
			allJobs = append(allJobs, jobs...)

		} else {
			jobs, err := sitesmain.FetchJobsByCompany(company)
			if err != nil {
				fmt.Println(err.Error())
				cachedError = err
				return
			}
			fmt.Println("All "+company+" Jobs: ", len(jobs))
			allJobs = append(allJobs, jobs...)
		}

		fmt.Println(allJobs[0])
		cachedJobs, cachedError = processDublicateJobs(allJobs)
		if cachedError != nil {
			fmt.Println(cachedError.Error())
		}
	})

	return cachedJobs, cachedError
}
