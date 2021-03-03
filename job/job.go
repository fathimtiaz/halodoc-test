package job

import (
	"local/halodoc/config"
	"strings"
)

func getUsers(users []config.User, logType string) (userNames []string) {
	for _, user := range users {
		if strings.Contains(user.Logs, logType) {
			userNames = append(userNames, user.Name)
		}
	}

	return userNames
}
