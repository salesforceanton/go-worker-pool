package activity_generator

import (
	"fmt"
	"math/rand"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

func GenerateUser(id int) User {
	return User{
		Id:    id + 1,
		Email: fmt.Sprintf("user%d@company.com", id+1),
		Logs:  generateLogs(rand.Intn(1000)),
	}
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
