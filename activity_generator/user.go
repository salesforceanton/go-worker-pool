package activity_generator

import (
	"fmt"
	"time"
)

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	Id    int
	Email string
	Logs  []logItem
}

func (u User) GetActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.Id, u.Email)
	for index, item := range u.Logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}
