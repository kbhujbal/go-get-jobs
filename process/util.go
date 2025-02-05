package process

import (
	"encoding/json"
	"os"

	commonconst "github.com/neyaadeez/go-get-jobs/common_const"
)

func loadJobIDs() (map[string]struct{}, error) {
	file, err := os.Open(commonconst.JobIdFile)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]struct{}), nil
		}
		return nil, err
	}
	defer file.Close()

	var jobIDs []string
	if err := json.NewDecoder(file).Decode(&jobIDs); err != nil {
		return nil, err
	}

	jobIDSet := make(map[string]struct{})
	for _, id := range jobIDs {
		jobIDSet[id] = struct{}{}
	}
	return jobIDSet, nil
}

func saveJobIDs(jobIDSet map[string]struct{}) error {
	jobIDs := make([]string, 0, len(jobIDSet))
	for id := range jobIDSet {
		jobIDs = append(jobIDs, id)
	}

	data, err := json.Marshal(jobIDs)
	if err != nil {
		return err
	}
	return os.WriteFile(commonconst.JobIdFile, data, 0644)
}
