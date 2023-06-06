package readme

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/process"
)

// maxTableRows caps the total number of job rows in the README so that it
// stays under GitHub's 512KB markdown rendering limit (~1700 rows).
const maxTableRows = 1500

// allowedMonths defines which months to keep in the README.
// Since the README only stores "Mon DD" without a year, we can't distinguish
// years reliably. Set this to the months you want to retain.
var allowedMonths = map[string]bool{
	"Feb": true,
	"Mar": true,
}

func ReadMeProcessNewJobs() error {
	jobs, err := process.GetProcessedNewJobs()
	if err != nil {
		fmt.Println("error while getting new processed jobs: ", err.Error())
	}

	return appendJobsToReadme(jobs)
}

// isRowInAllowedMonth checks if a table row's date (last column, "Mon DD" format)
// belongs to one of the allowed months.
func isRowInAllowedMonth(row string) bool {
	row = strings.TrimSpace(row)
	if row == "" || !strings.HasPrefix(row, "|") {
		return false
	}

	cols := strings.Split(row, "|")
	dateStr := ""
	for i := len(cols) - 1; i >= 0; i-- {
		col := strings.TrimSpace(cols[i])
		if col != "" {
			dateStr = col
			break
		}
	}

	if len(dateStr) < 3 {
		return false
	}

	month := dateStr[:3]
	return allowedMonths[month]
}

// extractLink extracts the href URL from a table row to use as a unique key.
func extractLink(row string) string {
	idx := strings.Index(row, "href=\"")
	if idx == -1 {
		return ""
	}
	start := idx + len("href=\"")
	end := strings.Index(row[start:], "\"")
	if end == -1 {
		return ""
	}
	return row[start : start+end]
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

	seen := make(map[string]bool)

	// Build new rows from today's scraped jobs
	var newRows []string
	today := time.Now().Format("Jan 02")
	for _, job := range jobPostings {
		row := fmt.Sprintf("| **%s** | %s | %s | <a href=\"%s\" target=\"_blank\"><img src=\"https://i.imgur.com/u1KNU8z.png\" width=\"118\" alt=\"Apply\"></a> | %s |",
			job.Company, job.JobTitle, job.Location, job.ExternalPath, today)
		link := extractLink(row)
		if link != "" && !seen[link] {
			seen[link] = true
			newRows = append(newRows, row)
		}
	}

	// Filter existing rows: keep allowed months + deduplicate by link
	existingLines := strings.Split(splitContent[1], "\n")
	var keptRows []string
	var footerLines []string
	for _, line := range existingLines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || !strings.HasPrefix(trimmed, "|") {
			footerLines = append(footerLines, line)
			continue
		}
		if !isRowInAllowedMonth(trimmed) {
			continue
		}
		link := extractLink(trimmed)
		if link != "" && seen[link] {
			continue // duplicate
		}
		if link != "" {
			seen[link] = true
		}
		keptRows = append(keptRows, line)
	}

	if len(newRows) == 0 && len(keptRows) == 0 {
		fmt.Println("no new data!!!")
		return nil
	}

	// Merge new + existing rows, capped at maxTableRows (newest first)
	allRows := append(newRows, keptRows...)
	if len(allRows) > maxTableRows {
		allRows = allRows[:maxTableRows]
	}

	var sb strings.Builder
	sb.WriteString(splitContent[0])
	sb.WriteString(tableMarker)
	sb.WriteString("\n")
	for _, row := range allRows {
		sb.WriteString(row)
		sb.WriteString("\n")
	}
	sb.WriteString(strings.Join(footerLines, "\n"))

	err = os.WriteFile("README.md", []byte(sb.String()), 0644)
	if err != nil {
		return fmt.Errorf("error writing to README.md: %v", err)
	}

	fmt.Println("Job postings appended successfully!")
	return nil
}
