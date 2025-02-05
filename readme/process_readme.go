package readme

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/process"
)

func ReadMeProcessNewJobs() error {
	jobs, err := process.GetProcessedNewJobs()
	if err != nil {
		fmt.Println("error while getting new processed jobs: ", err.Error())
	}

	return appendJobsToReadme(jobs)
}

func appendJobsToReadme(jobPostings []common.JobPosting) error {
	file, err := os.ReadFile("README.md")
	if err != nil {
		return fmt.Errorf("error reading README.md: %v", err)
	}

	content := string(file)

	tableMarker := "| --- | --- | --- | :---: | :---: |"
	splitContent := strings.Split(content, tableMarker)

	if len(splitContent) < 2 {
		return fmt.Errorf("table marker not found")
	}

	var newRows string
	today := time.Now().Format("Jan 02") // Get today's date as "Month Day"
	for _, job := range jobPostings {
		row := fmt.Sprintf("| **%s** | %s | %s | <a href=\"%s\" target=\"_blank\"><img src=\"https://i.imgur.com/u1KNU8z.png\" width=\"118\" alt=\"Apply\"></a> | %s |",
			job.Company, job.JobTitle, job.Location, job.ExternalPath, today)
		newRows += row + "\n"
	}

	newRows = strings.TrimSpace(newRows)

	if newRows == "" {
		fmt.Println("no new data!!!")
		return nil
	}
	updatedContent := splitContent[0] + tableMarker + "\n" + newRows + splitContent[1]

	err = os.WriteFile("README.md", []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to README.md: %v", err)
	}

	fmt.Println("Job postings appended successfully!")
	return nil
}
