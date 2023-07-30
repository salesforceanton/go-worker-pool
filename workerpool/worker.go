package workerpool

import (
	"fmt"
	"os"
)

type SaveUserLogsWorker struct {
}

func (w *SaveUserLogsWorker) processJob(job SaveUserLogsJob) SaveUserLogsResult {
	user := job.User

	filename := fmt.Sprintf("users/uid%d.txt", user.Id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	result := SaveUserLogsResult{
		email: user.Email,
	}

	if err != nil {
		result.errorMessage = err.Error()
		return result
	}

	_, err = file.WriteString(user.GetActivityInfo())
	if err != nil {
		result.errorMessage = err.Error()
		return result
	}

	result.filepath = filename
	return result
}

func newSaveUserLogsWorker() *SaveUserLogsWorker {
	return &SaveUserLogsWorker{}
}
